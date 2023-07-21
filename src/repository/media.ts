import { Media } from "../domain/media";
import { Result } from "../helpers/result";
import { Snowflake } from "../helpers/id_generator";

export interface IMediaRepository {
  Create(m: Media): Promise<Result<Media, Error>>;
  Update(m: Media): Promise<Result<Media, Error>>;

  FindByID(id: Snowflake): Promise<Result<Media, Error>>;
  FindByPostID(id: Snowflake): Promise<Result<Array<Media>, Error>>;
  FindByUserID(id: Snowflake): Promise<Result<Array<Media>, Error>>;
}
