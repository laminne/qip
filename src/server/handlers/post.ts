import { PostController } from "../controller/post";
import { FastifyHandlerMethod } from "../../helpers/fastify";
import { CommonPostRequest } from "../types/post";

export class PostHandler {
  private readonly controller: PostController;
  constructor(controller: PostController) {
    this.controller = controller;
  }

  public FindByID: FastifyHandlerMethod<{ Params: { id: string } }> = async (
    req,
    res,
  ) => {
    const re = await this.controller.FindByID(req.params.id);
    console.log(req.params);
    if (re.isFailure()) {
      res.code(500).send({ message: "failed to find post by ID" });
      console.log(re.value);
      return;
    }
    const sleep = async (ms: number) => {
      return new Promise((r) => setTimeout(r, ms));
    };

    await sleep(3000);
    res.code(200).send(re.value);
    return;
  };

  public CreatePost: FastifyHandlerMethod<{ Body: CommonPostRequest }> = async (
    req,
    res,
  ) => {
    const re = await this.controller.CreatePost(req.body);
    if (re.isFailure()) {
      res.code(500).send({ message: "failed to create post" });
      console.log(re.value);
      return;
    }

    res.code(200).send(re.value);
  };
}
