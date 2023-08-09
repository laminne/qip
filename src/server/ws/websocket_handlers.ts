import * as http from "http";
import WebSocket, { WebSocketServer } from "ws";
import { IncomingMessage } from "http";
import { Snowflake } from "../../helpers/id_generator.js";
import logger from "../../helpers/logger.js";

export interface NotificationSendReceiver {
  Send(payload: string, to: string): Promise<void>;
}

export class DummyNotificationSendReceiver implements NotificationSendReceiver {
  async Send(): Promise<void> {
    return;
  }
}

export class WebSocketNotificationSendReceiver
  implements NotificationSendReceiver
{
  private ws: WebSocket.Server<
    typeof WebSocket.WebSocket,
    typeof IncomingMessage
  >;
  private connections: Map<Snowflake, WebSocket.WebSocket>;

  constructor(server: http.Server) {
    this.ws = new WebSocketServer({ noServer: true });
    server.on("upgrade", (q, so, head) => {
      // fixme: q.headers.authorizationを使って認証とユーザーの検証をする
      this.ws.handleUpgrade(q, so, head, (w) => {
        this.ws.emit("connection", w, q);
        this.connections.set("123" as Snowflake, w);
      });
    });

    this.connections = new Map<Snowflake, WebSocket.WebSocket>();

    this.ws.on("connection", (w) => {
      logger.info("connect");
      w.on("message", () => {
        console.log("received!");
      });
    });
  }

  async Send(payload: string, to: string): Promise<void> {
    const cli = this.connections.get(to as Snowflake);
    if (!cli) {
      return;
    }
    cli.send(payload);
    return;
  }
}
