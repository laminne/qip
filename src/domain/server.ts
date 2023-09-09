import { Snowflake } from "../helpers/id_generator.js";

export interface ServerArgs {
  id: Snowflake;
  host: string;
  softwareName: string;
  softwareVersion: string;
  name: string;
  description: string;
  maintainer: string;
  maintainerEmail: string;
  iconURL: string;
  faviconURL: string;
}

export class Server {
  get softwareName(): string {
    return this._softwareName;
  }

  set softwareName(value: string) {
    this._softwareName = value;
  }

  get softwareVersion(): string {
    return this._softwareVersion;
  }

  set softwareVersion(value: string) {
    this._softwareVersion = value;
  }

  get name(): string {
    return this._name;
  }

  set name(value: string) {
    this._name = value;
  }

  get description(): string {
    return this._description;
  }

  set description(value: string) {
    this._description = value;
  }

  get maintainer(): string {
    return this._maintainer;
  }

  set maintainer(value: string) {
    this._maintainer = value;
  }

  get maintainerEmail(): string {
    return this._maintainerEmail;
  }

  set maintainerEmail(value: string) {
    this._maintainerEmail = value;
  }

  get iconURL(): string {
    return this._iconURL;
  }

  set iconURL(value: string) {
    this._iconURL = value;
  }

  get faviconURL(): string {
    return this._faviconURL;
  }

  set faviconURL(value: string) {
    this._faviconURL = value;
  }
  get id(): Snowflake {
    return this._id;
  }

  get host(): string {
    return this._host;
  }

  // ID
  private readonly _id: Snowflake;
  // サーバーのFQDN
  private readonly _host: string;
  // ソフトウェア名
  private _softwareName: string;
  // ソフトウェアバージョン
  private _softwareVersion: string;
  // サーバー名
  private _name: string;
  // サーバー説明
  private _description: string;
  // メンテナー
  private _maintainer: string;
  // メンテナー連絡先
  private _maintainerEmail: string;
  // サーバーアイコンURL
  private _iconURL: string;
  // サーバーファビコンURL
  private _faviconURL: string;

  constructor(args: {
    id: Snowflake;
    host: string;
    softwareName: string;
    softwareVersion: string;
    name: string;
    description: string;
    maintainer: string;
    maintainerEmail: string;
    iconURL: string;
    faviconURL: string;
  }) {
    this.validate(args);

    this._id = args.id;
    this._host = args.host;
    this._softwareName = args.softwareName;
    this._softwareVersion = args.softwareVersion;
    this._name = args.name;
    this._description = args.description;
    this._maintainer = args.maintainer;
    this._maintainerEmail = args.maintainerEmail;
    this._iconURL = args.iconURL;
    this._faviconURL = args.faviconURL;
  }

  private validate(args: ServerArgs) {
    if ([...args.host].length > 128) {
      throw new Error("failed to create server: host is too long");
    } else if ([...args.host].length < 4) {
      throw new Error("failed to create server: host is too short");
    }

    if ([...args.softwareName].length > 128) {
      throw new Error("failed to create server: softwareName is too long");
    }

    if ([...args.softwareVersion].length > 128) {
      throw new Error("failed to create server: softwareVersion is too long");
    }

    if ([...args.name].length > 128) {
      throw new Error("failed to create server: name is too long");
    }

    if ([...args.description].length > 3000) {
      throw new Error("failed to create server: description is too long");
    }

    if ([...args.maintainer].length > 256) {
      throw new Error("failed to create server: maintainer is too long");
    }

    if ([...args.maintainerEmail].length > 256) {
      throw new Error("failed to create server: maintainerEmail is too long");
    }
  }
}
