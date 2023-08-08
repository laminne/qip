import { IServerRepository } from "../server.js";
import { Server } from "../../domain/server.js";
import { Failure, Result, Success } from "../../helpers/result.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { PrismaClient } from "@prisma/client";
import { PrismaErrorConverter } from "./error.js";

export class ServerRepository implements IServerRepository {
  private prisma: PrismaClient;

  constructor(prisma: PrismaClient) {
    this.prisma = prisma;
  }

  async Create(s: Server): Promise<Result<Server, Error>> {
    try {
      const res = await this.prisma.server.create({
        data: {
          id: s.id,
          host: s.host,
          softwareName: s.softwareName,
          softwareVersion: s.softwareVersion,
          name: s.name,
          description: s.description,
          maintainer: s.maintainer,
          maintainerEmail: s.maintainerEmail,
          iconURL: s.iconURL,
          faviconURL: s.faviconURL,
        },
      });

      return new Success(this.convertToDomain(res as ServerEntity));
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async FindByHost(host: string): Promise<Result<Server, Error>> {
    try {
      const res = await this.prisma.server.findUnique({
        where: {
          host: host,
        },
      });

      return new Success(this.convertToDomain(res as ServerEntity));
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async FindByID(id: Snowflake): Promise<Result<Server, Error>> {
    try {
      const res = await this.prisma.server.findUnique({
        where: {
          id: id,
        },
      });

      return new Success(this.convertToDomain(res as ServerEntity));
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async Update(): Promise<Result<Server, Error>> {
    return new Failure(new Error(""));
  }

  private convertToDomain(i: ServerEntity): Server {
    return new Server({
      id: i.id as Snowflake,
      description: i.description,
      faviconURL: i.faviconURL,
      host: i.host,
      iconURL: i.iconURL,
      maintainer: i.maintainer,
      maintainerEmail: i.maintainerEmail,
      name: i.name,
      softwareName: i.softwareName,
      softwareVersion: i.softwareVersion,
    });
  }
}

export type ServerEntity = {
  id: string;
  description: string;
  faviconURL: string;
  host: string;
  iconURL: string;
  maintainer: string;
  maintainerEmail: string;
  name: string;
  softwareName: string;
  softwareVersion: string;
};
