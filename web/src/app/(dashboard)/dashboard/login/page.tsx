"use client";

import { useState } from "react";
import {
  Button,
  Card,
  CardHeader,
  Field,
  Input,
  Text,
  Title3,
  Avatar,
  Divider,
  makeStyles,
  tokens,
  Spinner,
  useId,
  useToastController,
  Toast,
  ToastTitle,
  Toaster,
} from "@fluentui/react-components";
import { PersonRegular, LockClosedRegular } from "@fluentui/react-icons";

const useStyles = makeStyles({
  container: {
    minHeight: "100%",
    top: 0,
    bottom: 0,
    left: 0,
    right: 0,
    position: "absolute",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
    backgroundImage:
      'url("https://logincdn.msauth.net/shared/5/js/../images/fluent_web_light_2_145a07dcb971527a82b8.svg")', // 占位符背景
    backgroundSize: "cover",
    backgroundPosition: "center",
    "@media (max-width: 768px)": {
      backgroundImage: "none",
      backgroundColor: tokens.colorNeutralBackground1,
      padding: 0,
    },
  },
  backLink: {
    position: "absolute",
    top: tokens.spacingVerticalL,
    left: tokens.spacingHorizontalL,
    display: "flex",
    alignItems: "center",
    gap: tokens.spacingHorizontalXS,
    color: tokens.colorNeutralForeground1,
    textDecoration: "none",
    "&:hover": {
      color: tokens.colorBrandForeground1,
    },
    "@media (max-width: 768px)": {
      position: "static",
      alignSelf: "flex-start",
      marginBottom: tokens.spacingVerticalM,
      marginLeft: tokens.spacingHorizontalM,
      marginTop: tokens.spacingVerticalM,
    },
  },
  cardContainer: {
    position: "relative",
    width: "100%",
    maxWidth: "400px",
    "@media (max-width: 768px)": {
      maxWidth: "100%",
      height: "100vh",
      display: "flex",
      flexDirection: "column",
    },
  },
  card: {
    width: "100%",
    padding: (tokens.spacingVerticalXXL, tokens.spacingHorizontalXXL),
    transition: "height 0.6s ease-in-out",
    overflow: "hidden",
    "@media (max-width: 768px)": {
      border: "none",
      boxShadow: "none",
      borderRadius: "0",
      flex: "1",
      display: "flex",
      flexDirection: "column",
      justifyContent: "center",
    },
  },
  header: {
    textAlign: "center",
    marginBottom: tokens.spacingVerticalXL,
    flexDirection: "column",
  },
  title: {
    color: tokens.colorNeutralForeground1,
    marginBottom: tokens.spacingVerticalS,
    fontWeight: "normal",
    transition: "all 0.6s ease-in-out",
  },
  form: {
    display: "flex",
    flexDirection: "column",
    gap: tokens.spacingVerticalL,
    opacity: 1,
    transition: "opacity 0.6s ease-in-out",
  },
  formEntering: {
    opacity: 0,
  },
  formExiting: {
    opacity: 0,
  },
  formContainer: {
    position: "relative",
    height: "240px",
  },
  userBadge: {
    display: "flex",
    alignItems: "center",
    gap: tokens.spacingHorizontalM,
    marginBottom: tokens.spacingVerticalS,
    padding: (tokens.spacingVerticalS, tokens.spacingHorizontalS),
    backgroundColor: tokens.colorNeutralBackground2,
    borderRadius: tokens.borderRadiusMedium,
  },
  passwordToggle: {
    position: "relative",
  },
  passwordToggleButton: {
    position: "absolute",
    right: tokens.spacingHorizontalS,
    top: "50%",
    transform: "translateY(-50%)",
    background: "transparent",
    border: "none",
    cursor: "pointer",
    color: tokens.colorNeutralForeground2,
    "&:hover": {
      color: tokens.colorNeutralForeground1,
    },
  },
  submitButton: {
    marginTop: tokens.spacingVerticalM,
  },
  divider: {
    margin: `${tokens.spacingVerticalL} 0`,
  },
  githubButton: {
    width: "100%",
  },
  githubIcon: {
    width: "20px",
    height: "20px",
    paddingRight: tokens.spacingHorizontalS,
  },
  footer: {
    position: "absolute",
    bottom: tokens.spacingVerticalL,
    width: "100%",
    textAlign: "center",
    color: tokens.colorNeutralForeground2,
    fontSize: tokens.fontSizeBase200,
    "@media (max-width: 768px)": {
      position: "static",
      marginTop: tokens.spacingVerticalM,
    },
  },
});

