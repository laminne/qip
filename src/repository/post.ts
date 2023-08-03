import { AsyncResult, Result } from "../helpers/result.js";
import { Post } from "../domain/post.js";
import { Snowflake } from "../helpers/id_generator.js";
import { User } from "../domain/user.js";

export interface IPostRepository {
  Create(p: Post): Promise<Result<Post, Error>>;
  Update(p: Post): Promise<Result<Post, Error>>;

  FindByID(id: Snowflake): Promise<Result<Post, Error>>;
  FindByAuthor(id: Snowflake): Promise<Result<Array<Post>, Error>>;

  ChronologicalPosts(
    authorIDs: Snowflake,
    cursor: number,
  ): AsyncResult<{ posts: Post; author: User }[], Error>;

  Delete(id: Snowflake): AsyncResult<void, Error>;
}
