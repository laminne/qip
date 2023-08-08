import { IMediaRepository } from "../media.js";
import { Media } from "../../domain/media.js";
import { Failure, Result, Success } from "../../helpers/result.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { PrismaClient } from "@prisma/client";
import { PrismaErrorConverter } from "./error.js";

export class MediaRepository implements IMediaRepository {
  private prisma: PrismaClient;

  constructor(prisma: PrismaClient) {
    this.prisma = prisma;
  }

  async Create(m: Media): Promise<Result<Media, Error>> {
    try {
      const res = await this.prisma.media.create({
        data: {
          id: m.id,
          postID: m.postID,
          authorID: m.authorID,
          name: m.name,
          type: m.type,
          md5Sum: m.md5Sum,
          size: m.size,
          isSensitive: m.isSensitive,
          blurhash: m.blurhash,
          url: m.url,
          thumbnailURL: m.thumbnailURL,
          cached: m.cached,
        },
      });
      return new Success(this.convertToDomain(res));
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async FindByID(id: Snowflake): Promise<Result<Media, Error>> {
    try {
      const res = await this.prisma.media.findUnique({
        where: {
          id: id,
        },
      });
      return new Success(this.convertToDomain(res as MediaEntity));
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async FindByPostID(id: Snowflake): Promise<Result<Array<Media>, Error>> {
    try {
      const res = await this.prisma.media.findMany({
        where: {
          postID: id,
        },
      });
      return new Success(
        (res as MediaEntity[]).map((v) => this.convertToDomain(v)),
      );
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async FindByUserID(id: Snowflake): Promise<Result<Array<Media>, Error>> {
    try {
      const res = await this.prisma.media.findMany({
        where: {
          authorID: id,
        },
      });
      return new Success(
        (res as MediaEntity[]).map((v) => this.convertToDomain(v)),
      );
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async Update(): Promise<Result<Media, Error>> {
    return new Failure(new Error(""));
  }

  private convertToDomain(v: MediaEntity): Media {
    return new Media({
      id: v.id as Snowflake,
      authorID: v.authorID as Snowflake,
      postID: v.postID as Snowflake,
      blurhash: v.blurhash,
      cached: v.cached,
      isSensitive: v.isSensitive,
      md5Sum: v.md5Sum,
      name: v.name,
      size: v.size,
      thumbnailURL: v.thumbnailURL,
      type: v.type,
      url: v.url,
    });
  }
}

export type MediaEntity = {
  id: string;
  authorID: string;
  postID: string;
  blurhash: string;
  cached: boolean;
  isSensitive: boolean;
  md5Sum: string;
  name: string;
  size: number;
  thumbnailURL: string;
  type: string;
  url: string;
};
