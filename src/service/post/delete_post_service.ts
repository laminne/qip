import { IPostRepository } from "../../repository/post.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { AsyncResult, Failure, Success } from "../../helpers/result.js";

export class DeletePostService {
  private readonly repository: IPostRepository;

  constructor(repository: IPostRepository) {
    this.repository = repository;
  }

  async Delete(id: Snowflake): AsyncResult<void, Error> {
    const res = await this.repository.Delete(id);
    if (res.isFailure()) {
      return new Failure(new Error("failed to delete post", res.value));
    }
    return new Success(void res.value);
  }
}
