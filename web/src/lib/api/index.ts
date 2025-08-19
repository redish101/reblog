import { Api } from "./generated";

function getApiBaseUrl(): string {
  // 开发环境
  if (process.env.NODE_ENV === "development") {
    return "http://localhost:3000/api/v1";
  }

  // 生产环境
  if (process.env.NODE_ENV === "production") {
    // 检查是否在浏览器环境（客户端）
    if (typeof window !== "undefined") {
      // 客户端：从 NEXT_PUBLIC_API_BASE_URL 读取
      return process.env.NEXT_PUBLIC_API_BASE_URL || "/api/v1";
    } else {
      // 服务端：使用内部服务地址
      return "http://backend:3000/api/v1";
    }
  }

  // 默认值（fallback）
  return "/api/v1";
}

const api = new Api({
  baseUrl: getApiBaseUrl(),
});

export { api };