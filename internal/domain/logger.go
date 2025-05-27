package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogLevel string

const (
	LogLevelInfo    LogLevel = "info"
	LogLevelError   LogLevel = "error"
	LogLevelWarn LogLevel = "warn"
	LogLevelDebug   LogLevel = "debug"
)

type MethodType string

const (
	MethodTypeGet    MethodType = "GET"
	MethodTypePost   MethodType = "POST"
	MethodTypePut    MethodType = "PUT"
	MethodTypeDelete MethodType = "DELETE"
	MethodTypePatch  MethodType = "PATCH"
	MethodTypeOptions MethodType = "OPTIONS"
)

type Logging struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Method    MethodType        `gorm:"type:varchar(10)" json:"method"`
	Level     LogLevel       `gorm:"type:varchar(20)" json:"level"`
	Message   string         `gorm:"type:text" json:"message"`
	Source    string         `gorm:"type:varchar(100)" json:"source"`
	UserID    *uuid.UUID        `gorm:"type:char(36);index;null" json:"user_id,omitempty"`
	UserIP    *string        `gorm:"type:varchar(45);null" json:"user_ip,omitempty"`   
	UserAgent  *string        `gorm:"type:text;null" json:"user_agent,omitempty"` 
	Metadata  *string        `gorm:"type:json;null" json:"metadata,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type LoggerRepository interface {
	CreateLog(log *Logging) error
}