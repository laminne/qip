import { IServerRepository } from "../../repository/server";
import { Server } from "../../domain/server";
import { Failure, Success } from "../../helpers/result";
import { ServerToServerData } from "../data/server";

export class CreateServerService {
  private readonly repository: IServerRepository;
  constructor(repository: IServerRepository) {
    this.repository = repository;
  }

  async Handle(s: Server) {
    const res = await this.repository.Create(s);
    if (res.isFailure()) {
      return new Failure(new Error("failed to create server", res.value));
    }

    return new Success(ServerToServerData(res.value));
  }
}
