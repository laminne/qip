import { FastifyHandlerMethod } from "../../helpers/fastify.js";
import { UserController } from "../controller/user.js";
export class UserHandlers {
  private readonly controller: UserController;
  constructor(controller: UserController) {
    this.controller = controller;
  }

  public FindByHandle: FastifyHandlerMethod<{ Params: { name: string } }> =
    async (q, r) => {
      const res = await this.controller.FindByHandle(q.params.name);
      if (res.isFailure()) {
        r.code(500).send({ message: "failed to find id by ID" });
        return;
      }
      r.code(200).send(res.value);
      return;
    };

  public FindUserPosts: FastifyHandlerMethod<{ Params: { name: string } }> =
    async (q, r) => {
      const res = await this.controller.FindUserPosts(q.params.name);
      if (res.isFailure()) {
        r.code(500).send({ message: "failed to find user posts" });
        return;
      }

      r.code(200).send(res.value);
      return;
    };
}
