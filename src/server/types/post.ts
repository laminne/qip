export interface CommonPostRequest {
  text: string;
  visibility: number;
  attachments: Array<string>;
}

export interface CommonPostResponse {
  id: string;
  text: string;
  author: {
    id: string;
    nickName: string;
    host: string;
    iconImageURL: string;
  };
  createdAt: Date;
  attachments: Array<Omit<CommonMediaResponse, "postID" | "md5Sum" | "cached">>;
  reactions: Array<PostReactionResponse>;
}

export interface CommonMediaResponse {
  id: string;
  authorID: string;
  postID: string;
  name: string;
  type: string;
  md5Sum: string;
  size: number;
  isSensitive: boolean;
  blurhash: string;
  url: string;
  thumbnailURL: string;
  cached: boolean;
}

export interface PostReactionResponse {
  postID: string;
  userID: string;
}
