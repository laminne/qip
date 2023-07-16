import { User } from "../domain/user";
import { Result } from "../helpers/result";
import { Snowflake } from "../helpers/id_generator";

export interface IUserRepository {
  Create(u: User): Promise<Result<User, Error>>;
  Update(u: User): Promise<Result<User, Error>>;

  FindByID(id: Snowflake): Promise<Result<User, Error>>;
  FindByHandle(handle: string): Promise<Result<Array<User>, Error>>;
}
