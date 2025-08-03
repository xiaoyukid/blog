package models

import "fmt"

// 自定义错误类型
type AppError struct {
	Code       int    `json:"code"`    // 业务错误码
	Message    string `json:"message"` // 用户可见的错误信息
	Details    string `json:"details"` // 开发人员可见的错误详情
	StatusCode int    `json:"-"`       // HTTP 状态码
}

// 实现 error 接口
func (e *AppError) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e.StatusCode, e.Message, e.Details)
}

// 添加错误详情
func (e *AppError) WithDetails(details string) *AppError {
	e.Details = details
	return e
}
