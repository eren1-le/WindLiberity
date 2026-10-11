import { createContext, useContext } from "react";
import type { LoginParams, RegisterParams } from "../api/auth";
import type { UserInfo } from "../api/types";

export interface AuthContextValue {
  user: UserInfo | null;
  login: (params: LoginParams) => Promise<UserInfo>;
  register: (params: RegisterParams) => Promise<void>;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextValue | undefined>(undefined);

export function useAuth(): AuthContextValue {
  const ctx = useContext(AuthContext);
  if (!ctx) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return ctx;
}
