import { describe, expect, it } from "vitest";
import { PostReactionEvent } from "../../domain/post.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { CreateReactionService } from "./create_reaction_service.js";
import { ReactionRepository } from "../../repository/inmemory/reaction.js";

describe("CreateReactionService", () => {
  const exp = [
    new PostReactionEvent("100" as Snowflake, "200" as Snowflake),
    new PostReactionEvent("101" as Snowflake, "201" as Snowflake),
  ];
  const repository = new ReactionRepository([exp[1]]);
  const service = new CreateReactionService({
    repository: repository,
  });

  it("リアクションを作成できる", async () => {
    const res = await service.Handle("100" as Snowflake, "200" as Snowflake);
    expect(res.value).toStrictEqual(exp[0]);
  });

  it("1つの投稿に対して複数回リアクションできない", async () => {
    const res2 = await service.Handle("101" as Snowflake, "201" as Snowflake);
    expect(res2.isSuccess()).toBe(false);
  });
});
