package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/redish101/reblog/internal/env"
	"github.com/redish101/reblog/internal/jwt"
	"github.com/redish101/reblog/server/common"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// GitHubUser GitHub 用户信息结构
type GitHubUser struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

// GitHubEmail GitHub 用户邮箱信息结构
type GitHubEmail struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

// GitHubTokenResponse GitHub token 响应结构
type GitHubTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

//	@summary		GitHub OAuth 登录
//	@description	跳转到 GitHub OAuth 授权页面
//	@tags			auth
//	@accept			json
//	@produce		json
//	@param			redirect_uri	query		string	false	"登录成功后的重定向地址"
//	@success		302				{string}	string	"重定向到 GitHub 授权页面"
//	@router			/auth/github [get]
//
// GitHubLogin 跳转到 GitHub OAuth 授权页面
func (h *AuthHandler) GitHubLogin(ctx context.Context, c *app.RequestContext) {
	// 构建 GitHub OAuth 授权 URL
	authURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=%s",
		env.GitHubClientID,
		env.GitHubRedirectURL,
		"user:email", // 请求用户基本信息和邮箱权限
	)

	// 如果有前端重定向地址，可以通过 state 参数传递
	redirectURI := c.Query("redirect_uri")
	if redirectURI != "" {
		authURL += "&state=" + url.QueryEscape(string(redirectURI))
	}

	// 重定向到 GitHub 授权页面
	c.Redirect(302, []byte(authURL))
}

type GithubCallBackResponse struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

//	@summary		GitHub OAuth 回调
//	@description	处理 GitHub OAuth 回调，获取用户信息并生成 JWT token
//	@tags			auth
//	@accept			json
//	@produce		json
//	@param			code	query		string					true	"GitHub OAuth 授权码"
//	@param			state	query		string					false	"状态参数"
//	@success		200		{object}	GithubCallBackResponse	"登录成功"
//	@failure		400		{object}	common.FailureResponse	"请求参数错误"
//	@failure		500		{object}	common.FailureResponse	"服务器内部错误"
//	@router			/auth/github/callback [get]
//
// GitHubCallback 处理 GitHub OAuth 回调
func (h *AuthHandler) GitHubCallback(ctx context.Context, c *app.RequestContext) {
	// 获取授权码
	code := c.Query("code")
	if code == "" {
		common.RespBadRequest(c, "缺少授权码")
		return
	}

	// 使用授权码换取访问令牌
	accessToken, err := h.exchangeCodeForToken(string(code))
	if err != nil {
		logrus.Errorf("[AUTH] 获取访问令牌失败: %v", err)
		common.RespInternalServerError(c, "获取访问令牌失败")
		return
	}

	// 使用访问令牌获取用户信息
	githubUser, err := h.getUserInfo(accessToken)
	if err != nil {
		logrus.Errorf("[AUTH] 获取用户信息失败: %v", err)
		common.RespInternalServerError(c, "获取用户信息失败")
		return
	}

	// 获取用户邮箱信息 (只使用已验证的邮箱)
	email, err := h.getUserPrimaryEmail(accessToken)
	if err != nil {
		logrus.Errorf("[AUTH] 获取用户邮箱失败: %v", err)
		common.RespInternalServerError(c, "获取用户邮箱失败")
		return
	}

	// 使用显示名称（如果为空则使用用户名）
	nickname := githubUser.Name
	if nickname == "" {
		nickname = githubUser.Login
	}

	// 生成 JWT token
	token, err := jwt.GenerateToken(githubUser.Login, nickname, email)
	if err != nil {
		logrus.Errorf("[AUTH} 生成令牌失败: %v", err)
		common.RespInternalServerError(c, "生成令牌失败")
		return
	}

	// 设置 Cookie (30天过期, HttpOnly)
	c.SetCookie(common.TokenCookieKey, token, 86400*30, "/", "", protocol.CookieSameSiteLaxMode, false, true)

	// 检查是否有重定向地址 (通过 state 参数传递)
	state := c.Query("state")
	if state != "" {
		// 重定向到指定页面
		c.Redirect(302, []byte(string(state)))
		return
	}

	// 返回成功响应 (API 调用或没有重定向地址)
	common.RespSuccess(c, GithubCallBackResponse{
		Username: githubUser.Login,
		Nickname: nickname,
		Email:    email,
	})
}

// exchangeCodeForToken 使用授权码换取访问令牌
func (h *AuthHandler) exchangeCodeForToken(code string) (string, error) {
	data := url.Values{}
	data.Set("client_id", env.GitHubClientID)
	data.Set("client_secret", env.GitHubSecret)
	data.Set("code", code)

	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResp GitHubTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", err
	}

	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("no access token received")
	}

	return tokenResp.AccessToken, nil
}

// getUserInfo 获取用户基本信息
func (h *AuthHandler) getUserInfo(accessToken string) (*GitHubUser, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user GitHubUser
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// getUserPrimaryEmail 获取用户已验证的邮箱地址
func (h *AuthHandler) getUserPrimaryEmail(accessToken string) (string, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var emails []GitHubEmail
	if err := json.Unmarshal(body, &emails); err != nil {
		return "", err
	}

	// 优先查找主要且已验证的邮箱
	for _, email := range emails {
		if email.Primary && email.Verified {
			return email.Email, nil
		}
	}

	// 如果没有主要验证邮箱，返回第一个已验证的邮箱
	for _, email := range emails {
		if email.Verified {
			return email.Email, nil
		}
	}

	// 如果用户没有任何已验证的邮箱，拒绝登录
	return "", fmt.Errorf("用户没有已验证的邮箱地址")
}

type LogoutResponse struct {
	Message string `json:"message"`
}

//	@summary		登出
//	@description	清除登录状态
//	@tags			auth
//	@accept			json
//	@produce		json
//	@success		200	{object}	LogoutResponse			"登出成功"
//	@failure		401	{object}	common.FailureResponse	"未授权"
//	@security		ApiKeyAuth
//	@router			/auth/logout [post]
//
// Logout 登出
func (h *AuthHandler) Logout(ctx context.Context, c *app.RequestContext) {
	// 清除 Cookie
	c.SetCookie(common.TokenCookieKey, "", -1, "/", "", protocol.CookieSameSiteLaxMode, false, true)

	common.RespSuccess(c, LogoutResponse{
		Message: "登出成功",
	})
}

type UserInfoResponse struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

//	@summary		获取当前用户信息
//	@description	获取当前登录用户的信息
//	@tags			auth
//	@accept			json
//	@produce		json
//	@success		200	{object}	UserInfoResponse			"用户信息"
//	@failure		401	{object}	common.FailureResponse	"未授权"
//	@security		ApiKeyAuth
//	@router			/auth/me [get]
//
// GetCurrentUser 获取当前用户信息
func (h *AuthHandler) GetCurrentUser(ctx context.Context, c *app.RequestContext) {
	// 从中间件中获取当前用户信息
	currentUser := common.GetCurrentUser(ctx)
	if currentUser == nil {
		common.RespFailure(c, http.StatusUnauthorized, "未登录")
		return
	}

	common.RespSuccess(c, UserInfoResponse{
		Username: currentUser.Username,
		Nickname: currentUser.Nickname,
		Email:    currentUser.Email,
	})
}
