import { IUserRepository } from "../../repository/user.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { Failure, Success } from "../../helpers/result.js";
import { UserToUserData } from "../data/user.js";
import logger from "../../helpers/logger.js";

export class FindUserService {
  private repository: IUserRepository;

  constructor(repository: IUserRepository) {
    this.repository = repository;
  }

  async FindByID(id: Snowflake) {
    const res = await this.repository.FindByID(id);
    if (res.isFailure()) {
      return new Failure(res.value);
    }
    const resp = UserToUserData(res.value);
    return new Success(resp);
  }

  async FindByHandle(handle: string) {
    const res = await this.repository.FindByHandle(handle);
    if (res.isFailure()) {
      logger.error(res.value);
      return new Failure(res.value);
    }
    const resp = UserToUserData(res.value);
    return new Success(resp);
  }
}
