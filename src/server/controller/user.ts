import { FindUserService } from "../../service/user/find_user_service";
import { Snowflake } from "../../helpers/id_generator";
import { AsyncResult, Failure, Success } from "../../helpers/result";
import { UserResponse } from "../types/user";
import { FindServerService } from "../../service/server/find_server_service";

export class UserController {
  private readonly findUserService: FindUserService;
  private readonly findServerService: FindServerService;
  constructor(args: {
    findUserService: FindUserService;
    findServerService: FindServerService;
  }) {
    this.findUserService = args.findUserService;
    this.findServerService = args.findServerService;
  }

  async FindByHandle(name: string): AsyncResult<UserResponse, any> {
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
        return new Failure("failed to parse handle");
    }
    const user = await this.findUserService.FindByHandle(uName);
    if (user.isFailure()) {
      return new Failure(new Error("failed to find user", user.value));
    }
    // ToDo: ここのNを解決する(Userが属するサーバーを一発で取れるようにする)
    // ToDo: 取得したユーザーリストからacctが合致するものを取り出す
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
      return new Failure("failed to find user");
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
}
