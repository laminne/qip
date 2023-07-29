import { Snowflake } from "../../helpers/id_generator.js";
import { User, UserAPData, UserFollowEvent } from "../../domain/user.js";

export class UserData {
  get id(): Snowflake {
    return this._id;
  }

  get handle(): string {
    return this._handle;
  }

  get fullHandle(): string {
    return this._fullHandle;
  }

  get serverID(): Snowflake {
    return this._serverID;
  }

  get nickName(): string {
    return this._nickName;
  }

  get role(): number {
    return this._role;
  }

  get bio(): string {
    return this._bio;
  }

  get headerImageURL(): string {
    return this._headerImageURL;
  }

  get iconImageURL(): string {
    return this._iconImageURL;
  }

  get password(): string {
    return this._password;
  }

  get isLocalUser(): boolean {
    return this._isLocalUser;
  }

  get createdAt(): Date {
    return this._createdAt;
  }

  get apData(): UserAPDataData {
    return this._apData;
  }

  get following(): Array<UserFollowEventData> {
    return [...this._following];
  }

  // id
  private readonly _id: Snowflake;
  // @ユーザー名(ユーザーハンドル)
  private readonly _handle: string;
  // @handle@host (フルハンドル)
  private readonly _fullHandle: string;
  // ユーザーが属するサーバーID
  private readonly _serverID: Snowflake;
  // 表示名
  private _nickName: string;
  // ユーザーロール: 0 normal/1 admin
  private _role: number;
  // bio
  private _bio: string;
  // ヘッダー画像URL
  private _headerImageURL: string;
  // アイコン画像URL
  private _iconImageURL: string;
  // パスワード
  private _password: string;
  // ローカルユーザーか
  private readonly _isLocalUser: boolean;
  // 作成日時
  private _createdAt: Date;
  // APのデータ
  private _apData: UserAPDataData;
  // フォロー中のユーザー
  private _following: Set<UserFollowEventData>;

  constructor(args: {
    id: Snowflake;
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
    apData: UserAPDataData;
    following: Array<UserFollowEventData>;
  }) {
    // 不変
    this._id = args.id;
    this._handle = args.handle;
    this._fullHandle = args.fullHandle;
    this._isLocalUser = args.isLocalUser;
    this._serverID = args.serverID;

    this._nickName = args.nickName;
    this._role = args.role;
    this._bio = args.bio;
    this._password = args.password;
    this._headerImageURL = args.headerImageURL;
    this._iconImageURL = args.iconImageURL;
    this._createdAt = args.createdAt;
    this._apData = args.apData;
    this._following = new Set(args.following);
  }

  public toDomain() {
    return new User({
      id: this._id,
      serverID: this._serverID,
      bio: this._bio,
      handle: this._handle,
      fullHandle: this._fullHandle,
      headerImageURL: this._headerImageURL,
      iconImageURL: this._iconImageURL,
      isLocalUser: this._isLocalUser,
      nickName: this._nickName,
      password: this._password,
      role: this._role,
      createdAt: this._createdAt,
      following: [...this._following].map((v) => v.toDomain()),
      apData: this._apData.toDomain(),
    });
  }
}

