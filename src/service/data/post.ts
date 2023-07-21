import { Snowflake } from "../../helpers/id_generator";
import { Media } from "../../domain/media";
import { Post, PostReactionEvent } from "../../domain/post";

export class PostData {
  private readonly _id: Snowflake;
  private readonly _authorID: Snowflake;
  private readonly _text: string;
  private readonly _visibility: number;
  private readonly _createdAt: Date;
  private readonly _attachments: Set<Media>;
  private _reactions: Set<PostReactionEvent>;
  constructor(args: {
    id: Snowflake;
    authorID: Snowflake;
    text: string;
    visibility: number;
    createdAt: Date;
    attachments: Array<Media>;
    reactions: Array<PostReactionEvent>;
  }) {
    this._id = args.id;
    this._authorID = args.authorID;

    this._text = args.text;
    this._visibility = args.visibility;
    this._createdAt = args.createdAt;
    this._attachments = new Set<Media>(args.attachments);
    this._reactions = new Set<PostReactionEvent>(args.reactions);
  }

  get id(): Snowflake {
    return this._id;
  }

  get authorID(): Snowflake {
    return this._authorID;
  }

  get text(): string {
    return this._text;
  }

  get visibility(): number {
    return this._visibility;
  }

  get createdAt(): Date {
    return this._createdAt;
  }

  get attachments(): Array<Media> {
    return [...this._attachments];
  }

  get reactions(): Array<PostReactionEvent> {
    return [...this._reactions];
  }

  public toDomain() {
    return new Post({
      id: this._id,
      authorID: this._authorID,
      attachments: [...this._attachments],
      text: this._text,
      visibility: this._visibility,
      createdAt: this._createdAt,
      reactions: [...this._reactions],
    });
  }
}

export function PostToPostData(v: Post) {
  return new PostData({
    id: v.id,
    authorID: v.authorID,
    text: v.text,
    visibility: v.visibility,
    createdAt: v.createdAt,
    reactions: v.reactions(),
    attachments: v.attachments,
  });
}
