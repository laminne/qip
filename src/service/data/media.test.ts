import { describe, expect, it } from "vitest";
import { Media } from "../../domain/media.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { MediaData, MediaToMediaData } from "./media.js";

describe("正しくドメインモデルからDTOに変換できる", () => {
  it("MediaをMediaDataに変換できる", () => {
    const d = new Media({
      id: "100" as Snowflake,
      authorID: "123" as Snowflake,
      blurhash: "",
      cached: true,
      isSensitive: true,
      md5Sum: "abc",
      name: "okayama.jpg",
      postID: "114" as Snowflake,
      size: 999,
      thumbnailURL: "http://example.jp/thumbnail.jpg",
      type: "image/jpg",
      url: "http://example.jp/okayama.jpg",
    });

    const exp = new MediaData({
      id: "100" as Snowflake,
      authorID: "123" as Snowflake,
      blurhash: "",
      cached: true,
      isSensitive: true,
      md5Sum: "abc",
      name: "okayama.jpg",
      postID: "114" as Snowflake,
      size: 999,
      thumbnailURL: "http://example.jp/thumbnail.jpg",
      type: "image/jpg",
      url: "http://example.jp/okayama.jpg",
    });

    expect(MediaToMediaData(d)).toStrictEqual(exp);
  });
});
