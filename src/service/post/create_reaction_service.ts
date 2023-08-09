import { Snowflake } from "../../helpers/id_generator.js";
import { IReactionRepository } from "../../repository/reaction.js";
import { AsyncResult, Failure, Success } from "../../helpers/result.js";
import { PostReactionEvent } from "../../domain/post.js";

export class CreateReactionService {
  private readonly repository: IReactionRepository;
  constructor(args: { repository: IReactionRepository }) {
    this.repository = args.repository;
  }

  async Handle(
    postID: Snowflake,
    userID: Snowflake,
  ): AsyncResult<PostReactionEvent, Error> {
    const data = new PostReactionEvent(postID, userID);

    if (await this.isExists(data)) {
      return new Failure(new Error("already exists"));
    }
    const res = await this.repository.Create(data);
    if (res.isFailure()) {
      return new Failure(new Error("failed to create reaction", res.value));
    }

    return new Success(data);
  }

  async Undo(postID: Snowflake, userID: Snowflake): AsyncResult<void, Error> {
    const res = await this.repository.Undo(postID, userID);
    if (res.isFailure()) {
      return new Failure(res.value);
    }
    return new Success(void "");
  }

  private async isExists(d: PostReactionEvent): Promise<boolean> {
    const res = await this.repository.Find(d.postID, d.userID);
    return res.isSuccess();
  }
}
