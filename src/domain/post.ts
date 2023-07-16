import { Snowflake } from "../helpers/id_generator";
import { User } from "./user";
import { Media } from "./media";

export class Post {
  get visibility(): number {
    return this._visibility;
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

  get createdAt(): Date {
    return this._createdAt;
  }

  get attachments(): Array<Media> {
    return [...this._attachments];
  }

  public addReactions(user: User) {
    if (!this._reactions.has(new PostReactionEvent(this._id, user.id))) {
      this._reactions.add(new PostReactionEvent(this._id, user.id));
    }
  }

  public reactions(): Array<PostReactionEvent> {
    return [...this._reactions];
  }

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

    this._text = this.validateText(args.text);
    this._visibility = args.visibility;
    this._createdAt = args.createdAt;
    this._attachments = new Set<Media>(args.attachments);
    this._reactions = new Set<PostReactionEvent>(args.reactions);
  }

  private validateText(text: string): string {
    if ([...text].length >= 5000) {
      return text.substring(0, 5000);
    }
    return text;
  }
}

export class PostReactionEvent {
  private readonly _postID: Snowflake;
  private readonly _userID: Snowflake;

  constructor(postID: Snowflake, userID: Snowflake) {
    this._postID = postID;
    this._userID = userID;
  }

  get postID(): Snowflake {
    return this._postID;
  }

  get userID(): Snowflake {
    return this._userID;
  }
}
