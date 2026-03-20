package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"devops-cli/internal/config"
)

// Client API 客户端
type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// Response 通用 API 响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// UserInfo 用户信息
type UserInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// NewClient 创建 API 客户端
func NewClient() *Client {
	cfg := config.Get()
	return &Client{
		baseURL: cfg.BaseURL,
		token:   cfg.APIToken,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SetToken 设置 token
func (c *Client) SetToken(token string) {
	c.token = token
}

// request 发送 HTTP 请求
func (c *Client) request(method, path string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("序列化请求失败：%w", err)
		}
		reqBody = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, c.baseURL+path, reqBody)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败：%w", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("User-Agent", "devops-cli/1.0")

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败：%w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API 错误 (状态码 %d): %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// ValidateToken 验证 token
func (c *Client) ValidateToken() (*UserInfo, error) {
	respBody, err := c.request("GET", "/api/v1/auth/validate", nil)
	if err != nil {
		return nil, err
	}

	var resp Response
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("token 验证失败：%s", resp.Message)
	}

	userInfo, ok := resp.Data.(*UserInfo)
	if !ok {
		// 手动解析嵌套结构
		var rawResp struct {
			Code    int             `json:"code"`
			Message string          `json:"message"`
			Data    json.RawMessage `json:"data"`
		}
		if err := json.Unmarshal(respBody, &rawResp); err != nil {
			return nil, fmt.Errorf("解析响应失败：%w", err)
		}
		
		var user UserInfo
		if err := json.Unmarshal(rawResp.Data, &user); err != nil {
			return nil, fmt.Errorf("解析用户信息失败：%w", err)
		}
		return &user, nil
	}

	return userInfo, nil
}

// GetUser 获取当前用户信息
func (c *Client) GetUser() (*UserInfo, error) {
	respBody, err := c.request("GET", "/api/v1/user", nil)
	if err != nil {
		return nil, err
	}

	var resp Response
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	var user UserInfo
	data, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, fmt.Errorf("解析用户信息失败：%w", err)
	}

	return &user, nil
}
