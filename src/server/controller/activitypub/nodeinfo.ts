import { NodeInfo20 } from "../../handlers/activitypub/nodeinfo.js";

export class NodeInfoController {
  constructor() {}

  Handle(): NodeInfo20 {
    return {
      version: "2.0",
      software: {
        name: "Qip2",
        version: "v0.0.1",
      },
      protocols: ["activitypub"],
      services: {
        inbound: [],
        outbound: [],
      },
      openRegistrations: false,
      // ToDo: ユーザー数と投稿数を計算する
      usage: {
        users: {
          total: 0,
        },
        localPosts: 0,
      },
      metadata: {
        // ToDo: インスタンス名 / メンテナなどを取得できるようにする
        nodeName: "Whiterabbit",
        nodeDescription: "",
        maintainer: {
          name: "",
          email: "",
        },
      },
    };
  }
}
