import { IUserRepository } from "../../repository/user";
import { User } from "../../domain/user";
import { Failure, Success } from "../../helpers/result";
import { UserToUserData } from "../data/user";

export class CreateUserService {
  private readonly repository: IUserRepository;

  constructor(repository: IUserRepository) {
    this.repository = repository;
  }

  async Handle(u: User) {
    const res = await this.repository.Create(u);
    if (res.isFailure()) {
      return new Failure(new Error("failed to create user", res.value));
    }

    return new Success(UserToUserData(u));
  }
}
