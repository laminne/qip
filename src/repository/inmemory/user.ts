import { IUserRepository } from "../user";
import { AsyncResult, Failure, Result, Success } from "../../helpers/result";
import { User } from "../../domain/user";
import { Snowflake } from "../../helpers/id_generator";

export class UserRepository implements IUserRepository {
  private data: Set<User>;

  constructor(data: User[]) {
    this.data = new Set(data);
  }

  async Create(u: User): Promise<Result<User, Error>> {
    try {
      const res = this.data.add(u);

      return new Success(u);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByHandle(handle: string): Promise<Result<User, Error>> {
    try {
      const res = [...this.data].filter((v) => v.fullHandle === handle);
      return new Success(res[0]);
    } catch (e: unknown) {
      return new Failure(new Error(e as Error as any));
    }
  }

  async FindByID(id: Snowflake): Promise<Result<User, Error>> {
    try {
      const res = [...this.data].filter((v) => v.id === id);
      return new Success(res[0]);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async Update(u: User): Promise<Result<User, Error>> {
    return new Failure(new Error(""));
  }

  // 指定したユーザーがフォローしているユーザー一覧を取得
  async FindFollowing(id: Snowflake): AsyncResult<Array<User>, Error> {
    try {
      // ToDo: 実装する
      return new Success(new Array<User>());
    } catch (e: unknown) {
      return new Failure(
        new Error("failed to find user follow", e as Error as any),
      );
    }
  }

  // 指定したユーザーをフォローしているユーザー一覧を取得
  async FindFollower(id: Snowflake): AsyncResult<Array<User>, Error> {
    try {
      // ToDo: 実装する
      return new Success(new Array<User>());
    } catch (e: unknown) {
      return new Failure(
        new Error("failed to find user follow", e as Error as any),
      );
    }
  }
}
