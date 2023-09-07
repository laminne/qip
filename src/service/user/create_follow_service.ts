import { Failure, Success } from "../../helpers/result.js";
import { IUserRepository } from "../../repository/user.js";
import { UserFollowEventToUserFollowEventData } from "../data/user.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { UserFollowEvent } from "../../domain/user.js";

export class CreateFollowService {
  private readonly repository: IUserRepository;

  constructor(repository: IUserRepository) {
    this.repository = repository;
  }

  // Following - フォロー > Follower
  async Handle(followingID: Snowflake, followerID: Snowflake) {
    // すでにフォローしている場合は処理を切る
    const isExists = await this.isExists(followingID, followerID);
    if (isExists) {
      return new Failure(new Error("already following"));
    }

    // ユーザーを取ってくる
    const following = await this.repository.FindByID(followingID);
    if (following.isFailure()) {
      return new Failure(new Error("failed to find user", following.value));
    }
    const follower = await this.repository.FindByID(followerID);
    if (follower.isFailure()) {
      return new Failure(new Error("failed to find user", follower.value));
    }

    const req = new UserFollowEvent(following.value, follower.value);
    const res = await this.repository.CreateFollow(req);
    if (res.isFailure()) {
      return new Failure(new Error("failed to create follow", res.value));
    }

    return new Success(UserFollowEventToUserFollowEventData(res.value));
  }

  private async isExists(followingID: Snowflake, followerID: Snowflake) {
    const obj = await this.repository.FindFollowEvent(followingID, followerID);
    return !obj.isFailure();
  }
}
