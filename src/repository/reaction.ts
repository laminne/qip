import { Snowflake } from "../helpers/id_generator.js";
import { AsyncResult } from "../helpers/result.js";
import { PostReactionEvent } from "../domain/post.js";

export interface IReactionRepository {
  Create(d: PostReactionEvent): AsyncResult<PostReactionEvent, Error>;
  Undo(postID: Snowflake, userID: Snowflake): AsyncResult<void, Error>;
  Find(
    postID: Snowflake,
    userID: Snowflake,
  ): AsyncResult<PostReactionEvent, Error>;
}
