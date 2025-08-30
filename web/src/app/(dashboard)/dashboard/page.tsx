"use client";

import { Button, Title1, Body1 } from "@fluentui/react-components";

export default function Dashboard() {
  return (
    <div style={{ padding: "24px" }}>
      <Title1>Dashboard</Title1>
      <Body1 style={{ marginBottom: "16px" }}>reblog</Body1>
      <Button appearance="primary">主要操作</Button>
    </div>
  );
}
