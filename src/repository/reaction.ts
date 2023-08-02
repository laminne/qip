import { Snowflake } from "../helpers/id_generator.js";
import { AsyncResult } from "../helpers/result.js";
import { PostReactionEvent } from "../domain/post.js";

export interface IReactionRepository {
  Create(
    postID: Snowflake,
    userID: Snowflake,
  ): AsyncResult<PostReactionEvent, Error>;
  Undo(id: Snowflake): AsyncResult<void, Error>;
}
