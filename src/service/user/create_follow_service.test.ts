import { describe, expect, it } from "vitest";
import { CreateFollowService } from "./create_follow_service.js";
import { UserRepository } from "../../repository/inmemory/user.js";
import { User, UserAPData } from "../../domain/user.js";
import { Snowflake } from "../../helpers/id_generator.js";

describe("create_follow_service", () => {
  const repository = new UserRepository([
    new User({
      id: "1" as Snowflake,
      fullHandle: "testuser@example.com",
      password: "testUserPassword",
      role: 0,
      nickName: "test",
      handle: "test",
      bio: "test",
      headerImageURL: "test",
      iconImageURL: "test",
      isLocalUser: true,
      serverID: "1" as Snowflake,
      apData: new UserAPData({
        userID: "1" as Snowflake,
        userAPID: "2" as Snowflake,
        inboxURL: "test",
        outboxURL: "test",
        followersURL: "https://example.com",
        followingURL: "https://example.com",
        publicKey: "test",
        privateKey: null,
      }),
      following: [],
      createdAt: new Date(),
    }),
    new User({
      id: "2" as Snowflake,
      fullHandle: "testuser2@example.com",
      password: "testuser2password",
      role: 0,
      nickName: "test",
      handle: "test",
      bio: "test",
      headerImageURL: "test",
      iconImageURL: "test",
      isLocalUser: true,
      serverID: "1" as Snowflake,
      apData: new UserAPData({
        userID: "2" as Snowflake,
        userAPID: "3" as Snowflake,
        inboxURL: "test",
        outboxURL: "test",
        followersURL: "https://example.com",
        followingURL: "https://example.com",
        publicKey: "test",
        privateKey: null,
      }),
      following: [],
      createdAt: new Date(),
    }),
  ]);
  const service = new CreateFollowService(repository);

  it("フォローできる", async () => {
    const res = await service.Handle("1" as Snowflake, "2" as Snowflake);
    expect(res.value).not.toBeUndefined();
    expect(res.isFailure()).toBe(false);
  });

  it("すでにフォローしている相手はフォローできない", async () => {
    await service.Handle("1" as Snowflake, "2" as Snowflake);
    const res2 = await service.Handle("1" as Snowflake, "2" as Snowflake);
    expect(res2.isFailure()).toBe(true);
  });
});
