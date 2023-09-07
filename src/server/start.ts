/* eslint-disable @typescript-eslint/no-explicit-any */
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
import { PreCheck } from "./pre_check.js";
import { CreateReactionService } from "../service/post/create_reaction_service.js";
import { ReactionRepository } from "../repository/prisma/reaction.js";
import { WebFingerHandler } from "./handlers/activitypub/webfinger.js";
import { WebFingerController } from "./controller/activitypub/webfinger.js";
import { NodeInfoHandlers } from "./handlers/activitypub/nodeinfo.js";
import { NodeInfoController } from "./controller/activitypub/nodeinfo.js";
import { PersonHandler } from "./handlers/activitypub/person.js";
import { PersonController } from "./controller/activitypub/person.js";
import logger from "../helpers/logger.js";
import { CreateFollowService } from "../service/user/create_follow_service.js";

export async function StartServer(port: number) {
  const app = fastify({
    logger: true,
  });
  app.addContentTypeParser(
    "application/activity+json",
    { parseAs: "string" },
    (req, body, done) => {
      try {
        const parsedBody = JSON.parse(typeof body === "string" ? body : "");
        done(null, parsedBody);
      } catch (err) {
        done(err as Error as any);
      }
    },
  );
  app.register(cors, {});

  const prisma = new PrismaClient();
  await PreCheck(prisma);
  const postRepository = new PostRepository(prisma);
  const serverRepository = new ServerRepository(prisma);
  const userRepository = new UserRepository(prisma);
  const reactionRepository = new ReactionRepository(prisma);
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
      createReactionService: new CreateReactionService({
        repository: reactionRepository,
      }),
    }),
  );
  const userHandler = new UserHandlers(
    new UserController({
      findServerService: new FindServerService(serverRepository),
      findUserService: new FindUserService(userRepository),
      findPostService: new FindPostService(postRepository),
      createFollowService: new CreateFollowService(userRepository),
    }),
  );
  const apHandler = new WebFingerHandler(
    new WebFingerController(new FindUserService(userRepository)),
  );
  const nodeinfoHandler = new NodeInfoHandlers(new NodeInfoController());
  const personHandler = new PersonHandler(
    new PersonController(new FindUserService(userRepository)),
  );
  app.get("/", () => {
    return { version: "Qip2 Server v0.0.1 (pre-alpha)" };
  });

  app.get("/api/v1/posts/:id", postHandler.FindByID);
  app.delete("/api/v1/posts/:id", postHandler.DeletePost);
  app.post("/api/v1/posts/:id/reaction", postHandler.CreateReaction);
  app.delete("/api/v1/posts/:id/reaction", postHandler.UndoReaction);
  app.post("/api/v1/posts", postHandler.CreatePost);
  app.get("/api/v1/users/:name", userHandler.FindByHandle);
  app.post("/api/v1/users/:id/follow", userHandler.CreateFollow);
  app.get("/api/v1/users/:name/posts", userHandler.FindUserPosts);
  app.get("/api/v1/timeline/home", postHandler.GetTimeline);

  app.get("/.well-known/webfinger", apHandler.Handle);
  app.get("/nodeinfo/2.0", nodeinfoHandler.Handle);
  app.get("/users/:id", personHandler.Handle);
  app.post("/users/:id/inbox", (q, r) => {
    logger.info(q.body, "inbox");
    r.code(503).send();
  });

  try {
    await app.listen({ port: port, host: "0.0.0.0" });
    return;
  } catch (e: unknown) {
    return new Error("failed to start server", e as Error as any);
  }
}
