import { IMediaRepository } from "../../repository/media";
import { Snowflake } from "../../helpers/id_generator";
import { Failure, Success } from "../../helpers/result";
import { MediaToMediaData } from "../data/media";

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

    return new Success(
      res.value.map((v) => {
        MediaToMediaData(v);
      }),
    );
  }
  public async findByPostID(id: Snowflake) {
    const res = await this.repository.FindByPostID(id);
    if (res.isFailure()) {
      return new Failure(
        new Error("failed to find media by postID", res.value),
      );
    }

    return new Success(
      res.value.map((v) => {
        MediaToMediaData(v);
      }),
    );
  }
}
