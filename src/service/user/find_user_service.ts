import { IUserRepository } from "../../repository/user.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { Failure, Success } from "../../helpers/result.js";
import { UserToUserData } from "../data/user.js";

export class FindUserService {
  private repository: IUserRepository;

  constructor(repository: IUserRepository) {
    this.repository = repository;
  }

  async FindByID(id: Snowflake) {
    const res = await this.repository.FindByID(id);
    if (res.isFailure()) {
      return new Failure(new Error("failed to find user by id", res.value));
    }
    return new Success(UserToUserData(res.value));
  }

  async FindByHandle(handle: string) {
    const res = await this.repository.FindByHandle(handle);
    if (res.isFailure()) {
      console.log(res.value);
      return new Failure(new Error("failed to find user by id", res.value));
    }
    return new Success(UserToUserData(res.value));
  }
}
