import { IPostRepository } from "../../repository/post";
import { IUserRepository } from "../../repository/user";
import { Snowflake } from "../../helpers/id_generator";
import { Failure, Success } from "../../helpers/result";
import { UserToUserData } from "../data/user";
import { PostToPostData } from "../data/post";

export class CreateTimelineService {
  private readonly postRepository: IPostRepository;
  private readonly userRepository: IUserRepository;
  constructor(args: {
    postRepository: IPostRepository;
    userRepository: IUserRepository;
  }) {
    this.postRepository = args.postRepository;
    this.userRepository = args.userRepository;
  }

  async Handle(userID: Snowflake) {
    // 投稿を取得
    const posts = await this.postRepository.ChronologicalPosts(userID, 0);
    if (posts.isFailure()) {
      return new Failure(
        new Error("find to find following users", posts.value),
      );
    }

    return new Success(
      posts.value.map((v) => {
        return {
          author: UserToUserData(v.author),
          posts: PostToPostData(v.posts),
        };
      }),
    );
  }
}
