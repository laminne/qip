import { AsyncResult, Failure, Success } from "../../helpers/result";
import {
  CommonMediaResponse,
  CommonPostRequest,
  CommonPostResponse,
  PostReactionResponse,
} from "../types/post";
import { FindPostService } from "../../service/post/find_post_service";
import { Snowflake } from "../../helpers/id_generator";
import { FindUserService } from "../../service/user/find_user_service";
import { FindServerService } from "../../service/server/find_server_service";
import { CreatePostService } from "../../service/post/create_post_service";
import { PostData } from "../../service/data/post";
import { UserData } from "../../service/data/user";
import { ServerData } from "../../service/data/server";
import { NormalizeURL } from "../../helpers/url";

export class PostController {
  private readonly findPostService: FindPostService;
  private readonly findUserService: FindUserService;
  private readonly findServerService: FindServerService;
  private readonly createPostService: CreatePostService;

  constructor(args: {
    findPostService: FindPostService;
    findUserService: FindUserService;
    findServerService: FindServerService;
    createPostService: CreatePostService;
  }) {
    this.findPostService = args.findPostService;
    this.findUserService = args.findUserService;
    this.findServerService = args.findServerService;
    this.createPostService = args.createPostService;
  }

  async FindByID(id: string): AsyncResult<CommonPostResponse, Error> {
    const res = await this.findPostService.FindByID(id as Snowflake);
    if (res.isFailure()) {
      return new Failure(new Error("failed to find post", res.value));
    }
    const user = await this.findUserService.FindByID(res.value.authorID);
    if (user.isFailure()) {
      return new Failure(
        new Error("failed to find post author's data", user.value),
      );
    }
    const server = await this.findServerService.FindByID(user.value.serverID);
    if (server.isFailure()) {
      return new Failure(
        new Error("failed to find post author's server data", server.value),
      );
    }

    return new Success(
      this.convertToCommonResponse({
        post: res.value,
        user: user.value,
        server: server.value,
      }),
    );
  }

  async FindByHandler(
    handle: string,
  ): AsyncResult<Array<CommonPostResponse>, Error> {
    // ToDo: Implement
    return new Success(Array<CommonPostResponse>());
  }

  async CreatePost(req: CommonPostRequest) {
    const res = await this.createPostService.Handle({
      attachments: req.attachments as Array<Snowflake>,
      // ToDo: 投稿ユーザーのIDを入れる
      authorID: "123" as Snowflake,
      text: req.text,
      visibility: req.visibility,
    });
    if (res.isFailure()) {
      return new Failure(new Error("failed to create post", res.value));
    }
    const user = await this.findUserService.FindByID(res.value.authorID);
    if (user.isFailure()) {
      return new Failure(
        new Error("failed to find post author's data", user.value),
      );
    }
    const server = await this.findServerService.FindByID(user.value.serverID);
    if (server.isFailure()) {
      return new Failure(
        new Error("failed to find post author's server data", server.value),
      );
    }

    return new Success(
      this.convertToCommonResponse({
        post: res.value,
        server: server.value,
        user: user.value,
      }),
    );
  }

  private convertToCommonResponse(arg: {
    post: PostData;
    user: UserData;
    server: ServerData;
  }) {
    const resp: CommonPostResponse = {
      id: arg.post.id,
      text: arg.post.text,
      author: {
        id: arg.post.authorID,
        nickName: arg.user.nickName,
        host: `@${arg.user.handle}@${NormalizeURL(arg.server.host).host}`,
        iconImageURL: arg.user.iconImageURL,
      },
      createdAt: arg.post.createdAt,
      attachments: arg.post.attachments.map((v): CommonMediaResponse => {
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
      reactions: arg.post.reactions.map((v): PostReactionResponse => {
        return { postID: v.postID, userID: v.userID };
      }),
    };
    return resp;
  }
}
