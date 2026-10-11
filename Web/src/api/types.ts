export interface ApiResponse<T> {
  code: number;
  msg: string;
  data: T;
}

export interface UserInfo {
  id: number;
  phone: number;
  username: string;
  nickname: string;
  email: string;
  avatar: string;
  gender: number;
}

export interface LoginResult {
  token: string;
  user: UserInfo;
}
