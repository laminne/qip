import { PostController } from "../controller/post.js";
import { FastifyHandlerMethod } from "../../helpers/fastify.js";
import { CommonPostRequest } from "../types/post.js";

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

  public DeletePost: FastifyHandlerMethod<{ Params: { id: string } }> = async (
    q,
    r,
  ) => {
    const re = await this.controller.DeletePost(q.params.id);
    if (re.isFailure()) {
      return r.code(500).send({ message: "failed to delete post" });
    }
    r.code(204).send();
  };

  public GetTimeline: FastifyHandlerMethod<{}> = async (q, r) => {
    const res = await this.controller.ChronologicalPosts("123");
    if (res.isFailure()) {
      r.code(500).send({ message: "failed to get timeline" });
      return;
    }

    r.code(200).send(res.value);
  };

  public CreateReaction: FastifyHandlerMethod<{ Params: { id: string } }> =
    async (q, r) => {
      const re = await this.controller.Reaction("123", q.params.id);
      if (re.isFailure()) {
        return r.code(500).send({ message: "failed to reaction" });
      }
      r.code(200).send(re.value);
    };
}
