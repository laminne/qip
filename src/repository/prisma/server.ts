import { IServerRepository } from "../server";
import { Server } from "../../domain/server";
import { Failure, Result, Success } from "../../helpers/result";
import { Snowflake } from "../../helpers/id_generator";
import { PrismaClient } from "@prisma/client";

export class ServerRepository implements IServerRepository {
  private prisma: PrismaClient;

  constructor(prisma: any) {
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

      return new Success(this.convertToDomain(res));
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByHost(host: string): Promise<Result<Server, Error>> {
    try {
      const res = await this.prisma.server.findUnique({
        where: {
          host: host,
        },
      });

      return new Success(this.convertToDomain(res));
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByID(id: Snowflake): Promise<Result<Server, Error>> {
    try {
      const res = await this.prisma.server.findUnique({
        where: {
          id: id,
        },
      });

      return new Success(this.convertToDomain(res));
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async Update(s: Server): Promise<Result<Server, Error>> {
    return new Failure(new Error(""));
  }

  private convertToDomain(i: any): Server {
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
