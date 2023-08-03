export interface UserResponse {
  id: string;
  host: string;
  nickName: string;
  role: number;
  bio: string;
  headerImageURL: string;
  iconImageURL: string;
  following: Array<object>;
  softwareName: string;
}
