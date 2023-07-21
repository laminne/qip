import { fastify } from "fastify";
import { PostHandler } from "./handlers/post";
import { PostRepository } from "../repository/prisma/post";
import { PostController } from "./controller/post";
import { CreatePostService } from "../service/post/create_post_service";
import { FindPostService } from "../service/post/find_post_service";
import { PrismaClient } from "@prisma/client";
import { FindServerService } from "../service/server/find_server_service";
import { FindUserService } from "../service/user/find_user_service";
import { SnowflakeIDGenerator } from "../helpers/id_generator";
import { ServerRepository } from "../repository/prisma/server";
import { UserRepository } from "../repository/prisma/user";
import cors from "@fastify/cors";
export async function StartServer(port: number) {
  const app = fastify({
    logger: true,
  });
  app.register(cors, {});

  const prisma = new PrismaClient();

  const postRepository = new PostRepository(prisma);
  const serverRepository = new ServerRepository(prisma);
  const userRepository = new UserRepository(prisma);
  const idGen = new SnowflakeIDGenerator(1);
  const postHandler = new PostHandler(
    new PostController({
      createPostService: new CreatePostService(postRepository, idGen),
      findPostService: new FindPostService(postRepository),
      findServerService: new FindServerService(serverRepository),
      findUserService: new FindUserService(userRepository),
    }),
  );

  app.get("/", (q, s) => {
    return { version: "Qip2 Server v0.0.1 (pre-alpha)" };
  });

  app.get("/api/v1/posts/:id", postHandler.FindByID);
  app.post("/api/v1/posts", postHandler.CreatePost);

  try {
    await app.listen({ port: port });
  } catch (e: unknown) {
    return new Error("failed to start server", e as Error as any);
  }
}
