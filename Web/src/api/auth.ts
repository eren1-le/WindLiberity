import { request } from "./request";
import type { LoginResult } from "./types";

export interface LoginParams {
  username: string;
  password: string;
}

export interface RegisterParams {
  phone: string;
  username: string;
  password: string;
  confirm_password: string;
}

export function login(params: LoginParams): Promise<LoginResult> {
  return request<LoginResult>("/login", {
    method: "POST",
    body: JSON.stringify(params),
  });
}

export function register(params: RegisterParams): Promise<Record<string, unknown>> {
  return request<Record<string, unknown>>("/reg", {
    method: "POST",
    body: JSON.stringify(params),
  });
}