interface LoginPageProps {}

export default function Login({}: LoginPageProps) {
  const styles = useStyles();
  const [step, setStep] = useState<"email" | "password">("email");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [isTransitioning, setIsTransitioning] = useState(false);

  const handleEmailSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);

    setTimeout(() => {
      setIsTransitioning(true);
      setTimeout(() => {
        setStep("password");
        setIsTransitioning(false);
        setIsLoading(false);
      }, 300); // 等待淡出动画完成一半再切换步骤
    }, 1000);
  };

  const handlePasswordSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);

    console.log("Login attempt:", { email, password });
    setTimeout(() => {
      window.location.href = "https://www.bilibili.com/video/BV1GJ411x7h7";
      setIsLoading(false);
    }, 1000);
  };

  const handleGitHubLogin = () => {
    // TODO: 实现GitHub OAuth登录
    console.log("GitHub login clicked");
  };

  // 返回到邮箱步骤
  const handleBackToEmail = () => {
    setIsTransitioning(true);
    setTimeout(() => {
      setStep("email");
      setPassword("");
      setIsTransitioning(false);
    }, 300); // 等待淡出动画完成一半再切换步骤
  };

  return (
    <div className={styles.container}>
      <div className={styles.cardContainer}>
        <Card className={styles.card}>
          <CardHeader
            className={styles.header}
            header={<Title3 className={styles.title}>登录</Title3>}
          ></CardHeader>

          <div className={styles.formContainer}>
            {step === "email" ? (
              <form
                onSubmit={handleEmailSubmit}
                className={`${styles.form} ${isTransitioning ? styles.formExiting : ""}`}
              >
                <Field label="邮箱地址" required>
                  <Input
                    type="email"
                    value={email}
                    onChange={(_, data) => setEmail(data.value)}
                    disabled={isLoading}
                    contentBefore={<PersonRegular />}
                    required
                  />
                </Field>

                <Button
                  type="submit"
                  appearance="primary"
                  className={styles.submitButton}
                  disabled={isLoading || !email}
                >
                  {isLoading ? <Spinner size="tiny" /> : "下一步"}
                </Button>

                <Divider className={styles.divider}>或</Divider>

                <Button
                  appearance="secondary"
                  className={styles.githubButton}
                  onClick={handleGitHubLogin}
                  disabled={isLoading}
                >
                  <svg
                    className={styles.githubIcon}
                    viewBox="0 0 24 24"
                    fill="currentColor"
                  >
                    <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
                  </svg>
                  使用 GitHub 登录
                </Button>
              </form>
            ) : (
              <form
                onSubmit={handlePasswordSubmit}
                className={`${styles.form} ${isTransitioning ? styles.formExiting : ""}`}
              >
                <div className={styles.userBadge}>
                  <Avatar name={email} size={32} />
                  <div>
                    <Text weight="semibold">{email}</Text>
                    <br />
                    <Text
                      size={200}
                      style={{ color: tokens.colorNeutralForeground2 }}
                    >
                      管理员
                    </Text>
                  </div>
                </div>

                <Field label="密码" required>
                  <Input
                    value={password}
                    type="password"
                    onChange={(_, data) => setPassword(data.value)}
                    placeholder="请输入密码"
                    disabled={isLoading}
                    contentBefore={<LockClosedRegular />}
                    required
                  />
                </Field>

                <Button
                  type="submit"
                  appearance="primary"
                  className={styles.submitButton}
                  disabled={isLoading || !password}
                  onClick={handlePasswordSubmit}
                >
                  {isLoading ? <Spinner size="tiny" /> : "登录"}
                </Button>

                <Button
                  appearance="secondary"
                  onClick={handleBackToEmail}
                  disabled={isLoading}
                >
                  上一步
                </Button>
              </form>
            )}
          </div>
        </Card>
      </div>
      <footer className={styles.footer}>
        <Text>Powered by reblog on kubernetes</Text>
      </footer>
    </div>
  );
}
