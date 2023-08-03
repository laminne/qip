export interface LoginRequest {
  name: string;
  password: string;
}

export interface LoginResponse {
  token: string;
}
