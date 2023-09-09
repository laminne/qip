import { describe, expect, it } from "vitest";
import { Server, ServerArgs } from "./server.js";
import { Snowflake } from "../helpers/id_generator.js";

describe("Server", () => {
  const defaultServerArgs: ServerArgs = {
    id: "10101294855225344" as Snowflake,
    host: "social.example.jp",
    softwareName: "Mastodon",
    softwareVersion: "4.0.0",
    name: "example social",
    description: "this is example social",
    maintainer: "example maintainer",
    maintainerEmail: "johndoe@example.jp",
    iconURL: "https://social.example.jp/icon.png",
    faviconURL: "https://social.example.jp/favicon.png",
  };

  it("128文字以上のホスト名は設定できない", () => {
    expect(() => {
      new Server({
        ...defaultServerArgs,
        host: "a".repeat(128 + 1),
      });
    }).toThrowError(new Error("failed to create server: host is too long"));
  });

  it("4文字以下のホスト名は設定できない", () => {
    expect(() => {
      new Server({
        ...defaultServerArgs,
        host: "a.b",
      });
    }).toThrowError(new Error("failed to create server: host is too short"));
  });

  it("ホスト名を設定できる", () => {
    expect(() => {
      new Server({
        ...defaultServerArgs,
      });
    }).toBeTruthy();
  });

  it("128文字以上のソフトウェア名は設定できない", () => {
    expect(() => {
      new Server({
        ...defaultServerArgs,
        softwareName: "a".repeat(128 + 1),
      });
    }).toThrowError(
      new Error("failed to create server: softwareName is too long"),
    );
  });

  it("ソフトウェア名を設定できる", () => {
    expect(() => {
      new Server({
        ...defaultServerArgs,
      });
    }).toBeTruthy();
  });

  it("128文字以上のソフトウェアバージョンは設定できない", () => {
    expect(() => {
      new Server({
        ...defaultServerArgs,
        softwareVersion: "a".repeat(128 + 1),
      });
    }).toThrowError(
      new Error("failed to create server: softwareVersion is too long"),
    );
  });

  it("128文字以上のサーバー名は設定できない", () => {
    expect(() => {
      new Server({
        ...defaultServerArgs,
        name: "a".repeat(128 + 1),
      });
    }).toThrowError(new Error("failed to create server: name is too long"));
  });

  it("3000文字以上のサーバー説明文は設定できない", () => {
    expect(() => {
      new Server({
        ...defaultServerArgs,
        description: "a".repeat(3000 + 1),
      });
    }).toThrowError(
      new Error("failed to create server: description is too long"),
    );
  });

  it("256文字以上の管理者名は設定できない", () => {
    expect(() => {
      new Server({
        ...defaultServerArgs,
        maintainer: "a".repeat(256 + 1),
      });
    }).toThrowError(
      new Error("failed to create server: maintainer is too long"),
    );
  });

  it("256文字以上の管理者メールアドレスは設定できない", () => {
    expect(() => {
      new Server({
        ...defaultServerArgs,
        maintainerEmail: "a".repeat(256 + 1),
      });
    }).toThrowError(
      new Error("failed to create server: maintainerEmail is too long"),
    );
  });
});
