import { FindUserService } from "../../service/user/find_user_service";
import { AsyncResult, Failure, Result, Success } from "../../helpers/result";
import { UserResponse } from "../types/user";
import { FindServerService } from "../../service/server/find_server_service";
import { FindPostService } from "../../service/post/find_post_service";
import { CommonPostResponse } from "../types/post";
import { Snowflake } from "../../helpers/id_generator";
import { UserData } from "../../service/data/user";
import { User } from "../../domain/user";
import { ServerData } from "../../service/data/server";

export class UserController {
  private readonly findUserService: FindUserService;
  private readonly findServerService: FindServerService;
  private readonly findPostService: FindPostService;

  constructor(args: {
    findUserService: FindUserService;
    findServerService: FindServerService;
    findPostService: FindPostService;
  }) {
    this.findUserService = args.findUserService;
    this.findServerService = args.findServerService;
    this.findPostService = args.findPostService;
  }

  async FindByHandle(name: string): AsyncResult<UserResponse, Error> {
    // acctを分ける
    const parsed = name.split("@");
    let uName: string;
    let host: string;
    switch (parsed.length) {
      case 2:
        // ABC@EXAMPLE.COM
        uName = parsed[0];
        host = parsed[1];
        break;
      case 3:
        // @ABC@EXAMPLE.COM
        uName = parsed[1];
        host = parsed[2];
        break;
      default:
        return new Failure(new Error("failed to parse handle"));
    }
    const user = await this.findUserService.FindByHandle(uName);
    if (user.isFailure()) {
      return new Failure(new Error("failed to find user", user.value));
    }
    const server = await this.findServerService.FindByHost(host);
    if (server.isFailure()) {
      return new Failure(
        new Error("failed to find user server data", server.value),
      );
    }
    const r = user.value.find((v) => {
      return v.serverID === server.value.id;
    });
    if (!r) {
      return new Failure(new Error("failed to find user"));
    }

    const res: UserResponse = {
      id: r.id,
      host: `@${r.handle}@${server.value.host}`,
      nickName: r.nickName,
      role: r.role,
      bio: r.bio,
      headerImageURL: r.headerImageURL,
      iconImageURL: r.iconImageURL,
      following: r.following,
      softwareName: server.value.softwareName,
    };
    return new Success(res);
  }

  async FindUserPosts(id: string): AsyncResult<CommonPostResponse[], Error> {
    const u = await this.FindByHandle(id);
    if (u.isFailure()) {
      return new Failure(u.value);
    }
    // ユーザーの投稿を取ってくる
    const posts = await this.findPostService.FindByAuthor(
      u.value.id as Snowflake,
    );
    if (posts.isFailure()) {
      return new Failure(posts.value);
    }

    // ユーザーの情報を取ってくる
    // RN/BT/RTに相当するものがないのでこれで良い(全てリンクによる引用になる)
    const user = await this.findUserService.FindByID(u.value.id as Snowflake);
    if (user.isFailure()) {
      return new Failure(user.value);
    }

    // ユーザーのサーバー情報を持ってくる
    const server = await this.findServerService.FindByID(user.value.serverID);
    if (server.isFailure()) {
      return new Failure(server.value);
    }

    return new Success(
      posts.value.map((v): CommonPostResponse => {
        return {
          id: v.id,
          author: {
            host: `@${user.value.handle}@${server.value.host}`,
            iconImageURL: user.value.iconImageURL,
            id: user.value.id,
            nickName: user.value.nickName,
          },
          createdAt: v.createdAt,
          reactions: v.reactions,
          attachments: v.attachments,
          text: v.text,
        };
      }),
    );
  }
}
