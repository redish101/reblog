"use client";

import React from "react";
import { Spinner } from "@fluentui/react-components";

interface GlobalLoadingProps {
  isVisible: boolean;
}

export function GlobalLoading({ isVisible }: GlobalLoadingProps) {
  if (!isVisible) return null;

  return (
    <div
      style={{
        position: "fixed",
        top: 0,
        left: 0,
        right: 0,
        bottom: 0,
        backgroundColor: "rgba(255, 255, 255, 0.8)",
        backdropFilter: "blur(4px)",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        flexDirection: "column",
        gap: "16px",
        zIndex: 9999,
      }}
    >
      <Spinner size="large" />
      <div style={{ fontSize: "14px", color: "#666" }}>正在加载...</div>
    </div>
  );
}
