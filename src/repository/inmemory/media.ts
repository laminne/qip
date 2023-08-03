import { IMediaRepository } from "../media.js";
import { Media } from "../../domain/media.js";
import { Failure, Result, Success } from "../../helpers/result.js";
import { Snowflake } from "../../helpers/id_generator.js";

export class MediaRepository implements IMediaRepository {
  private media: Set<Media>;

  constructor(data: Media[]) {
    this.media = new Set(data);
  }

  async Create(m: Media): Promise<Result<Media, Error>> {
    try {
      this.media.add(m);
      return new Success(m);
    } catch (e: unknown) {
      return new Failure(new Error("failed to create media", e as any));
    }
  }

  async FindByID(id: Snowflake): Promise<Result<Media, Error>> {
    try {
      const res = [...this.media].filter((v) => v.id == id);
      return new Success(res[0]);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByPostID(id: Snowflake): Promise<Result<Array<Media>, Error>> {
    try {
      const res = [...this.media].filter((v) => v.postID == id);
      return new Success(res);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async FindByUserID(id: Snowflake): Promise<Result<Array<Media>, Error>> {
    try {
      const res = [...this.media].filter((v) => {
        return v.authorID == id;
      });
      console.log(res);
      return new Success(res);
    } catch (e: unknown) {
      return new Failure(new Error(e as any));
    }
  }

  async Update(m: Media): Promise<Result<Media, Error>> {
    return new Failure(new Error(""));
  }
}
