"use client";

import NoSSR from "@/components/no-ssr";
import { lightTheme } from "@/lib/theme";
import { AuthProvider } from "@/contexts/auth-context";
import { FluentProvider, makeStyles } from "@fluentui/react-components";

const styles = makeStyles({
  root: {
    margin: 0,
    padding: 0,
  },
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="zh-CN">
      <body className={styles().root}>
        <NoSSR>
          <FluentProvider theme={lightTheme}>
            <AuthProvider>{children}</AuthProvider>
          </FluentProvider>
        </NoSSR>
      </body>
    </html>
  );
}
