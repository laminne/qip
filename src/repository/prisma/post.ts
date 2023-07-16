import { IPostRepository } from "../post";
import { Failure, Result, Success } from "../../helpers/result";
import { Post, PostReactionEvent } from "../../domain/post";
import { Snowflake } from "../../helpers/id_generator";
import { Media } from "../../domain/media";
import { PrismaClient } from "@prisma/client";

export class PostRepository implements IPostRepository {
  private prisma: PrismaClient;

  constructor(prisma: PrismaClient) {
    this.prisma = prisma;
  }

  async Create(p: Post): Promise<Result<Post, Error>> {
    try {
      const attachments = p.attachments.map((v) => {
        return {
          id: v.id,
          name: v.name,
          type: v.type,
          md5Sum: v.md5Sum,
          size: v.size,
          isSensitive: v.isSensitive,
          blurhash: v.blurhash,
          url: v.url,
          thumbnailURL: v.thumbnailURL,
          cached: v.cached,
          authorID: v.authorID,
          User: {
            connect: {
              id: v.authorID,
            },
          },
          Post: {
            connect: {
              id: v.postID,
            },
          },
        };
      });

      const res = await this.prisma.post.create({
        data: {
          User: {
            connect: {
              id: p.authorID,
            },
          },
          id: p.id,
          text: p.text,
          visibility: p.visibility,
          attachments: {
            createMany: {
              data: attachments,
            },
          },
        },
        include: {
          attachments: true,
        },
      });
      return new Success(this.convertToDomain(res));
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByAuthor(id: Snowflake): Promise<Result<Array<Post>, Error>> {
    try {
      const res = await this.prisma.post.findMany({
        where: {
          authorID: id,
        },
        include: {
          attachments: true,
          reactions: true,
        },
      });

      return new Success(res.map((r) => this.convertToDomain(r)));
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByID(id: Snowflake): Promise<Result<Post, Error>> {
    try {
      const res = await this.prisma.post.findUnique({
        where: {
          id: id,
        },
        include: {
          attachments: true,
          reactions: true,
        },
      });

      return new Success(this.convertToDomain(res));
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async Update(p: Post): Promise<Result<Post, Error>> {
    return new Failure(new Error(""));
  }

  private convertToDomain(i: any): Post {
    try {
      return new Post({
        id: i.id as Snowflake,
        authorID: i.authorID,
        createdAt: i.createdAt,
        text: i.text,
        visibility: i.visibility,
        reactions: i.reactions.map((v: any) => {
          return new PostReactionEvent(
            v.postId as Snowflake,
            v.userId as Snowflake,
          );
        }),
        attachments: i.attachments.map((v: any) => {
          return new Media({
            authorID: v.authorID,
            blurhash: v.blurhash,
            cached: v.cached,
            id: v.id,
            isSensitive: v.isSensitive,
            md5Sum: v.md5Sum,
            name: v.name,
            postID: v.postID,
            size: v.size,
            thumbnailURL: v.thumbnailURL,
            type: v.type,
            url: v.url,
          });
        }),
      });
    } catch (e: unknown) {
      throw new Error(e as any);
    }
  }
}
