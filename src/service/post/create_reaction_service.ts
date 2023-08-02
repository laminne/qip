import { Snowflake, SnowflakeIDGenerator } from "../../helpers/id_generator.js";
import { IReactionRepository } from "../../repository/reaction.js";
import { AsyncResult } from "../../helpers/result.js";
import { PostReactionEvent } from "../../domain/post.js";

export class CreateReactionService {
  private readonly repository: IReactionRepository;
  private readonly idGenerator: SnowflakeIDGenerator;
  constructor(args: {
    repository: IReactionRepository;
    idGenerator: SnowflakeIDGenerator;
  }) {
    this.repository = args.repository;
    this.idGenerator = args.idGenerator;
  }

  async Handle(
    postID: Snowflake,
    userID: Snowflake,
  ): AsyncResult<PostReactionEvent, Error> {
    throw new Error("todo");
  }
}
