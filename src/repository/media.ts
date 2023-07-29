import { Media } from "../domain/media.js";
import { Result } from "../helpers/result.js";
import { Snowflake } from "../helpers/id_generator.js";

export interface IMediaRepository {
  Create(m: Media): Promise<Result<Media, Error>>;
  Update(m: Media): Promise<Result<Media, Error>>;

  FindByID(id: Snowflake): Promise<Result<Media, Error>>;
  FindByPostID(id: Snowflake): Promise<Result<Array<Media>, Error>>;
  FindByUserID(id: Snowflake): Promise<Result<Array<Media>, Error>>;
}
