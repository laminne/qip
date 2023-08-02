import { AsyncResult, Failure, Success } from "../../helpers/result.js";
import {
  CommonMediaResponse,
  CommonPostRequest,
  CommonPostResponse,
  PostReactionResponse,
} from "../types/post.js";
import { FindPostService } from "../../service/post/find_post_service.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { FindUserService } from "../../service/user/find_user_service.js";
import { CreatePostService } from "../../service/post/create_post_service.js";
import { PostData } from "../../service/data/post.js";
import { UserData } from "../../service/data/user.js";
import { CreateTimelineService } from "../../service/post/create_timeline_service.js";
import { DeletePostService } from "../../service/post/delete_post_service.js";
import { CreateReactionService } from "../../service/post/create_reaction_service.js";

export class PostController {
  private readonly findPostService: FindPostService;
  private readonly findUserService: FindUserService;
  private readonly createPostService: CreatePostService;
  private readonly createTimelineService: CreateTimelineService;
  private readonly deletePostService: DeletePostService;
  private readonly createReactionService: CreateReactionService;

  constructor(args: {
    findPostService: FindPostService;
    findUserService: FindUserService;
    createPostService: CreatePostService;
    createTimelineService: CreateTimelineService;
    deletePostService: DeletePostService;
    createReactionService: CreateReactionService;
  }) {
    this.findPostService = args.findPostService;
    this.findUserService = args.findUserService;
    this.createPostService = args.createPostService;
    this.createTimelineService = args.createTimelineService;
    this.deletePostService = args.deletePostService;
    this.createReactionService = args.createReactionService;
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

    return new Success(
      this.convertToCommonResponse({
        post: res.value,
        user: user.value,
      }),
    );
  }

  async FindByHandle(
    handle: string,
  ): AsyncResult<Array<CommonPostResponse>, Error> {
    // ToDo: Implement
    return new Success(Array<CommonPostResponse>());
  }

  async ChronologicalPosts(
    id: string,
  ): AsyncResult<CommonPostResponse[], Error> {
    const res = await this.createTimelineService.Handle(id as Snowflake);
    if (res.isFailure()) {
      return new Failure(new Error("failed to create timeline", res.value));
    }

    return new Success(
      res.value.map((v) =>
        this.convertToCommonResponse({
          post: v.posts,
          user: v.author,
        }),
      ),
    );
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

    return new Success(
      this.convertToCommonResponse({
        post: res.value,
        user: user.value,
      }),
    );
  }

  async DeletePost(id: string) {
    const res = await this.deletePostService.Delete(id as Snowflake);
    if (res.isFailure()) {
      return new Failure(new Error("failed to delete post", res.value));
    }

    return new Success(res.value);
  }

  async Reaction(userID: string, postID: string) {
    const res = await this.createReactionService.Handle(
      postID as Snowflake,
      userID as Snowflake,
    );
    if (res.isFailure()) {
      return new Failure(res.value);
    }

    // ToDo: 独立した型として提供する
    const resp = {
      userID: res.value.userID,
      postID: res.value.postID,
    };
    return new Success(resp);
  }

  private convertToCommonResponse(arg: { post: PostData; user: UserData }) {
    const resp: CommonPostResponse = {
      id: arg.post.id,
      text: arg.post.text,
      author: {
        id: arg.post.authorID,
        nickName: arg.user.nickName,
        host: arg.user.fullHandle,
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
