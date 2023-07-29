import { IMediaRepository } from "../../repository/media.js";
import { Snowflake } from "../../helpers/id_generator.js";
import { Failure, Success } from "../../helpers/result.js";
import { MediaToMediaData } from "../data/media.js";
import { Media } from "../../domain/media.js";

export class FindMediaService {
  private readonly repository: IMediaRepository;

  constructor(repository: IMediaRepository) {
    this.repository = repository;
  }

  public async findByID(id: Snowflake) {
    const res = await this.repository.FindByID(id);
    if (res.isFailure()) {
      return new Failure(new Error("failed to find media by id", res.value));
    }

    return new Success(MediaToMediaData(res.value));
  }

  public async findByUserID(id: Snowflake) {
    const res = await this.repository.FindByUserID(id);
    if (res.isFailure()) {
      return new Failure(
        new Error("failed to find media by userID", res.value),
      );
    }

    const resp = res.value.map((v: Media) => MediaToMediaData(v));
    return new Success(resp);
  }

  public async findByPostID(id: Snowflake) {
    const res = await this.repository.FindByPostID(id);
    if (res.isFailure()) {
      return new Failure(
        new Error("failed to find media by postID", res.value),
      );
    }

    const resp = res.value.map((v: Media) => MediaToMediaData(v));

    return new Success(resp);
  }
}
