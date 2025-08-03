package models

import "time"

// 定义一个全局的错误记录结构体
type ErrorLog struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Time       time.Time `json:"time" gorm:"column:log_time"`
	Code       int       `json:"code"` // 业务错误码
	StatusCode int       `json:"status" gorm:"column:status_code"`
	Message    string    `json:"message" gorm:"column:message"`
	Path       string    `json:"path" gorm:"column:request_path"`
	Method     string    `json:"method" gorm:"column:http_method"`
	Details    string    `json:"error" gorm:"column:error_details"`
	IP         string    `json:"ip" gorm:"column:client_ip"`
	UserAgent  string    `gorm:"column:user_agent"`
	UserID     int       `json:"user_id" gorm:"column:user_id"`
}
