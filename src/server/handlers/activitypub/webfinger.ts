import { FastifyHandlerMethod } from "../../../helpers/fastify.js";
import {
  WebFinger,
  WebFingerController,
} from "../../controller/activitypub/webfinger.js";
import logger from "../../../helpers/logger.js";

export class WebFingerHandler {
  private readonly controller: WebFingerController;

  constructor(controller: WebFingerController) {
    this.controller = controller;
  }

  public Handle: FastifyHandlerMethod<{
    Querystring: { resource: string };
    Reply: WebFinger;
  }> = async (q, r) => {
    // 方針: ここからQueryのacct(フルハンドル)をFindUserServiceで検索して返す
    // なお、ハンドラにはエラー処理等しか書かない
    // APのJSONLDの変換自体はハンドラではなくコントローラーで行う
    const res = await this.controller.Handle(q.query.resource.substring(5));
    if (res.isFailure()) {
      logger.error(res.value);
      return r.code(400).send();
    }

    return r
      .code(200)
      .send(res.value)
      .header("Content-Type", "application/activity+json");
  };
}
