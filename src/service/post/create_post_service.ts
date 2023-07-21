import { Post } from "../../domain/post";
import { AsyncResult, Failure, Success } from "../../helpers/result";
import { IPostRepository } from "../../repository/post";
import { PostData, PostToPostData } from "../data/post";
import { Snowflake, SnowflakeIDGenerator } from "../../helpers/id_generator";
import { Media } from "../../domain/media";

export class CreatePostService {
  private readonly repository: IPostRepository;
  private readonly idGenerator: SnowflakeIDGenerator;
  constructor(repository: IPostRepository, idGenerator: SnowflakeIDGenerator) {
    this.repository = repository;
    this.idGenerator = idGenerator;
  }

  async Handle(p: CreatePostArgs): AsyncResult<PostData, Error> {
    const id = this.idGenerator.generate();
    const req = new Post({
      id: id,
      attachments: new Array<Media>(),
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
  attachments: Array<Snowflake>;
}
