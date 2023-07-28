import { User } from "../domain/user";
import { AsyncResult, Result } from "../helpers/result";
import { Snowflake } from "../helpers/id_generator";

export interface IUserRepository {
  Create(u: User): AsyncResult<User, Error>;
  Update(u: User): AsyncResult<User, Error>;

  FindByID(id: Snowflake): AsyncResult<User, Error>;
  FindByHandle(handle: string): AsyncResult<Array<User>, Error>;
  FindFollowing(id: Snowflake): AsyncResult<Array<User>, Error>;
  FindFollower(id: Snowflake): AsyncResult<Array<User>, Error>;
}