export function UserToUserData(v: User): UserData {
  try {
    return new UserData({
      id: v.id,
      serverID: v.serverID,
      bio: v.bio,
      fullHandle: v.fullHandle,
      handle: v.handle,
      headerImageURL: v.headerImageURL,
      iconImageURL: v.iconImageURL,
      isLocalUser: v.isLocalUser,
      nickName: v.nickName,
      password: v.password ?? "",
      role: v.isAdmin ? 1 : 0,
      following: v.following().map<UserFollowEventData>((e) => {
        return new UserFollowEventData(
          new UserData({
            id: e.follower.id,
            serverID: e.follower.serverID,
            bio: e.follower.bio,
            fullHandle: e.follower.fullHandle,
            handle: e.follower.handle,
            headerImageURL: e.follower.headerImageURL,
            iconImageURL: e.follower.iconImageURL,
            isLocalUser: e.follower.isLocalUser,
            nickName: e.follower.nickName,
            role: e.follower.isAdmin ? 1 : 0,
            createdAt: e.follower.createdAt,

            // 以下のデータはフォロー関係を示すのには必要ないので空欄
            password: "",
            following: new Array<UserFollowEventData>(),
            apData: new UserAPDataData({
              userID: e.following.id,
              userAPID: "",
              followersURL: "",
              followingURL: "",
              inboxURL: "",
              outboxURL: "",
              privateKey: null,
              publicKey: "",
            }),
          }),
          new UserData({
            id: e.following.id,
            serverID: e.following.serverID,
            bio: e.following.bio,
            fullHandle: e.following.fullHandle,
            handle: e.following.handle,
            headerImageURL: e.following.headerImageURL,
            iconImageURL: e.following.iconImageURL,
            isLocalUser: e.following.isLocalUser,
            nickName: e.following.nickName,
            role: e.following.isAdmin ? 1 : 0,
            createdAt: e.following.createdAt,

            // 以下のデータはフォロー関係を示すのには必要ないので空欄
            password: "",
            following: new Array<UserFollowEventData>(),
            apData: new UserAPDataData({
              userID: e.following.id,
              userAPID: "",
              followersURL: "",
              followingURL: "",
              inboxURL: "",
              outboxURL: "",
              privateKey: null,
              publicKey: "",
            }),
          }),
        );
      }),
      apData: new UserAPDataData({
        followersURL: "",
        followingURL: "",
        inboxURL: "",
        outboxURL: "",
        privateKey: null,
        publicKey: "",
        userAPID: "",
        userID: v.id,
      }),
      createdAt: v.createdAt,
    });
  } catch (e: unknown) {
    console.log(e);
    throw new Error(e as any);
  }
}

export class UserAPDataData {
  get userID(): Snowflake {
    return this._userID;
  }

  get userAPID(): string {
    return this._userAPID;
  }

  get inboxURL(): string {
    return this._inboxURL;
  }

  get outboxURL(): string {
    return this._outboxURL;
  }

  get followersURL(): string {
    return this._followersURL;
  }

  get followingURL(): string {
    return this._followingURL;
  }

  get publicKey(): string {
    return this._publicKey;
  }

  get privateKey(): string | null {
    return this._privateKey;
  }

  private readonly _userID: Snowflake;
  private readonly _userAPID: string;
  private _inboxURL: string;
  private _outboxURL: string;
  private _followersURL: string;
  private _followingURL: string;
  private readonly _publicKey: string;
  private readonly _privateKey: string | null;

  constructor(args: {
    userID: Snowflake;
    userAPID: string;
    inboxURL: string;
    outboxURL: string;
    followersURL: string;
    followingURL: string;
    publicKey: string;
    privateKey: string | null;
  }) {
    this._userID = args.userID;
    this._userAPID = args.userAPID;
    this._inboxURL = args.inboxURL;
    this._outboxURL = args.outboxURL;
    this._followersURL = args.followersURL;
    this._followingURL = args.followingURL;
    this._publicKey = args.publicKey;
    this._privateKey = args.privateKey;
  }

  public toDomain() {
    return new UserAPData({
      userID: this._userID,
      userAPID: this._userAPID,
      followersURL: this._followersURL,
      followingURL: this._followingURL,
      inboxURL: this._inboxURL,
      outboxURL: this._outboxURL,
      privateKey: this._privateKey,
      publicKey: this._publicKey,
    });
  }
}

export function UserAPDataToUserAPDataData(v: UserAPData) {
  return new UserAPDataData({
    userID: v.userID,
    userAPID: v.userAPID,
    inboxURL: v.inboxURL,
    outboxURL: v.outboxURL,
    followersURL: v.followersURL,
    followingURL: v.followingURL,
    privateKey: v.privateKey,
    publicKey: v.publicKey,
  });
}

export class UserFollowEventData {
  // フォローされたユーザー(dst)
  private readonly _follower: UserData;
  // フォローしたユーザー(from)
  private readonly _following: UserData;

  constructor(following: UserData, follower: UserData) {
    this._follower = follower;
    this._following = following;
  }

  get follower() {
    return this._follower;
  }

  get following() {
    return this._following;
  }

  public toDomain(): UserFollowEvent {
    return new UserFollowEvent(
      this._following.toDomain(),
      this._follower.toDomain(),
    );
  }
}
