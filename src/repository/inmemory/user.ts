/* eslint-disable @typescript-eslint/no-explicit-any */
import { IUserRepository } from "../user.js";
import { AsyncResult, Failure, Result, Success } from "../../helpers/result.js";
import { User } from "../../domain/user.js";
import { Snowflake } from "../../helpers/id_generator.js";

export class UserRepository implements IUserRepository {
  private data: Set<User>;

  constructor(data: User[]) {
    this.data = new Set(data);
  }

  async Create(u: User): Promise<Result<User, Error>> {
    try {
      this.data.add(u);

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

  async Update(): Promise<Result<User, Error>> {
    return new Failure(new Error(""));
  }

  // 指定したユーザーがフォローしているユーザー一覧を取得
  async FindFollowing(): AsyncResult<Array<User>, Error> {
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
  async FindFollower(): AsyncResult<Array<User>, Error> {
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
