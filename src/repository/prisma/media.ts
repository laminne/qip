import { IMediaRepository } from "../media";
import { Media } from "../../domain/media";
import { Failure, Result, Success } from "../../helpers/result";
import { Snowflake } from "../../helpers/id_generator";
import { PrismaClient } from "@prisma/client";

export class MediaRepository implements IMediaRepository {
  private prisma: PrismaClient;

  constructor(prisma: PrismaClient) {
    this.prisma = prisma;
  }

  // Mediaは投稿作成時にしか生成されないのでここでは書かない

  async FindByID(id: Snowflake): Promise<Result<Media, Error>> {
    try {
      const res = await this.prisma.media.findUnique({
        where: {
          id: id,
        },
      });
      return new Success(this.convertToDomain(res));
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByPostID(id: Snowflake): Promise<Result<Array<Media>, Error>> {
    try {
      const res = await this.prisma.media.findMany({
        where: {
          postID: id,
        },
      });
      return new Success(res.map((v: any) => this.convertToDomain(v)));
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByUserID(id: Snowflake): Promise<Result<Array<Media>, Error>> {
    try {
      const res = await this.prisma.media.findMany({
        where: {
          authorID: id,
        },
      });
      return new Success(res.map((v: any) => this.convertToDomain(v)));
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async Update(m: Media): Promise<Result<Media, Error>> {
    return new Failure(new Error(""));
  }

  private convertToDomain(v: any): Media {
    return new Media({
      id: v.id,
      authorID: v.authorID,
      postID: v.postID,
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
