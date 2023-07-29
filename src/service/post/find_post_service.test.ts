import { describe, expect, it } from "vitest";
import { PostRepository } from "../../repository/inmemory/post";
import { Post, PostReactionEvent } from "../../domain/post";
import { Snowflake } from "../../helpers/id_generator";
import { Media } from "../../domain/media";
import { FindPostService } from "./find_post_service";
import { PostToPostData } from "../data/post";

describe("FindPostService", () => {
  const exp = [
    new Post({
      id: "123" as Snowflake,
      authorID: "100" as Snowflake,
      createdAt: new Date("2021-10-20 00:00:00"),
      text: "あいうえお",
      visibility: 0,
      attachments: new Array<Media>(),
      reactions: new Array<PostReactionEvent>(),
    }),
    new Post({
      id: "121" as Snowflake,
      authorID: "101" as Snowflake,
      attachments: new Array<Media>(),
      reactions: new Array<PostReactionEvent>(),
      text: "テスト投稿",
      visibility: 0,
      createdAt: new Date("2020-10-20 00:00:00"),
    }),
  ];
  const postRepository = new PostRepository(exp);

  const findService = new FindPostService(postRepository);

  it("IDで取得できる", async () => {
    const res = await findService.FindByID("123" as Snowflake);
    expect(res.value).toStrictEqual(PostToPostData(exp[0]));
  });
  it("投稿者のIDで取得できる", async () => {
    const res = await findService.FindByAuthor("101" as Snowflake);
    expect(res.value).toStrictEqual([PostToPostData(exp[1])]);
  });
});
