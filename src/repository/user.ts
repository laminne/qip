import { User, UserFollowEvent } from "../domain/user.js";
import { AsyncResult } from "../helpers/result.js";
import { Snowflake } from "../helpers/id_generator.js";

export interface IUserRepository {
  Create(u: User): AsyncResult<User, Error>;
  Update(u: User): AsyncResult<User, Error>;
  CreateFollow(u: UserFollowEvent): AsyncResult<UserFollowEvent, Error>;

  FindByID(id: Snowflake): AsyncResult<User, Error>;
  FindByHandle(handle: string): AsyncResult<User, Error>;
  FindFollowing(id: Snowflake): AsyncResult<Array<User>, Error>;
  FindFollower(id: Snowflake): AsyncResult<Array<User>, Error>;
  FindFollowEvent(
    followingID: Snowflake,
    followerID: Snowflake,
  ): AsyncResult<UserFollowEvent, Error>;
}
