import { IPostRepository } from "../post.js";
import { AsyncResult, Failure, Result, Success } from "../../helpers/result.js";
import { Post } from "../../domain/post.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { User } from "../../domain/user.js";

export class PostRepository implements IPostRepository {
  private data: Set<Post>;

  constructor(data: Post[]) {
    this.data = new Set(data);
  }

  async Create(p: Post): Promise<Result<Post, Error>> {
    try {
      this.data.add(p);
      return new Success(p);
    } catch (e: unknown) {
      console.log(e);
      return new Failure(new Error(e as any));
    }
  }

  async FindByAuthor(id: Snowflake): Promise<Result<Array<Post>, Error>> {
    try {
      const res = [...this.data].filter((v) => v.authorID === id);

      return new Success(res);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByID(id: Snowflake): Promise<Result<Post, Error>> {
    try {
      const res = [...this.data].filter((v) => v.id == id);

      return new Success(res[0]);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async Update(p: Post): Promise<Result<Post, Error>> {
    return new Failure(new Error(""));
  }

  async Delete(id: Snowflake): AsyncResult<void, Error> {
    return new Failure(new Error("todo"));
  }

  async ChronologicalPosts(
    userID: Snowflake,
    cursor: number,
  ): AsyncResult<{ posts: Post; author: User }[], Error> {
    try {
      // ToDo: 自分がフォローしているユーザーの投稿の取得
      [...this.data].filter((v) => {
        return v.id;
      });

      return new Success(new Array<{ posts: Post; author: User }>());
    } catch (e: unknown) {
      return new Failure(new Error(e as Error as any));
    }
  }
}
