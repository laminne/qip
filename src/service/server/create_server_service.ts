import { IServerRepository } from "../../repository/server.js";
import { Server } from "../../domain/server.js";
import { Failure, Success } from "../../helpers/result.js";
import { ServerToServerData } from "../data/server.js";
import { SnowflakeIDGenerator } from "../../helpers/id_generator.js";

export class CreateServerService {
  private readonly repository: IServerRepository;
  private readonly idGenerator: SnowflakeIDGenerator;
  constructor(
    repository: IServerRepository,
    idGenerator: SnowflakeIDGenerator,
  ) {
    this.repository = repository;
    this.idGenerator = idGenerator;
  }

  async Handle(s: CreateServerArgs) {
    const req = new Server({
      id: this.idGenerator.generate(),
      name: s.name,
      description: s.description,
      faviconURL: s.faviconURL,
      host: s.host,
      iconURL: s.iconURL,
      maintainer: s.maintainer,
      maintainerEmail: s.maintainerEmail,
      softwareName: s.softwareName,
      softwareVersion: s.softwareVersion,
    });
    const res = await this.repository.Create(req);
    if (res.isFailure()) {
      return new Failure(new Error("failed to create server", res.value));
    }

    return new Success(ServerToServerData(res.value));
  }
}

export interface CreateServerArgs {
  host: string;
  softwareName: string;
  softwareVersion: string;
  name: string;
  description: string;
  maintainer: string;
  maintainerEmail: string;
  iconURL: string;
  faviconURL: string;
}
