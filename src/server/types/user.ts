import { UserFollowEventData } from "../../service/data/user.js";

export interface UserResponse {
  id: string;
  host: string;
  nickName: string;
  role: number;
  bio: string;
  headerImageURL: string;
  iconImageURL: string;
  following: Array<object>;
  softwareName: string;
}

export interface UserFollowResponse {
  following: {
    id: string;
    nickName: string;
    fullHandle: string;
    iconImageURL: string;
    bio: string;
  };
  follower: {
    id: string;
    nickName: string;
    fullHandle: string;
    iconImageURL: string;
    bio: string;
  };
}

export function DomainToUserFollowResponse(
  d: UserFollowEventData,
): UserFollowResponse {
  return {
    following: {
      id: d.following.id,
      nickName: d.following.nickName,
      fullHandle: d.following.fullHandle,
      iconImageURL: d.following.iconImageURL,
      bio: d.following.bio,
    },
    follower: {
      id: d.follower.id,
      nickName: d.follower.nickName,
      fullHandle: d.follower.fullHandle,
      iconImageURL: d.follower.iconImageURL,
      bio: d.follower.bio,
    },
  };
}
