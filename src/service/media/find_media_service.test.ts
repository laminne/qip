import { describe, expect, it } from "vitest";
import { FindMediaService } from "./find_media_service.js";
import { MediaRepository } from "../../repository/inmemory/media.js";
import { Media } from "../../domain/media.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { MediaToMediaData } from "../data/media.js";

describe("FindMediaService", () => {
  const data = new Media({
    id: "999" as Snowflake,
    postID: "115" as Snowflake,
    authorID: "100" as Snowflake,
    blurhash: "abcabcabcabc",
    cached: false,
    isSensitive: false,
    md5Sum: "sumsumsum",
    name: "test.jpg",
    size: 90000,
    thumbnailURL: "https://example.jp",
    type: "image/jpeg",
    url: "https://example.jp",
  });
  const service = new FindMediaService(new MediaRepository([data]));
  const exp = MediaToMediaData(data);

  it("IDで検索できる", async () => {
    const res = await service.findByID("999" as Snowflake);
    expect(res.value).toStrictEqual(exp);
  });

  it("投稿者で検索できる", async () => {
    const res = await service.findByUserID("100" as Snowflake);
    expect(res.value).toStrictEqual([exp]);
  });

  it("投稿IDで検索できる", async () => {
    const res = await service.findByPostID("115" as Snowflake);
    expect(res.value).toStrictEqual([exp]);
  });
});
