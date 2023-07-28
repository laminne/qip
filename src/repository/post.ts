import { AsyncResult, Result } from "../helpers/result";
import { Post } from "../domain/post";
import { Snowflake } from "../helpers/id_generator";
import { User } from "../domain/user";

export interface IPostRepository {
  Create(p: Post): Promise<Result<Post, Error>>;
  Update(p: Post): Promise<Result<Post, Error>>;

  FindByID(id: Snowflake): Promise<Result<Post, Error>>;
  FindByAuthor(id: Snowflake): Promise<Result<Array<Post>, Error>>;

  ChronologicalPosts(
    authorIDs: Snowflake,
    cursor: number,
  ): AsyncResult<{ posts: Post; author: User }[], Error>;
}
