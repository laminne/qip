import { FindUserService } from "../../../service/user/find_user_service.js";
import {
  AsyncResult,
  Failure,
  Result,
  Success,
} from "../../../helpers/result.js";
import logger from "../../../helpers/logger.js";

export class WebFingerController {
  private readonly findUserService: FindUserService;

  constructor(service: FindUserService) {
    this.findUserService = service;
  }

  async Handle(acct: string): AsyncResult<WebFinger, Error> {
    // ここでのacctはフルハンドルと一緒
    const a = this.acctConverter(acct);
    if (a.isFailure()) {
      return new Failure(a.value);
    }
    logger.info(a.value);
    const user = await this.findUserService.FindByHandle(a.value);
    if (user.isFailure()) {
      return new Failure(user.value);
    }

    const res: WebFinger = {
      subject: `acct:${a.value}`,
      // ToDo: 自分自身のFQDNをグローバル定数として用意する
      aliases: [`https://wrt2.laminne33569.net/users/${user.value.id}`],
      links: [
        {
          rel: "self",
          type: "application/activity+json",
          href: `https://wrt2.laminne33569.net/users/${user.value.id}`,
        },
      ],
    };

    return new Success(res);
  }

  private acctConverter(acct: string): Result<string, Error> {
    const split = acct.split("@");
    switch (split.length) {
      case 2:
        return new Success(`${split[0]}@${split[1]}`);
      case 3:
        return new Success(`${split[1]}@${split[2]}`);
      default:
        // パースエラー
        return new Failure(new Error("failed to parse acct"));
    }
  }
}

export interface WebFinger {
  subject: string;
  aliases: string[];
  links: [{ rel: string; type: string; href: string }];
}
