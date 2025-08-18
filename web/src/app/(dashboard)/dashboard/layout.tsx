"use client";

import NoSSR from "@/components/no-ssr";
import { FluentProvider, webLightTheme } from "@fluentui/react-components";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="zh-CN">
      <body>
          <NoSSR>
            <FluentProvider theme={webLightTheme}>
            {children}
          </FluentProvider>
          </NoSSR>
      </body>
    </html>
  );
}
