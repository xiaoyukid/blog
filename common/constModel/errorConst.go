package constModel

import (
	"blog/common/models"
	"net/http"
)

// 错误常量定义
var (
	// 数据库错误
	ErrDBConnection = &models.AppError{
		Code:       1001,
		Message:    "数据库连接失败",
		StatusCode: http.StatusInternalServerError,
	}
	ErrDBNotFound = &models.AppError{
		Code:       1002,
		Message:    "请求的资源不存在",
		StatusCode: http.StatusNotFound,
	}
	ErrDBQuery = &models.AppError{
		Code:       1003,
		Message:    "数据库查询错误",
		StatusCode: http.StatusInternalServerError,
	}

	// 用户认证错误
	ErrAuthRequired = &models.AppError{
		Code:       2001,
		Message:    "需要登录才能访问",
		StatusCode: http.StatusUnauthorized,
	}
	ErrAuthInvalid = &models.AppError{
		Code:       2002,
		Message:    "无效的认证信息",
		StatusCode: http.StatusUnauthorized,
	}
	ErrAuthExpired = &models.AppError{
		Code:       2003,
		Message:    "登录已过期",
		StatusCode: http.StatusUnauthorized,
	}

	// 请求参数错误
	ErrInvalidParam = &models.AppError{
		Code:       3001,
		Message:    "请求参数错误",
		StatusCode: http.StatusBadRequest,
	}
	ErrInvalidJSON = &models.AppError{
		Code:       3002,
		Message:    "无效的JSON格式",
		StatusCode: http.StatusBadRequest,
	}

	// 业务逻辑错误
	ErrArticleNotFound = &models.AppError{
		Code:       4001,
		Message:    "文章不存在",
		StatusCode: http.StatusNotFound,
	}
	ErrCommentNotFound = &models.AppError{
		Code:       4002,
		Message:    "评论不存在",
		StatusCode: http.StatusNotFound,
	}
	ErrPermissionDenied = &models.AppError{
		Code:       4003,
		Message:    "没有操作权限",
		StatusCode: http.StatusForbidden,
	}

	// 系统错误
	ErrInternalServer = &models.AppError{
		Code:       5001,
		Message:    "服务器内部错误",
		StatusCode: http.StatusInternalServerError,
	}
)
