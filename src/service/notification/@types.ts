export interface CommonNotificationObject<T, U> {
  type: T;
  payload: U;
}

export type FollowedNotification = CommonNotificationObject<
  "Followed",
  FollowedNotificationPayload
>;

export interface FollowedNotificationPayload {
  id: string;
  nickName: string;
  iconImageURL: string;
}

export type FollowSuccessNotification = CommonNotificationObject<
  "FollowSuccess",
  FollowSuccessNotificationPayload
>;
export interface FollowSuccessNotificationPayload {
  id: string;
  nickName: string;
  iconImageURL: string;
}

export type ReactionNotification = CommonNotificationObject<
  "Reaction",
  ReactionNotificationPayload
>;
export interface ReactionNotificationPayload {
  id: string;
  nickName: string;
  iconImageURL: string;
}
