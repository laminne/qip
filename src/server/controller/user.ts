import { FindUserService } from "../../service/user/find_user_service.js";
import { AsyncResult, Failure, Result, Success } from "../../helpers/result.js";
import {
  DomainToUserFollowResponse,
  UserFollowResponse,
  UserResponse,
} from "../types/user.js";
import { FindServerService } from "../../service/server/find_server_service.js";
import { FindPostService } from "../../service/post/find_post_service.js";
import {
  CommonMediaResponse,
  CommonPostResponse,
  PostReactionResponse,
} from "../types/post.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { PostData } from "../../service/data/post.js";
import { CreateFollowService } from "../../service/user/create_follow_service.js";

export class UserController {
  private readonly findUserService: FindUserService;
  private readonly findServerService: FindServerService;
  private readonly findPostService: FindPostService;
  private readonly createFollowService: CreateFollowService;

  constructor(args: {
    findUserService: FindUserService;
    findServerService: FindServerService;
    findPostService: FindPostService;
    createFollowService: CreateFollowService;
  }) {
    this.findUserService = args.findUserService;
    this.findServerService = args.findServerService;
    this.findPostService = args.findPostService;
    this.createFollowService = args.createFollowService;
  }

  async FindByHandle(name: string): AsyncResult<UserResponse, Error> {
    const acct = this.acctConverter(name);
    if (acct.isFailure()) {
      return new Failure(acct.value);
    }
    const user = await this.findUserService.FindByHandle(acct.value);
    if (user.isFailure()) {
      return new Failure(user.value);
    }

    const server = await this.findServerService.FindByID(user.value.serverID);
    if (server.isFailure()) {
      return new Failure(server.value);
    }

    const r = user.value;
    const res: UserResponse = {
      id: r.id,
      host: r.fullHandle,
      nickName: r.nickName,
      role: r.role,
      bio: r.bio,
      headerImageURL: r.headerImageURL,
      iconImageURL: r.iconImageURL,
      following: r.following.map((v) => {
        return {
          id: v.following.id,
          fullHandle: v.following.fullHandle,
          nickName: v.following.nickName,
          bio: v.following.bio,
          iconImageURL: v.following.iconImageURL,
        };
      }),
      softwareName: server.value.softwareName,
    };
    return new Success(res);
  }

  async FindUserPosts(id: string): AsyncResult<CommonPostResponse[], Error> {
    const u = await this.FindByHandle(id);
    if (u.isFailure()) {
      return new Failure(u.value);
    }
    // ユーザーの投稿を取ってくる
    const posts = await this.findPostService.FindByAuthor(
      u.value.id as Snowflake,
    );
    if (posts.isFailure()) {
      return new Failure(posts.value);
    }

    // ユーザーの情報を取ってくる
    // RN/BT/RTに相当するものがないのでこれで良い(全てリンクによる引用になる)
    const user = await this.findUserService.FindByID(u.value.id as Snowflake);
    if (user.isFailure()) {
      return new Failure(user.value);
    }

    return new Success(
      posts.value.map((v: PostData): CommonPostResponse => {
        return {
          id: v.id,
          author: {
            host: user.value.fullHandle,
            iconImageURL: user.value.iconImageURL,
            id: user.value.id,
            nickName: user.value.nickName,
          },
          createdAt: v.createdAt,
          reactions: v.reactions.map((v): PostReactionResponse => {
            return { postID: v.postID, userID: v.userID };
          }),
          attachments: v.attachments.map((v): CommonMediaResponse => {
            return {
              id: v.id,
              authorID: v.authorID,
              postID: v.postID,
              blurhash: v.blurhash,
              cached: v.cached,
              isSensitive: v.isSensitive,
              size: v.size,
              thumbnailURL: v.thumbnailURL,
              url: v.url,
              name: v.name,
              type: v.type,
              md5Sum: v.md5Sum,
            };
          }),
          text: v.text,
        };
      }),
    );
  }

  async CreateFollow(
    followerID: string,
    followingID: string,
  ): AsyncResult<UserFollowResponse, Error> {
    const res = await this.createFollowService.Handle(
      followingID as Snowflake,
      followerID as Snowflake,
    );
    if (res.isFailure()) {
      return new Failure(res.value);
    }
    return new Success(DomainToUserFollowResponse(res.value));
  }

  private acctConverter(acct: string): Result<string, Error> {
    const split = acct.split("@");
    switch (split.length) {
      case 2:
        return new Success(`${split[0]}@${split[1]}`);
      case 3:
        return new Success(`${split[1]}@${split[2]}`);
      default:
        // パースエラー
        return new Failure(new Error("failed to parse acct"));
    }
  }
}
