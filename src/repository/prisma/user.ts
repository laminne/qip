/* eslint-disable @typescript-eslint/no-explicit-any */
import { IUserRepository } from "../user.js";
import { AsyncResult, Failure, Result, Success } from "../../helpers/result.js";
import { User, UserAPData, UserFollowEvent } from "../../domain/user.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { PrismaClient } from "@prisma/client";
import { PrismaErrorConverter } from "./error.js";

export class UserRepository implements IUserRepository {
  private prisma: PrismaClient;

  constructor(prisma: PrismaClient) {
    this.prisma = prisma;
  }

  async Create(u: User): Promise<Result<User, Error>> {
    try {
      const res = await this.prisma.user.create({
        data: {
          id: u.id,
          handle: u.handle,
          fullHandle: u.fullHandle,
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
          following: { include: { following: true } },
          follower: { include: { follower: true } },
        },
      });

      const response: Array<User> = this.convertToDomain(new Array<any>(res));
      return new Success(response[0]);
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async CreateFollow(u: UserFollowEvent): AsyncResult<UserFollowEvent, Error> {
    try {
      const res = await this.prisma.userFollowEvent.create({
        data: {
          follower: {
            connect: {
              id: u.follower.id,
            },
          },
          following: {
            connect: {
              id: u.following.id,
            },
          },
        },
        include: {
          follower: true,
          following: true,
        },
      });
      const resp: UserFollowEvent = this.convertToFollowEventDomain(res);
      return new Success(resp);
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async FindByHandle(handle: string): Promise<Result<User, Error>> {
    try {
      const res = await this.prisma.user.findUniqueOrThrow({
        where: {
          fullHandle: handle,
        },
        include: {
          userAPData: true,
          following: { include: { following: true, follower: true } },
          follower: { include: { follower: true, following: true } },
        },
      });
      // Memo:
      //  follower[N].following -> そのユーザーがフォローしているユーザー
      //  following[N].follower -> そのユーザーをフォローしているユーザー
      return new Success(this.convertToDomain([res as UserEntity])[0]);
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async FindByID(id: Snowflake): Promise<Result<User, Error>> {
    try {
      const res = await this.prisma.user.findUnique({
        where: {
          id: id,
        },
        include: {
          userAPData: true,
          following: { include: { following: true, follower: true } },
          follower: { include: { follower: true, following: true } },
        },
      });

      return new Success(this.convertToDomain([res as UserEntity])[0]);
      // return new Success(this.convertToDomain([res as UserEntity])[0]);
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }

  async Update(): Promise<Result<User, Error>> {
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
      return new Failure(PrismaErrorConverter(e));
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
      return new Failure(PrismaErrorConverter(e));
    }
  }

  private convertToDomain<T extends UserEntity>(ew: Array<T>): Array<User> {
    return ew.map((e) => {
      return new User({
        id: e.id as Snowflake,
        bio: e.bio,
        createdAt: e.createdAt,
        handle: e.handle,
        fullHandle: e.fullHandle,
        headerImageURL: e.headerImageURL,
        iconImageURL: e.iconImageURL,
        isLocalUser: e.isLocalUser,
        nickName: e.nickName,
        password: e.password,
        role: e.role,
        following: e.follower.map((v) => {
          return new UserFollowEvent(
            {
              id: v.following.id as Snowflake,
              bio: v.following.bio,
              iconImageURL: v.following.iconImageURL,
              nickName: v.following.nickName,
              fullHandle: v.following.fullHandle,
            },
            {
              id: v.follower.id as Snowflake,
              bio: v.follower.bio,
              iconImageURL: v.follower.iconImageURL,
              nickName: v.follower.nickName,
              fullHandle: v.follower.fullHandle,
            },
          );
        }),
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
        fullHandle: e.follower.fullHandle,
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
        fullHandle: e.following.fullHandle,
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

  private convertToFollowEventDomain<T extends UserFollowEventEntity>(
    i: T,
  ): UserFollowEvent {
    return new UserFollowEvent(
      { ...i.following, id: i.following.id as Snowflake },
      { ...i.follower, id: i.follower.id as Snowflake },
    );
  }

  async FindFollowEvent(
    followingID: Snowflake,
    followerID: Snowflake,
  ): AsyncResult<UserFollowEvent, Error> {
    try {
      const res = await this.prisma.userFollowEvent.findUniqueOrThrow({
        where: {
          followingID_followerID: {
            followingID: followingID,
            followerID: followerID,
          },
        },
        include: {
          follower: true,
          following: true,
        },
      });
      return new Success(this.convertToFollowEventDomain(res));
    } catch (e: unknown) {
      return new Failure(PrismaErrorConverter(e));
    }
  }
}

export type UserEntity = {
  id: string;
  serverId: string;
  bio: string;
  createdAt: Date;
  handle: string;
  fullHandle: string;
  headerImageURL: string;
  iconImageURL: string;
  nickName: string;
  isLocalUser: boolean;
  password: string;
  role: number;
  follower: Array<{
    follower: {
      id: string;
      fullHandle: string;
      nickName: string;
      bio: string;
      iconImageURL: string;
    };
    following: {
      id: string;
      fullHandle: string;
      nickName: string;
      bio: string;
      iconImageURL: string;
    };
  }>;
  userAPData: {
    followersURL: string;
    followingURL: string;
    inboxURL: string;
    outboxURL: string;
    privateKey: string | null;
    publicKey: string;
    id: string;
  };
};

export type UserFollowEventEntity = {
  follower: {
    id: string;
    fullHandle: string;
    nickName: string;
    bio: string;
    iconImageURL: string;
  };
  following: {
    id: string;
    fullHandle: string;
    nickName: string;
    bio: string;
    iconImageURL: string;
  };
};
