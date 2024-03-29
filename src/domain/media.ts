import { Snowflake } from "../helpers/id_generator.js";

export class Media {
  get id(): Snowflake {
    return this._id;
  }

  get authorID(): Snowflake {
    return this._authorID;
  }

  get postID(): Snowflake {
    return this._postID;
  }

  get name(): string {
    return this._name;
  }

  get type(): string {
    return this._type;
  }

  get md5Sum(): string {
    return this._md5Sum;
  }

  get size(): number {
    return this._size;
  }

  get isSensitive(): boolean {
    return this._isSensitive;
  }

  get blurhash(): string {
    return this._blurhash;
  }

  get url(): string {
    return this._url;
  }

  get thumbnailURL(): string {
    return this._thumbnailURL;
  }

  get cached(): boolean {
    return this._cached;
  }

  constructor(args: {
    id: Snowflake;
    authorID: Snowflake;
    postID: Snowflake;
    name: string;
    type: string;
    md5Sum: string;
    size: number;
    isSensitive: boolean;
    blurhash: string;
    url: string;
    thumbnailURL: string;
    cached: boolean;
  }) {
    this._id = args.id;
    this._authorID = args.authorID;
    this._postID = args.postID;
    this._name = args.name;
    this._type = args.type;
    this._md5Sum = args.md5Sum;
    this._size = args.size;
    this._isSensitive = args.isSensitive;
    this._blurhash = args.blurhash;
    this._url = args.url;
    this._thumbnailURL = args.thumbnailURL;
    this._cached = args.cached;
  }
  // ID
  private readonly _id: Snowflake;
  // 投稿者
  private readonly _authorID: Snowflake;
  // 紐づく先の投稿
  private readonly _postID: Snowflake;
  // ファイル名
  private readonly _name: string;
  // mimeタイプ
  private readonly _type: string;
  // MD5サム
  private readonly _md5Sum: string;
  // ファイルサイズ
  private readonly _size: number;
  // NSFWか
  private readonly _isSensitive: boolean;
  // blurhash
  private readonly _blurhash: string;
  // ファイルのURL
  private readonly _url: string;
  // サムネイルのURL
  private readonly _thumbnailURL: string;
  // ストレージにキャッシュされているか
  private readonly _cached: boolean;
}
