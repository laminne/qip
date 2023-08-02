import { IReactionRepository } from "../reaction.js";
import { PostReactionEvent } from "../../domain/post.js";
import { AsyncResult, Failure, Success } from "../../helpers/result.js";
import { Snowflake } from "../../helpers/id_generator.js";

export class ReactionRepository implements IReactionRepository {
  private data: Set<PostReactionEvent>;

  constructor(data: PostReactionEvent[]) {
    this.data = new Set(data);
  }

  async Create(d: PostReactionEvent): AsyncResult<PostReactionEvent, Error> {
    this.data.add(d);
    return new Success(d);
  }

  async Undo(postID: Snowflake, userID: Snowflake): AsyncResult<void, Error> {
    for (const v of [...this.data]) {
      if (v.userID === userID && v.postID === postID) {
        this.data.delete(v);
        return new Success(void "");
      }
    }

    return new Failure(new Error("not found"));
  }

  async Find(
    postID: Snowflake,
    userID: Snowflake,
  ): AsyncResult<PostReactionEvent, Error> {
    for (const v of this.data) {
      if (v.userID === userID && v.postID === postID) {
        return new Success(v);
      }
    }

    return new Failure(new Error("not found"));
  }
}
