import { NotificationSendReceiver } from "../../server/ws/websocket_handlers.js";
import { UserData } from "../data/user.js";
import { Snowflake } from "../../helpers/id_generator.js";
import {
  FollowedNotification,
  FollowSuccessNotification,
  ReactionNotification,
} from "./@types.js";

export class NotificationSenderService {
  private sender: NotificationSendReceiver;
  constructor(sender: NotificationSendReceiver) {
    this.sender = sender;
  }

  // フォローされた通知
  async Followed(u: UserData, to: Snowflake) {
    const payload: FollowedNotification = {
      type: "Followed",
      payload: {
        id: u.id,
        nickName: u.nickName,
        iconImageURL: u.iconImageURL,
      },
    };
    const serialized = JSON.stringify(payload);
    await this.sender.Send(serialized, to);
    return;
  }
  // フォローが成功した通知
  async FollowSuccess(u: UserData, to: Snowflake) {
    const payload: FollowSuccessNotification = {
      type: "FollowSuccess",
      payload: {
        id: u.id,
        nickName: u.nickName,
        iconImageURL: u.iconImageURL,
      },
    };
    const serialized = JSON.stringify(payload);
    await this.sender.Send(serialized, to);
    return;
  }
  // リアクションされた通知
  async Reaction(u: UserData, to: Snowflake) {
    const payload: ReactionNotification = {
      type: "Reaction",
      payload: {
        id: u.id,
        nickName: u.nickName,
        iconImageURL: u.iconImageURL,
      },
    };
    const serialized = JSON.stringify(payload);
    await this.sender.Send(serialized, to);
    return;
  }
}
