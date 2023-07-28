import { StartServer } from "./server/start";

console.log("Qip2\n(C) 2023 Tatsuto 'Laminne' YAMAMOTO");
console.log(`  ___  _         ____  
 / _ \\(_)_ __   |___ \\ 
| | | | | '_ \\    __) |
| |_| | | |_) |  / __/ 
 \\__\\_\\_| .__/  |_____|
        |_|`);
console.log("H e l l o ( a g a i n )");
StartServer(6300).then((r) => {});
// const prisma = new PrismaClient();
// const r = new PostRepository(prisma);

// const c = new PostController({
//   findPostService: new FindPostService(new PostRepository(prisma)),
//   findServerService: new FindServerService(new ServerRepository(prisma)),
//   findUserService: new FindUserService(new UserRepository(prisma)),
//   createTimelineService: new CreateTimelineService({
//     postRepository: new PostRepository(prisma),
//     userRepository: new UserRepository(prisma),
//   }),
//   createPostService: new CreatePostService(
//     new PostRepository(prisma),
//     new SnowflakeIDGenerator(1),
//   ),
// });

// c.ChronologicalPosts("9554100359925760").then((r) => console.log(r.value));

// r.ChronologicalPosts("9554100359925760" as Snowflake, 0)
//   .then((r) => {
//     // if (r.isFailure()) {
//     //   return;
//     // }
//     console.log(r.value);
//   })
//   .catch((e) => console.log(e));
