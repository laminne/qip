import { fastify } from "fastify";
import { PostHandler } from "./handlers/post.js";
import { PostRepository } from "../repository/prisma/post.js";
import { PostController } from "./controller/post.js";
import { CreatePostService } from "../service/post/create_post_service.js";
import { FindPostService } from "../service/post/find_post_service.js";
import { PrismaClient } from "@prisma/client";
import { FindServerService } from "../service/server/find_server_service.js";
import { FindUserService } from "../service/user/find_user_service.js";
import { SnowflakeIDGenerator } from "../helpers/id_generator.js";
import { ServerRepository } from "../repository/prisma/server.js";
import { UserRepository } from "../repository/prisma/user.js";
import cors from "@fastify/cors";
import { UserHandlers } from "./handlers/user.js";
import { UserController } from "./controller/user.js";
import { CreateTimelineService } from "../service/post/create_timeline_service.js";
import { DeletePostService } from "../service/post/delete_post_service.js";

export async function StartServer(port: number) {
  const app = fastify({
    logger: false,
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
      findUserService: new FindUserService(userRepository),
      createTimelineService: new CreateTimelineService({
        postRepository: postRepository,
      }),
      deletePostService: new DeletePostService(postRepository),
    }),
  );
  const userHandler = new UserHandlers(
    new UserController({
      findServerService: new FindServerService(serverRepository),
      findUserService: new FindUserService(userRepository),
      findPostService: new FindPostService(postRepository),
    }),
  );

  app.get("/", (q, s) => {
    return { version: "Qip2 Server v0.0.1 (pre-alpha)" };
  });

  app.get("/api/v1/posts/:id", postHandler.FindByID);
  app.delete("/api/v1/posts/:id", postHandler.DeletePost);
  app.post("/api/v1/posts", postHandler.CreatePost);
  app.get("/api/v1/users/:name", userHandler.FindByHandle);
  app.get("/api/v1/users/:name/posts", userHandler.FindUserPosts);
  app.get("/api/v1/timeline/home", postHandler.GetTimeline);

  try {
    await app.listen({ port: port });
    return;
  } catch (e: unknown) {
    return new Error("failed to start server", e as Error as any);
  }
}
