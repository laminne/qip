import { IReactionRepository } from "../reaction.js";
import { PrismaClient } from "@prisma/client";
import { PostReactionEvent } from "../../domain/post.js";
import { AsyncResult, Failure, Success } from "../../helpers/result.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { PrismaErrorConverter } from "./error.js";

export class ReactionRepository implements IReactionRepository {
  private readonly prisma: PrismaClient;

  constructor(prisma: PrismaClient) {
    this.prisma = prisma;
  }

  async Create(d: PostReactionEvent): AsyncResult<PostReactionEvent, Error> {
    try {
      const res = await this.prisma.reaction.create({
        data: {
          User: {
            connect: {
              id: d.userID,
            },
          },
          Post: {
            connect: {
              id: d.postID,
            },
          },
        },
      });

      return new Success(this.toDomain(res as PostReactionEntity));
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async Find(
    postID: Snowflake,
    userID: Snowflake,
  ): AsyncResult<PostReactionEvent, Error> {
    try {
      const res = await this.prisma.reaction.findUnique({
        where: {
          userId_postId: {
            userId: userID,
            postId: postID,
          },
        },
      });

      return new Success(this.toDomain(res as PostReactionEntity));
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async Undo(postID: Snowflake, userID: Snowflake): AsyncResult<void, Error> {
    try {
      await this.prisma.reaction.delete({
        where: {
          userId_postId: {
            userId: userID,
            postId: postID,
          },
        },
      });
      return new Success(void "");
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  private toDomain(v: PostReactionEntity) {
    return new PostReactionEvent(v.postId as Snowflake, v.userId as Snowflake);
  }
}

export type PostReactionEntity = { postId: string; userId: string };
