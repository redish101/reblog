"use client";

import React, { ComponentType, useEffect } from "react";
import { useAuth } from "@/contexts/auth-context";

interface WithAuthOptions {
  requireAuth?: boolean;
  redirectTo?: string;
}

/**
 * 高阶组件，用于页面级别的鉴权
 * @param WrappedComponent 被包装的组件
 * @param options 鉴权选项
 */
export function withAuth<P extends object>(
  WrappedComponent: ComponentType<P>,
  options: WithAuthOptions = {},
) {
  const {
    requireAuth = true,
    redirectTo = "/dashboard/login",
  } = options;

  const AuthenticatedComponent = (props: P) => {
    const { isAuthenticated, isLoading } = useAuth();

    useEffect(() => {
      if (!isLoading && requireAuth && !isAuthenticated) {
        window.location.href = redirectTo;
      }
    }, [isAuthenticated, isLoading]);

    // 如果正在加载，不渲染组件，由全局加载组件处理
    if (isLoading) {
      return null;
    }

    // 如果需要认证但未认证，不渲染组件（即将重定向）
    if (requireAuth && !isAuthenticated) {
      return null;
    }

    // 渲染原组件
    return <WrappedComponent {...props} />;
  };

  // 设置组件显示名称，方便调试
  AuthenticatedComponent.displayName = `withAuth(${WrappedComponent.displayName || WrappedComponent.name})`;

  return AuthenticatedComponent;
}

/**
 * 要求用户已登录的HOC
 */
export function withRequiredAuth<P extends object>(
  WrappedComponent: ComponentType<P>,
  redirectTo?: string,
) {
  return withAuth(WrappedComponent, { requireAuth: true, redirectTo });
}

/**
 * 可选认证的HOC（登录或未登录都可以访问）
 */
export function withOptionalAuth<P extends object>(
  WrappedComponent: ComponentType<P>,
) {
  return withAuth(WrappedComponent, { requireAuth: false });
}

/**
 * 自定义加载组件的HOC（已废弃，使用全局加载组件）
 * @deprecated 使用全局加载组件，直接使用 withAuth 即可
 */
export function withAuthAndLoading<P extends object>(
  WrappedComponent: ComponentType<P>,
  LoadingComponent: ComponentType,
  options?: WithAuthOptions,
) {
  // 忽略LoadingComponent参数，使用全局加载
  return withAuth(WrappedComponent, options);
}
