import { IPostRepository } from "../../repository/post";
import { Snowflake } from "../../helpers/id_generator";
import { Failure, Success } from "../../helpers/result";
import { PostToPostData } from "../data/post";

export class FindPostService {
  private repository: IPostRepository;

  constructor(repository: IPostRepository) {
    this.repository = repository;
  }

  async FindByID(id: Snowflake) {
    const res = await this.repository.FindByID(id);
    if (res.isFailure()) {
      return new Failure(new Error("failed to find post by id", res.value));
    }

    return new Success(PostToPostData(res.value));
  }

  async FindByAuthor(id: Snowflake) {
    const res = await this.repository.FindByAuthor(id);
    if (res.isFailure()) {
      return new Failure(
        new Error("failed to find post by authorID", res.value),
      );
    }

    return new Success(res.value.map((v) => PostToPostData(v)));
  }
}
