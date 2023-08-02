import { IReactionRepository } from "../reaction.js";
import { PostReactionEvent } from "../../domain/post.js";
import { AsyncResult } from "../../helpers/result.js";
import { Snowflake } from "../../helpers/id_generator.js";

export class ReactionRepository implements IReactionRepository {
  Create(
    postID: Snowflake,
    userID: Snowflake,
  ): AsyncResult<PostReactionEvent, Error> {
    throw new Error("todo");
  }

  Undo(id: Snowflake): AsyncResult<void, Error> {
    throw new Error("todo");
  }
}
