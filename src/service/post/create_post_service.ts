import { Post } from "../../domain/post";
import { Failure, Success } from "../../helpers/result";
import { IPostRepository } from "../../repository/post";
import { PostToPostData } from "../data/post";

export class CreatePostService {
  private readonly repository: IPostRepository;
  constructor(repository: IPostRepository) {
    this.repository = repository;
  }

  async Handle(p: Post) {
    const res = await this.repository.Create(p);
    if (res.isFailure()) {
      return new Failure(new Error("failed to create post", res.value));
    }
    return new Success(PostToPostData(res.value));
  }
}
