import { describe, expect, it } from "vitest";
import { Post, PostArgs } from "./post.js";
import { Snowflake } from "../helpers/id_generator.js";
import { Media } from "./media.js";

describe("Post", () => {
  const defaultPostArgs: PostArgs = {
    id: "10101450391945216" as Snowflake,
    authorID: "10101451896913920" as Snowflake,
    text: "Hello, world!",
    visibility: 0,
    createdAt: new Date(),
    attachments: [],
    reactions: [],
  };

  it("添付ファイルは16個以上添付できない", () => {
    expect(() => {
      new Post({
        ...defaultPostArgs,
        attachments: new Array(17).fill({} as Media),
      });
    }).toThrowError(
      new Error(
        "failed to create post: The number of attachments must be less than 16.",
      ),
    );
  });

  it("ファイルを添付できる", () => {
    expect(() => {
      new Post({
        ...defaultPostArgs,
        attachments: new Array(10).fill({} as Media),
      });
    }).toBeTruthy();
  });
});
