import { Server } from "../domain/server.js";
import { Result } from "../helpers/result.js";
import { Snowflake } from "../helpers/id_generator.js";

export interface IServerRepository {
  Create(s: Server): Promise<Result<Server, Error>>;
  Update(s: Server): Promise<Result<Server, Error>>;

  FindByID(id: Snowflake): Promise<Result<Server, Error>>;
  FindByHost(host: string): Promise<Result<Server, Error>>;
}
