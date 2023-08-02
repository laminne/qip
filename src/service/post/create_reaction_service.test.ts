import { describe, expect, it } from "vitest";
import { PostReactionEvent } from "../../domain/post.js";
import { Snowflake, SnowflakeIDGenerator } from "../../helpers/id_generator.js";
import { CreateReactionService } from "./create_reaction_service.js";
import { ReactionRepository } from "../../repository/inmemory/reaction.js";

describe("CreateReactionService", () => {
  const exp = [
    new PostReactionEvent("" as Snowflake, "" as Snowflake),
    new PostReactionEvent("" as Snowflake, "" as Snowflake),
  ];
  const generator = new SnowflakeIDGenerator(1);
  const repository = new ReactionRepository();
  const service = new CreateReactionService({
    repository: repository,
    idGenerator: generator,
  });

  it("リアクションを作成できる", async () => {
    const res = await service.Handle("" as Snowflake, "" as Snowflake);
    expect(res).toStrictEqual(exp[0]);
  });

  it("1つの投稿に対して複数回リアクションできない", async () => {
    await service.Handle("" as Snowflake, "" as Snowflake);
    const res2 = await service.Handle("" as Snowflake, "" as Snowflake);
    expect(res2.isFailure()).toBe(true);
  });
});
