// ユーザー
import { Snowflake } from "../helpers/id_generator.js";
import logger from "../helpers/logger.js";

export class User {
  get id(): Snowflake {
    return this._id;
  }

  get handle(): string {
    return this._handle;
  }

  get fullHandle(): string {
    return this._fullHandle;
  }

  get nickName(): string {
    return this._nickName;
  }

  set nickName(value: string) {
    this._nickName = value;
  }

  get isAdmin(): boolean {
    return this._role == 1;
  }

  public toAdmin(): void {
    this._role = 1;
  }

  public toNormal(): void {
    this._role = 0;
  }

  get bio(): string {
    return this._bio;
  }

  set bio(value: string) {
    this._bio = value;
  }

  get headerImageURL(): string {
    return this._headerImageURL;
  }

  set headerImageURL(value: string) {
    this._headerImageURL = value;
  }

  get password(): string | null {
    return this._password;
  }

  get iconImageURL(): string {
    return this._iconImageURL;
  }

  set iconImageURL(value: string) {
    this._iconImageURL = value;
  }

  get isLocalUser(): boolean {
    return this._isLocalUser;
  }

  get createdAt(): Date {
    return this._createdAt;
  }

  set createdAt(value: Date) {
    this._createdAt = value;
  }

  get apData(): UserAPData {
    return this._apData;
  }

  set apData(value: UserAPData) {
    this._apData = value;
  }

  get serverID(): Snowflake {
    return this._serverID;
  }

  public follow(u: User): void {
    // すでにフォローしていない状態なら
    if (!this._following.has(new UserFollowEvent(this, u))) {
      this._following.add(new UserFollowEvent(this, u));
    }
  }

  public unFollow(u: User): void {
    // フォローしていたらフォローを消す
    if (this._following.has(new UserFollowEvent(this, u))) {
      this._following.delete(new UserFollowEvent(this, u));
    }
  }

  public following() {
    //　FIXME: ここでなぜか {followerID: string, followingID: string}に型が変わってしまっている
    logger.debug([...this._following][0], "domain");
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
  private _apData: UserAPData;
  // フォロー中のユーザー
  private _following: Set<UserFollowEvent>;

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
    apData: UserAPData;
    following: Array<UserFollowEvent>;
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
}

export class UserAPData {
  set outboxURL(value: string) {
    this._outboxURL = value;
  }

  set followersURL(value: string) {
    this._followersURL = value;
  }

  set followingURL(value: string) {
    this._followingURL = value;
  }

  set inboxURL(value: string) {
    this._inboxURL = value;
  }

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
}

export class UserFollowEvent {
  // フォローされたユーザー(dst)
  private readonly _follower: User;
  // フォローしたユーザー(from)
  private readonly _following: User;

  constructor(following: User, follower: User) {
    this._follower = follower;
    this._following = following;
  }

  get follower(): User {
    return this._follower;
  }

  get following(): User {
    return this._following;
  }
}
