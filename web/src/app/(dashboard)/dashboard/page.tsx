"use client";

import { Button, Title1, Body1, Card } from "@fluentui/react-components";
import { withRequiredAuth } from "@/hoc/with-auth";
import { useAuth } from "@/contexts/auth-context";

function Dashboard() {
  const { user, logout } = useAuth();

  return (
    <div style={{ padding: "24px" }}>
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          marginBottom: "24px",
        }}
      >
        <Title1>Dashboard</Title1>
        <Button appearance="secondary" onClick={logout}>
          登出
        </Button>
      </div>

      <Card style={{ padding: "16px", marginBottom: "16px" }}>
        <Body1>欢迎，{user?.nickname || user?.username}！</Body1>
        <Body1 style={{ marginTop: "8px", color: "#666" }}>
          邮箱：{user?.email}
        </Body1>
      </Card>

      <Body1 style={{ marginBottom: "16px" }}>reblog 管理面板</Body1>
      <Button appearance="primary">主要操作</Button>
    </div>
  );
}

// 使用HOC包装组件，要求用户必须登录
export default withRequiredAuth(Dashboard);
