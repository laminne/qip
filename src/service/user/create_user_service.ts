import { IUserRepository } from "../../repository/user.js";
import { User, UserAPData } from "../../domain/user.js";
import { Failure, Success } from "../../helpers/result.js";
import { UserToUserData } from "../data/user.js";
import { Snowflake, SnowflakeIDGenerator } from "../../helpers/id_generator.js";

export class CreateUserService {
  private readonly repository: IUserRepository;
  private readonly idGenerator: SnowflakeIDGenerator;

  constructor(repository: IUserRepository, idGenerator: SnowflakeIDGenerator) {
    this.repository = repository;
    this.idGenerator = idGenerator;
  }

  async Handle(u: CreateUserArgs) {
    const id = this.idGenerator.generate();
    const req = new User({
      id: id,
      serverID: u.serverID,
      handle: u.handle,
      fullHandle: u.fullHandle,
      bio: u.bio,
      headerImageURL: u.headerImageURL,
      iconImageURL: u.iconImageURL,
      isLocalUser: u.isLocalUser,
      nickName: u.nickName,
      password: u.password,
      role: u.role,
      following: [],
      createdAt: u.createdAt,
      apData: new UserAPData({
        userID: id,
        userAPID: this.idGenerator.generate(),
        followersURL: u.apData.followersURL,
        followingURL: u.apData.followingURL,
        inboxURL: u.apData.inboxURL,
        outboxURL: u.apData.outboxURL,
        privateKey: u.apData.privateKey,
        publicKey: u.apData.publicKey,
      }),
    });
    const res = await this.repository.Create(req);
    if (res.isFailure()) {
      return new Failure(new Error("failed to create user", res.value));
    }

    return new Success(UserToUserData(res.value));
  }
}

export interface CreateUserArgs {
  handle: string;
  fullHandle: string;
  serverID: Snowflake;
  nickName: string;
  role: number;
  bio: string;
  headerImageURL: string;
  iconImageURL: string;
  password: string;
  isLocalUser: boolean;
  createdAt: Date;
  apData: {
    id: string;
    inboxURL: string;
    outboxURL: string;
    followersURL: string;
    followingURL: string;
    publicKey: string;
    privateKey: string | null;
  };
}
