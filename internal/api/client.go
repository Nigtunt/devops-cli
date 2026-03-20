package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"devops-cli/internal/config"
	"devops-cli/internal/pkg/retry"
)

// Client API 客户端
type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
	debug      bool
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
		debug: cfg.Debug,
	}
}

// SetDebug 设置 debug 模式
func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

// SetToken 设置 token
func (c *Client) SetToken(token string) {
	c.token = token
}

// request 发送 HTTP 请求（带重试）
func (c *Client) request(method, path string, body interface{}) ([]byte, error) {
	return retry.Do(func() ([]byte, error) {
		return c.doRequest(method, path, body)
	}, retry.DefaultConfig, func(attempt int, err error) {
		if c.debug {
			fmt.Fprintf(os.Stderr, "[DEBUG] 重试 #%d: %v\n", attempt+1, err)
		}
	})
}

// doRequest 执行单次 HTTP 请求
func (c *Client) doRequest(method, path string, body interface{}) ([]byte, error) {
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
	req.Header.Set("Authorization", "Bearer "+c.maskToken(c.token))
	req.Header.Set("User-Agent", "devops-cli/1.0")

	// Debug 模式打印请求详情
	if c.debug {
		fmt.Fprintf(os.Stderr, "\n===== [DEBUG] Request =====\n")
		fmt.Fprintf(os.Stderr, "%s %s\n", method, c.baseURL+path)
		fmt.Fprintf(os.Stderr, "Headers:\n")
		for key, values := range req.Header {
			for _, value := range values {
				// 脱敏 Authorization
				if key == "Authorization" {
					value = "Bearer " + c.maskToken(c.token)
				}
				fmt.Fprintf(os.Stderr, "  %s: %s\n", key, value)
			}
		}
		if body != nil {
			bodyJSON, _ := json.MarshalIndent(body, "", "  ")
			fmt.Fprintf(os.Stderr, "Body:\n%s\n", string(bodyJSON))
		}
		fmt.Fprintf(os.Stderr, "==========================\n\n")
	}

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

	// Debug 模式打印响应详情
	if c.debug {
		fmt.Fprintf(os.Stderr, "===== [DEBUG] Response =====\n")
		fmt.Fprintf(os.Stderr, "Status: %d %s\n", resp.StatusCode, resp.Status)
		fmt.Fprintf(os.Stderr, "Headers:\n")
		for key, values := range resp.Header {
			for _, value := range values {
				fmt.Fprintf(os.Stderr, "  %s: %s\n", key, value)
			}
		}
		if len(respBody) > 0 {
			fmt.Fprintf(os.Stderr, "Body:\n%s\n", c.maskSensitiveData(string(respBody)))
		}
		fmt.Fprintf(os.Stderr, "===========================\n\n")
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API 错误 (状态码 %d): %s", resp.StatusCode, c.maskSensitiveData(string(respBody)))
	}

	return respBody, nil
}

// maskToken 脱敏 token
func (c *Client) maskToken(token string) string {
	if len(token) <= 8 {
		return "****"
	}
	return token[:4] + "****" + token[len(token)-4:]
}

// maskSensitiveData 脱敏敏感数据
func (c *Client) maskSensitiveData(data string) string {
	// 脱敏 token 字段
	data = strings.ReplaceAll(data, c.token, "****")
	if len(c.token) > 8 {
		data = strings.ReplaceAll(data, c.token[:4]+"****"+c.token[len(c.token)-4:], "****")
	}
	return data
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
