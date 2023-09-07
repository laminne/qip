import { FastifyHandlerMethod } from "../../helpers/fastify.js";
import { UserController } from "../controller/user.js";
import { ErrorConverter } from "./error.js";
export class UserHandlers {
  private readonly controller: UserController;
  constructor(controller: UserController) {
    this.controller = controller;
  }

  public FindByHandle: FastifyHandlerMethod<{ Params: { name: string } }> =
    async (q, r) => {
      const res = await this.controller.FindByHandle(q.params.name);
      if (res.isFailure()) {
        const [code, message] = ErrorConverter(res.value);
        return r.code(code).send(message);
      }
      r.code(200).send(res.value);
      return;
    };

  public FindUserPosts: FastifyHandlerMethod<{ Params: { name: string } }> =
    async (q, r) => {
      const res = await this.controller.FindUserPosts(q.params.name);
      if (res.isFailure()) {
        const [code, message] = ErrorConverter(res.value);
        return r.code(code).send(message);
      }

      r.code(200).send(res.value);
      return;
    };

  public CreateFollow: FastifyHandlerMethod<{ Params: { id: string } }> =
    async (q, r) => {
      // APIを叩いたユーザー - フォロー > フォロー先(params.id)
      const res = await this.controller.CreateFollow(q.params.id, "123");
      if (res.isFailure()) {
        const [code, message] = ErrorConverter(res.value);
        return r.code(code).send(message);
      }

      r.code(200).send(res.value);
    };
}
