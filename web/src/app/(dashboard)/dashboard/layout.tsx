"use client";

import NoSSR from "@/components/no-ssr";
import { lightTheme } from "@/lib/theme";
import {
  FluentProvider,
  makeStyles,
  Toaster,
} from "@fluentui/react-components";

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
          <FluentProvider theme={lightTheme}>{children}</FluentProvider>
        </NoSSR>
      </body>
    </html>
  );
}
