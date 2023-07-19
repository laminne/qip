import { Post } from "../../domain/post";
import { Failure, Success } from "../../helpers/result";
import { IPostRepository } from "../../repository/post";
import { PostToPostData } from "../data/post";
import { Snowflake, SnowflakeIDGenerator } from "../../helpers/id_generator";
import { Media } from "../../domain/media";

export class CreatePostService {
  private readonly repository: IPostRepository;
  private readonly idGenerator: SnowflakeIDGenerator;
  constructor(repository: IPostRepository, idGenerator: SnowflakeIDGenerator) {
    this.repository = repository;
    this.idGenerator = idGenerator;
  }

  async Handle(p: CreatePostArgs) {
    const id = this.idGenerator.generate();
    const req = new Post({
      id: id,
      attachments: p.attachments.map((v) => {
        return new Media({
          id: this.idGenerator.generate(),
          authorID: p.authorID,
          blurhash: v.blurHash,
          cached: v.cached,
          isSensitive: v.isSensitive,
          md5Sum: v.md5Sum,
          name: v.name,
          postID: id,
          size: v.size,
          thumbnailURL: v.thumbnailURL,
          type: v.type,
          url: v.url,
        });
      }),
      authorID: p.authorID,
      createdAt: new Date(),
      reactions: [],
      text: p.text,
      visibility: p.visibility,
    });
    const res = await this.repository.Create(req);
    if (res.isFailure()) {
      return new Failure(new Error("failed to create post", res.value));
    }
    return new Success(PostToPostData(res.value));
  }
}

export interface CreatePostArgs {
  authorID: Snowflake;
  text: string;
  visibility: number;
  attachments: [
    {
      name: string;
      type: string;
      md5Sum: string;
      size: number;
      isSensitive: boolean;
      blurHash: string;
      url: string;
      thumbnailURL: string;
      cached: boolean;
    },
  ];
}
