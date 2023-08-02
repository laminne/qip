import { IPostRepository } from "../post.js";
import { AsyncResult, Failure, Result, Success } from "../../helpers/result.js";
import { Post, PostReactionEvent } from "../../domain/post.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { Media } from "../../domain/media.js";
import { PrismaClient } from "@prisma/client";
import { User, UserAPData, UserFollowEvent } from "../../domain/user.js";

export class PostRepository implements IPostRepository {
  private prisma: PrismaClient;

  constructor(prisma: PrismaClient) {
    this.prisma = prisma;
  }

  async Create(p: Post): Promise<Result<Post, Error>> {
    try {
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
            connect: p.attachments.map((v) => {
              return {
                id: v.id,
              };
            }),
          },
        },
        include: {
          attachments: true,
        },
      });
      return new Success(this.convertToDomain(res));
    } catch (e: unknown) {
      console.log(e);
      return new Failure(new Error(e as any));
    }
  }

  async Delete(id: Snowflake): AsyncResult<void, Error> {
    try {
      await this.prisma.post.update({
        where: {
          id: id,
        },
        data: {
          deletedAt: new Date(),
        },
      });
      return new Success(void "");
    } catch (e: unknown) {
      return new Failure(new Error("failed to delete post", e as Error as any));
    }
  }

  async FindByAuthor(id: Snowflake): Promise<Result<Array<Post>, Error>> {
    try {
      const res = await this.prisma.post.findMany({
        where: {
          authorID: id,
          deletedAt: null,
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
          deletedAt: null,
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

  // 時系列順にフォローしているユーザーと自分自身の投稿を取得
  async ChronologicalPosts(
    userID: Snowflake,
    cursor: number,
  ): AsyncResult<{ posts: Post; author: User }[], Error> {
    try {
      const posts = await this.prisma.post.findMany({
        where: {
          OR: [
            {
              User: {
                follower: {
                  some: {
                    followingID: userID,
                  },
                },
              },
              attachments: {
                some: {
                  deletedAt: null,
                },
              },
            },
            {
              authorID: userID,
            },
          ],
          deletedAt: null,
        },
        orderBy: {
          createdAt: "desc",
        },
        include: {
          User: true,
          attachments: true,
          reactions: true,
        },
      });

      return new Success(
        posts.map((p) => {
          return {
            posts: new Post({
              attachments: !p.attachments
                ? new Array<Media>()
                : p.attachments.map((v: any) => {
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
              authorID: p.authorID as Snowflake,
              createdAt: p.createdAt,
              id: p.id as Snowflake,
              reactions: !p.reactions
                ? new Array<PostReactionEvent>()
                : p.reactions.map((v: any) => {
                    return new PostReactionEvent(
                      v.postId as Snowflake,
                      v.userId as Snowflake,
                    );
                  }),
              text: p.text,
              visibility: 0,
            }),
            author: new User({
              bio: p.User.bio,
              apData: new UserAPData({
                followersURL: "",
                followingURL: "",
                inboxURL: "",
                outboxURL: "",
                privateKey: "",
                publicKey: "",
                userAPID: "",
                userID: "" as Snowflake,
              }),
              createdAt: p.User.createdAt,
              following: new Array<UserFollowEvent>(),
              handle: p.User.handle,
              fullHandle: p.User.fullHandle,
              headerImageURL: p.User.headerImageURL,
              iconImageURL: p.User.iconImageURL,
              id: p.authorID as Snowflake,
              isLocalUser: p.User.isLocalUser,
              nickName: p.User.nickName,
              password: "",
              role: p.User.role,
              serverID: p.User.serverId as Snowflake,
            }),
          };
        }),
      );
    } catch (e: unknown) {
      return new Failure(new Error(e as Error as any));
    }
  }

  private convertToDomain(i: any): Post {
    try {
      return new Post({
        id: i.id as Snowflake,
        authorID: i.authorID,
        createdAt: i.createdAt,
        text: i.text,
        visibility: i.visibility,
        reactions: !i.reactions
          ? new Array<PostReactionEvent>()
          : i.reactions.map((v: any) => {
              return new PostReactionEvent(
                v.postId as Snowflake,
                v.userId as Snowflake,
              );
            }),
        attachments: !i.attachments
          ? new Array<Media>()
          : i.attachments.map((v: any) => {
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
      console.log(i.reactions, i.attachments);
      throw new Error(e as any);
    }
  }
}
