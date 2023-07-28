import { IUserRepository } from "../user";
import { AsyncResult, Failure, Result, Success } from "../../helpers/result";
import { User, UserAPData } from "../../domain/user";
import { Snowflake } from "../../helpers/id_generator";
import { PrismaClient } from "@prisma/client";

export class UserRepository implements IUserRepository {
  private prisma: PrismaClient;

  constructor(prisma: any) {
    this.prisma = prisma;
  }

  async Create(u: User): Promise<Result<User, Error>> {
    try {
      const res = await this.prisma.user.create({
        data: {
          id: u.id,
          handle: u.handle,
          nickName: u.nickName,
          role: 0,
          bio: u.bio,
          headerImageURL: u.headerImageURL,
          iconImageURL: u.iconImageURL,
          password: u.password ?? "",
          isLocalUser: u.isLocalUser,
          Server: {
            connect: {
              id: u.serverID,
            },
          },
          userAPData: {
            create: {
              id: u.apData.userAPID,
              inboxURL: u.apData.inboxURL,
              outboxURL: u.apData.outboxURL,
              followersURL: u.apData.followersURL,
              followingURL: u.apData.followingURL,
              publicKey: u.apData.publicKey,
              privateKey: u.apData.privateKey,
            },
          },
        },
        include: {
          userAPData: true,
          following: true,
          follower: true,
        },
      });

      const response: Array<User> = this.convertToDomain(new Array<any>(res));
      return new Success(response[0]);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByHandle(handle: string): Promise<Result<Array<User>, Error>> {
    try {
      const res = await this.prisma.user.findMany({
        where: {
          handle: handle,
        },
        include: {
          userAPData: true,
          following: true,
          follower: true,
        },
      });
      return new Success(this.convertToDomain(res));
    } catch (e: unknown) {
      return new Failure(new Error(e as Error as any));
    }
  }

  async FindByID(id: Snowflake): Promise<Result<User, Error>> {
    try {
      const res = await this.prisma.user.findMany({
        where: {
          id: id,
        },
        include: {
          userAPData: true,
          following: true,
          follower: true,
        },
      });
      return new Success(this.convertToDomain(res)[0]);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async Update(u: User): Promise<Result<User, Error>> {
    return new Failure(new Error(""));
  }

  // 指定したユーザーがフォローしているユーザー一覧を取得
  async FindFollowing(id: Snowflake): AsyncResult<Array<User>, Error> {
    try {
      const res = await this.prisma.userFollowEvent.findMany({
        where: {
          following: {
            id: id,
          },
        },
        include: {
          follower: true,
        },
      });
      return new Success(this.convertToFollowDomain(res));
    } catch (e: unknown) {
      return new Failure(
        new Error("failed to find user follow", e as Error as any),
      );
    }
  }

  // 指定したユーザーをフォローしているユーザー一覧を取得
  async FindFollower(id: Snowflake): AsyncResult<Array<User>, Error> {
    try {
      const res = await this.prisma.userFollowEvent.findMany({
        where: {
          following: {
            id: id,
          },
        },
        include: {
          follower: true,
        },
      });
      return new Success(this.convertToFollowerDomain(res));
    } catch (e: unknown) {
      return new Failure(
        new Error("failed to find user follow", e as Error as any),
      );
    }
  }

  private convertToDomain(ew: Array<any>): Array<User> {
    return ew.map((e) => {
      return new User({
        id: e.id as Snowflake,
        bio: e.bio,
        createdAt: e.createdAt,
        handle: e.handle,
        headerImageURL: e.headerImageURL,
        iconImageURL: e.iconImageURL,
        isLocalUser: e.isLocalUser,
        nickName: e.nickName,
        password: e.password,
        role: e.role,
        following: e.following,
        apData: new UserAPData({
          followersURL: e.userAPData.followersURL ?? "",
          followingURL: e.userAPData.followingURL ?? "",
          inboxURL: e.userAPData.inboxURL ?? "",
          outboxURL: e.userAPData.outboxURL ?? "",
          privateKey: e.userAPData.privateKey ?? "",
          publicKey: e.userAPData.publicKey ?? "",
          userAPID: e.userAPData.id ?? "",
          userID: (e.id as Snowflake) ?? "",
        }),
        serverID: e.serverId as Snowflake,
      });
    });
  }

  private convertToFollowDomain(i: Array<any>): Array<User> {
    return i.map((e) => {
      return new User({
        id: e.follower.id as Snowflake,
        bio: e.follower.bio,
        createdAt: e.follower.createdAt,
        handle: e.follower.handle,
        headerImageURL: e.follower.headerImageURL,
        iconImageURL: e.follower.iconImageURL,
        isLocalUser: e.follower.isLocalUser,
        nickName: e.follower.nickName,
        password: e.follower.password,
        role: e.follower.role,
        following: e.follower.following,
        apData: new UserAPData({
          followersURL: "",
          followingURL: "",
          inboxURL: "",
          outboxURL: "",
          privateKey: "",
          publicKey: "",
          userAPID: "",
          userID: e.follower.id as Snowflake,
        }),
        serverID: e.follower.serverId as Snowflake,
      });
    });
  }

  private convertToFollowerDomain(i: Array<any>): Array<User> {
    return i.map((e) => {
      return new User({
        id: e.following.id as Snowflake,
        bio: e.following.bio,
        createdAt: e.following.createdAt,
        handle: e.following.handle,
        headerImageURL: e.following.headerImageURL,
        iconImageURL: e.following.iconImageURL,
        isLocalUser: e.following.isLocalUser,
        nickName: e.following.nickName,
        password: e.following.password,
        role: e.following.role,
        following: e.following.following,
        apData: new UserAPData({
          followersURL: "",
          followingURL: "",
          inboxURL: "",
          outboxURL: "",
          privateKey: "",
          publicKey: "",
          userAPID: "",
          userID: e.following.id as Snowflake,
        }),
        serverID: e.following.serverId as Snowflake,
      });
    });
  }
}
