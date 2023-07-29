import { IPostRepository } from "../post";
import { AsyncResult, Failure, Result, Success } from "../../helpers/result";
import { Post } from "../../domain/post";
import { Snowflake } from "../../helpers/id_generator";
import { User } from "../../domain/user";

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

  async ChronologicalPosts(
    userID: Snowflake,
    cursor: number,
  ): AsyncResult<{ posts: Post; author: User }[], Error> {
    try {
      // ToDo: 自分がフォローしているユーザーの投稿の取得
      const posts = [...this.data].filter((v) => {
        return v.id;
      });

      return new Success(new Array<{ posts: Post; author: User }>());
    } catch (e: unknown) {
      return new Failure(new Error(e as Error as any));
    }
  }
}
