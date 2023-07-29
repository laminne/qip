import { Snowflake } from "../../helpers/id_generator.js";
import { Server } from "../../domain/server.js";

export class ServerData {
  get id(): Snowflake {
    return this._id;
  }

  get host(): string {
    return this._host;
  }

  get softwareName(): string {
    return this._softwareName;
  }

  get softwareVersion(): string {
    return this._softwareVersion;
  }

  get name(): string {
    return this._name;
  }

  get description(): string {
    return this._description;
  }

  get maintainer(): string {
    return this._maintainer;
  }

  get maintainerEmail(): string {
    return this._maintainerEmail;
  }

  get iconURL(): string {
    return this._iconURL;
  }

  get faviconURL(): string {
    return this._faviconURL;
  }

  // ID
  private readonly _id: Snowflake;
  // サーバーのFQDN
  private readonly _host: string;
  // ソフトウェア名
  private readonly _softwareName: string;
  // ソフトウェアバージョン
  private readonly _softwareVersion: string;
  // サーバー名
  private readonly _name: string;
  // サーバー説明
  private readonly _description: string;
  // メンテナー
  private readonly _maintainer: string;
  // メンテナー連絡先
  private readonly _maintainerEmail: string;
  // サーバーアイコンURL
  private readonly _iconURL: string;
  // サーバーファビコンURL
  private readonly _faviconURL: string;

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

  public toDomain() {
    return new Server({
      description: this._description,
      faviconURL: this._faviconURL,
      host: this._host,
      iconURL: this._iconURL,
      id: this._id,
      maintainer: this._maintainer,
      maintainerEmail: this._maintainerEmail,
      name: this._name,
      softwareName: this._softwareName,
      softwareVersion: this._softwareVersion,
    });
  }
}

export function ServerToServerData(s: Server) {
  return new ServerData({
    description: s.description,
    faviconURL: s.faviconURL,
    host: s.host,
    iconURL: s.iconURL,
    id: s.id,
    maintainer: s.maintainer,
    maintainerEmail: s.maintainerEmail,
    name: s.name,
    softwareName: s.softwareName,
    softwareVersion: s.softwareVersion,
  });
}
