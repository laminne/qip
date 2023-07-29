import { IPostRepository } from "../../repository/post.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { AsyncResult, Failure, Success } from "../../helpers/result.js";
import { UserData, UserToUserData } from "../data/user.js";
import { PostData, PostToPostData } from "../data/post.js";
import { User } from "../../domain/user.js";
import { Post } from "../../domain/post.js";

export class CreateTimelineService {
  private readonly postRepository: IPostRepository;
  constructor(args: { postRepository: IPostRepository }) {
    this.postRepository = args.postRepository;
  }

  async Handle(
    userID: Snowflake,
  ): AsyncResult<Array<{ posts: PostData; author: UserData }>, Error> {
    // 投稿を取得
    const posts = await this.postRepository.ChronologicalPosts(userID, 0);
    if (posts.isFailure()) {
      return new Failure(
        new Error("find to find following users", posts.value),
      );
    }

    return new Success(
      posts.value.map((v: { posts: Post; author: User }) => {
        return {
          author: UserToUserData(v.author),
          posts: PostToPostData(v.posts),
        };
      }),
    );
  }
}
