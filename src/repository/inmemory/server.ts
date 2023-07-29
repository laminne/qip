import { IServerRepository } from "../server";
import { Server } from "../../domain/server";
import { Failure, Result, Success } from "../../helpers/result";
import { Snowflake } from "../../helpers/id_generator";

export class ServerRepository implements IServerRepository {
  private data: Set<Server>;

  constructor(data: Server[]) {
    this.data = new Set(data);
  }

  async Create(s: Server): Promise<Result<Server, Error>> {
    try {
      this.data.add(s);

      return new Success(s);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByHost(host: string): Promise<Result<Server, Error>> {
    try {
      return new Success([...this.data].filter((v) => v.host === host)[0]);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByID(id: Snowflake): Promise<Result<Server, Error>> {
    try {
      return new Success([...this.data].filter((v) => v.id === id)[0]);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async Update(s: Server): Promise<Result<Server, Error>> {
    return new Failure(new Error(""));
  }
}
