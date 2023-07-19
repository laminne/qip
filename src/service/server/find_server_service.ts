import { Snowflake } from "../../helpers/id_generator";
import { Failure, Success } from "../../helpers/result";
import { UserToUserData } from "../data/user";
import { IServerRepository } from "../../repository/server";
import { ServerToServerData } from "../data/server";

export class FindServerService {
  private repository: IServerRepository;

  constructor(repository: IServerRepository) {
    this.repository = repository;
  }

  async FindByID(id: Snowflake) {
    const res = await this.repository.FindByID(id);
    if (res.isFailure()) {
      return new Failure(new Error("failed to find user by id", res.value));
    }
    return new Success(ServerToServerData(res.value));
  }

  async FindByHost(host: string) {
    const res = await this.repository.FindByHost(host);
    if (res.isFailure()) {
      return new Failure(new Error("failed to find user by id", res.value));
    }
    return new Success(ServerToServerData(res.value));
  }
}
