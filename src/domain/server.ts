import { Snowflake } from "../helpers/id_generator";

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

  constructor(
    id: Snowflake,
    host: string,
    softwareName: string,
    softwareVersion: string,
    name: string,
    description: string,
    maintainer: string,
    maintainerEmail: string,
    iconURL: string,
    faviconURL: string,
  ) {
    this._id = id;
    this._host = host;
    this._softwareName = softwareName;
    this._softwareVersion = softwareVersion;
    this._name = name;
    this._description = description;
    this._maintainer = maintainer;
    this._maintainerEmail = maintainerEmail;
    this._iconURL = iconURL;
    this._faviconURL = faviconURL;
  }
}
