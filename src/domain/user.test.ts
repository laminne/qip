import { describe, expect, it } from "vitest";
import { User, UserAPData } from "./user.js";
import { Snowflake } from "../helpers/id_generator.js";

describe("User domain model", () => {
  const defaultDataArgs = {
    id: "10100912877076480" as Snowflake,
    serverID: "10100914303139840" as Snowflake,
    handle: "test",
    fullHandle: "test@example.jp",
    nickName: "TestUser",
    bio: "testのユーザーです",
    headerImageURL: "https://image.example.jp/header.png",
    iconImageURL: "https://image.example.jp/icon.png",
    password: "efnrkgnjkneiug",
    role: 0,
    createdAt: new Date(),
    isLocalUser: true,
    following: [],
    apData: new UserAPData({
      followersURL: "https://example.jp/followers",
      followingURL: "https://example.jp/following",
      inboxURL: "https://example.jp/inbox",
      outboxURL: "https://example.jp/outbox",
      privateKey: "",
      publicKey: "",
      userAPID: "10100938952278016",
      userID: "10100912877076480" as Snowflake,
    }),
  };

  it("64文字以内の表示名は設定できる", () => {
    const u = new User({ ...defaultDataArgs, nickName: "テストユーザー" });
    expect(u).toBeTruthy();
  });

  it("64文字以上の表示名は設定できない", () => {
    expect(
      () =>
        new User({
          ...defaultDataArgs,
          nickName: "テストユーザー".repeat(10),
        }),
    ).toThrowError(new Error("failed to create user: nickname is too long"));
  });

  it("64文字以上のハンドルは設定できない", () => {
    expect(
      () =>
        new User({
          ...defaultDataArgs,
          handle: "test_user".repeat(20),
        }),
    ).toThrowError(new Error("failed to create user: handle is too long"));
  });

  it("64文字以下のハンドルを設定できる", () => {
    expect(new User({ ...defaultDataArgs, handle: "test_user" })).toBeTruthy();
  });

  it("空文字列は設定できない", () => {
    expect(() => new User({ ...defaultDataArgs, handle: "" })).toThrowError(
      new Error("failed to create user: handle is too short"),
    );
  });

  it("フルハンドルを設定できる", () => {
    expect(
      new User({ ...defaultDataArgs, fullHandle: "test@social.example.jp" }),
    ).toBeTruthy();
  });

  it("FQDNとして成立しない長さのフルハンドルは設定できない", () => {
    expect(
      () =>
        new User({
          ...defaultDataArgs,
          fullHandle: "t@a.b",
        }),
    ).toThrowError(new Error("failed to create user: fullHandle is too short"));
  });

  it("3000文字以上のbioは設定できない", () => {
    expect(
      () =>
        new User({
          ...defaultDataArgs,
          bio: "a".repeat(3001),
        }),
    ).toThrowError(new Error("failed to create user: bio is too long"));
  });

  it("bioを設定できる", () => {
    expect(
      new User({
        ...defaultDataArgs,
        bio: "テストユーザーです",
      }),
    ).toBeTruthy();
  });
});
