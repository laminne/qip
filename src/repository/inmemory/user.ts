/* eslint-disable @typescript-eslint/no-explicit-any */
import { IUserRepository } from "../user.js";
import { AsyncResult, Failure, Result, Success } from "../../helpers/result.js";
import { User, UserFollowEvent } from "../../domain/user.js";
import { Snowflake } from "../../helpers/id_generator.js";

export class UserRepository implements IUserRepository {
  private data: Set<User>;
  private followData: Set<UserFollowEvent>;

  constructor(data: User[], followData?: UserFollowEvent[]) {
    this.data = new Set(data);
    this.followData = new Set(followData);
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

  async CreateFollow(u: UserFollowEvent): AsyncResult<UserFollowEvent, Error> {
    try {
      this.followData.add(u);
      return new Success(u);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindFollowEvent(
    followingID: Snowflake,
    followerID: Snowflake,
  ): AsyncResult<UserFollowEvent, Error> {
    for (const v of [...this.followData]) {
      if (v.follower.id === followerID && v.following.id === followingID) {
        return new Success(v);
      }
    }
    return new Failure(new Error("failed to find follow event"));
  }
}
