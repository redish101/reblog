# ReBlog Helm Chart

ReBlog 是一个前后端分离的博客系统的 Helm Chart，支持通过 Kubernetes 部署完整的博客应用。

## 功能特性

- 前后端分离架构
- 独立的 Ingress 配置（不同的域名）
- 外置数据库支持（PostgreSQL）
- 自动伸缩支持
- 完整的健康检查
- 安全的环境变量管理
- 灵活的资源配置

## 前置要求

- Kubernetes 1.19+
- Helm 3.2.0+
- 外置 PostgreSQL 数据库
- Nginx Ingress Controller（如果使用 Ingress）

## 安装

### 1. 准备配置文件

复制示例配置文件并修改：

```bash
cp values-example.yaml values-production.yaml
```

### 2. 修改配置

编辑 `values-production.yaml` 文件，至少需要配置以下必填项：

```yaml
backend:
  image:
    repository: your-registry/reblog-backend
    tag: "v1.0.0"
  env:
    databaseUrl: "postgres://username:password@your-postgres-host:5432/reblog?sslmode=require"
    secretKey: "your-super-secret-key"
    ownerEmail: "admin@yourdomain.com"
  ingress:
    host: "api.yourdomain.com"

frontend:
  image:
    repository: your-registry/reblog-frontend
    tag: "v1.0.0"
  ingress:
    host: "yourdomain.com"

ingress:
  className: "nginx"
```

### 3. 安装 Chart

```bash
helm install reblog ./helm/reblog -f values-production.yaml
```

### 4. 升级应用

```bash
helm upgrade reblog ./helm/reblog -f values-production.yaml
```

## 配置说明

### 后端配置

| 参数 | 描述 | 默认值 | 必填 |
|------|------|--------|------|
| `backend.enabled` | 是否启用后端 | `true` | 否 |
| `backend.image.repository` | 后端镜像地址 | `reblog-backend` | 是 |
| `backend.image.tag` | 后端镜像标签 | `latest` | 否 |
| `backend.env.databaseUrl` | 数据库连接字符串 | - | **是** |
| `backend.env.secretKey` | JWT 密钥 | `reblog-secret` | **是** |
| `backend.env.ownerEmail` | 管理员邮箱 | - | 否 |
| `backend.env.github.clientId` | GitHub OAuth Client ID | - | 否 |
| `backend.env.github.secret` | GitHub OAuth Secret | - | 否 |
| `backend.ingress.host` | 后端域名 | `api.reblog.local` | **是** |

### 前端配置

| 参数 | 描述 | 默认值 | 必填 |
|------|------|--------|------|
| `frontend.enabled` | 是否启用前端 | `true` | 否 |
| `frontend.image.repository` | 前端镜像地址 | `reblog-frontend` | 是 |
| `frontend.image.tag` | 前端镜像标签 | `latest` | 否 |
| `frontend.ingress.host` | 前端域名 | `reblog.local` | **是** |

**注意：** 前端的环境变量（如 `NEXT_PUBLIC_API_URL`）应该在构建 Docker 镜像时配置，而不是在运行时通过 Kubernetes 环境变量注入。

### Ingress 配置

| 参数 | 描述 | 默认值 | 必填 |
|------|------|--------|------|
| `ingress.className` | Ingress 类名 | `nginx` | 否 |

## 环境变量

### 后端环境变量

- `DEV`: 是否为开发模式
- `PORT`: 服务端口
- `DATABASE_URL`: PostgreSQL 数据库连接字符串
- `SECRET_KEY`: 用于生成 JWT 密钥对的密钥
- `OWNER_EMAIL`: 管理员邮箱
- `GITHUB_CLIENT_ID`: GitHub OAuth 客户端 ID
- `GITHUB_SECRET`: GitHub OAuth 密钥
- `GITHUB_REDIRECT_URL`: GitHub OAuth 回调 URL
- `FRONTEND_URL`: 前端 URL

### 前端环境变量（构建时）

前端使用 Next.js，其环境变量需要在构建时确定。请在构建 Docker 镜像时配置以下环境变量：

- `NEXT_PUBLIC_API_URL`: 后端 API 地址
- `NODE_ENV`: Node.js 环境（production/development）

示例 Dockerfile：

```dockerfile
FROM node:18-alpine AS builder

# 设置环境变量
ARG NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_API_URL=$NEXT_PUBLIC_API_URL

# ... 其他构建步骤
```

## 数据库准备

在部署之前，请确保：

1. PostgreSQL 数据库已创建
2. 数据库用户具有足够的权限
3. 网络连接正常

创建数据库示例：

```sql
CREATE DATABASE reblog;
CREATE USER reblog_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE reblog TO reblog_user;
```

## TLS/SSL 配置

如果需要 HTTPS 支持，可以：

1. 使用 cert-manager 自动管理证书：

```yaml
backend:
  ingress:
    annotations:
      cert-manager.io/cluster-issuer: "letsencrypt-prod"
    tls:
      - secretName: reblog-api-tls
        hosts:
          - api.yourdomain.com
```

2. 手动创建 TLS 证书 Secret：

```bash
kubectl create secret tls reblog-api-tls --cert=path/to/tls.crt --key=path/to/tls.key
```

## 监控和日志

### 健康检查

- 后端：`/healthz`
- 前端：`/`

### 查看日志

```bash
# 查看后端日志
kubectl logs -l app.kubernetes.io/name=reblog,app.kubernetes.io/component=backend

# 查看前端日志
kubectl logs -l app.kubernetes.io/name=reblog,app.kubernetes.io/component=frontend
```

## 卸载

```bash
helm uninstall reblog
```

## 故障排除

### 常见问题

1. **数据库连接失败**
   - 检查数据库连接字符串是否正确
   - 确认数据库服务是否可访问
   - 检查网络策略设置

2. **镜像拉取失败**
   - 确认镜像地址和标签是否正确
   - 检查是否需要配置 imagePullSecrets

3. **Ingress 无法访问**
   - 确认 Ingress Controller 是否正常运行
   - 检查 DNS 解析是否正确
   - 验证域名配置

## 开发和调试

### 本地开发

```bash
# 使用开发配置
helm install reblog-dev ./helm/reblog -f values-dev.yaml
```

### 配置验证

```bash
# 验证配置语法
helm lint ./helm/reblog

# 渲染模板但不安装
helm template reblog ./helm/reblog -f values-production.yaml
```
