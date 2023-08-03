import { describe, expect, it } from "vitest";
// import { User, UserAPData, UserFollowEvent } from "../../domain/user.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { UserRepository } from "../../repository/inmemory/user.js";
import { FindUserService } from "./find_user_service.js";
import { User, UserAPData, UserFollowEvent } from "../../domain/user.js";
import { UserToUserData } from "../data/user.js";

describe("FindUserService", () => {
  const exp = [
    new User({
      id: "100" as Snowflake,
      serverID: "120" as Snowflake,
      bio: "hello!",
      createdAt: new Date("2021-02-20 10:00:00"),
      fullHandle: "test@social.example.jp",
      handle: "test",
      headerImageURL: "https://media.example.jp",
      iconImageURL: "https://media.example.jp",
      isLocalUser: false,
      nickName: "Hej",
      password: "",
      role: 0,
      apData: new UserAPData({
        followersURL: "none",
        followingURL: "none",
        inboxURL: "none",
        outboxURL: "none",
        privateKey: null,
        publicKey: "none",
        userAPID: "10" as Snowflake,
        userID: "100" as Snowflake,
      }),
      following: new Array<UserFollowEvent>(),
    }),
    new User({
      id: "101" as Snowflake,
      serverID: "120" as Snowflake,
      bio: "こんにちは",
      createdAt: new Date("2020-02-20 10:00:00"),
      fullHandle: "hello@social.example.jp",
      handle: "hello",
      headerImageURL: "https://media.example.jp",
      iconImageURL: "https://media.example.jp",
      isLocalUser: false,
      nickName: "おはようございます",
      password: "",
      role: 0,
      apData: new UserAPData({
        followersURL: "none",
        followingURL: "none",
        inboxURL: "none",
        outboxURL: "none",
        privateKey: null,
        publicKey: "none",
        userAPID: "11" as Snowflake,
        userID: "101" as Snowflake,
      }),
      following: new Array<UserFollowEvent>(),
    }),
  ];
  exp[0].follow(exp[1]);

  const repository = new UserRepository(exp);
  const service = new FindUserService(repository);

  it("IDで検索できる", async () => {
    const res = await service.FindByID("100" as Snowflake);
    expect(res.value).toStrictEqual(UserToUserData(exp[0]));
  });

  it("ハンドルで検索できる", async () => {
    // ここでのハンドルはフルハンドル(<handle>@<host>)
    const res = await service.FindByHandle("hello@social.example.jp");
    expect(res.value).toStrictEqual(UserToUserData(exp[1]));
  });
});
