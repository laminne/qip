import { describe, expect, it } from "vitest";
import { ServerRepository } from "../../repository/inmemory/server.js";
import { Server } from "../../domain/server.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { FindServerService } from "./find_server_service.js";
import { ServerToServerData } from "../data/server.js";

describe("findServerService", () => {
  const exp = [
    new Server({
      id: "100" as Snowflake,
      description: "hello world!",
      faviconURL: "https://example.com",
      host: "https://social.example.com",
      iconURL: "https://example.com",
      maintainer: "John Doe",
      maintainerEmail: "john@example.com",
      name: "John's Home",
      softwareName: "Mastodon",
      softwareVersion: "v4.2.1",
    }),
    new Server({
      id: "101" as Snowflake,
      description: "Qip2 Server",
      faviconURL: "https://example.com",
      host: "https://q.example.jp",
      iconURL: "https://example.com",
      maintainer: "Yamada",
      maintainerEmail: "yamada@example.com",
      name: "myBase",
      softwareName: "Qip2",
      softwareVersion: "v0.0.1",
    }),
  ];
  const serverRepository = new ServerRepository(exp);
  const service = new FindServerService(serverRepository);
  it("IDで取得できる", async () => {
    const res = await service.FindByID("100" as Snowflake);
    expect(res.value).toStrictEqual(ServerToServerData(exp[0]));
  });

  it("ホスト名で取得できる", async () => {
    const res = await service.FindByHost("https://q.example.jp");
    expect(res.value).toStrictEqual(ServerToServerData(exp[1]));
  });
});
