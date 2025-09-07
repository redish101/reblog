"use client";

import React, { createContext, useContext, useEffect, useState } from "react";
import { api } from "@/lib/api";
import { GlobalLoading } from "@/components/global-loading";

export interface User {
  username: string;
  nickname: string;
  email: string;
}

interface AuthContextType {
  user: User | null;
  isLoading: boolean;
  isAuthenticated: boolean;
  logout: () => Promise<void>;
  checkAuthStatus: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  const checkAuthStatus = async () => {
    try {
      setIsLoading(true);
      // 调用 /auth/me 端点获取用户信息
      const response = await api.auth.getAuth();

      if (response.ok) {
        const userData = response.data;
        setUser({
          username: userData.username || "未知用户",
          nickname: userData.nickname || "未知用户",
          email: userData.email || "未知邮箱",
        });
      } else {
        // 未认证或其他错误
        setUser(null);
      }
    } catch (error) {
      console.log("[AUTH] 认证未通过: ", error);
      setUser(null);
    } finally {
      setIsLoading(false);
    }
  };

  const logout = async () => {
    try {
      await api.auth.logoutCreate();
      setUser(null);
      // 重定向到登录页面
      window.location.href = "/dashboard/login";
    } catch (error) {
      console.error("登出失败:", error);
      // 即使API调用失败，也清除本地状态
      setUser(null);
      window.location.href = "/dashboard/login";
    }
  };

  useEffect(() => {
    checkAuthStatus();
  }, []);

  const value = {
    user,
    isLoading,
    isAuthenticated: !!user,
    logout,
    checkAuthStatus,
  };

  return (
    <AuthContext.Provider value={value}>
      <GlobalLoading isVisible={isLoading} />
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}
