import { FastifyHandlerMethod } from "../../../helpers/fastify.js";
import { PersonController } from "../../controller/activitypub/person.js";
import logger from "../../../helpers/logger.js";

export class PersonHandler {
  private readonly controller: PersonController;

  constructor(controller: PersonController) {
    this.controller = controller;
  }

  public Handle: FastifyHandlerMethod<{ Params: { id: string } }> = async (
    q,
    r,
  ) => {
    const res = await this.controller.Handle(q.params.id);
    if (res.isFailure()) {
      logger.error(res.value);
      return r.code(404).send();
    }

    return r.code(200).send(res.value);
  };
}
