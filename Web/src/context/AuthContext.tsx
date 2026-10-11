import { useCallback, useState } from "react";
import type { ReactNode } from "react";
import { login as loginApi, register as registerApi } from "../api/auth";
import type { LoginParams, RegisterParams } from "../api/auth";
import type { UserInfo } from "../api/types";
import * as authStore from "../store/auth";
import { AuthContext } from "./auth-context";
import type { AuthContextValue } from "./auth-context";

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<UserInfo | null>(() => authStore.getUser());

  const login = useCallback(async (params: LoginParams): Promise<UserInfo> => {
    const result = await loginApi(params);
    authStore.setToken(result.token);
    authStore.setUser(result.user);
    setUser(result.user);
    return result.user;
  }, []);

  const register = useCallback(async (params: RegisterParams): Promise<void> => {
    await registerApi(params);
  }, []);

  const logout = useCallback((): void => {
    authStore.clearToken();
    authStore.clearUser();
    setUser(null);
  }, []);

  const value: AuthContextValue = { user, login, register, logout };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}
