// Package flash : message 提示信息
package flash

import (
	"github.com/labstack/echo/v4"
)

const (
	flashMessageKeyName = "_flash_message"

	flashMessageSuccess = "success"
	flashMessageInfo    = "info"
	flashMessageWarning = "warning"
	flashMessageDanger  = "danger"
)

type messageFlash struct {
	EchoContext echo.Context
	Data        flashDataValue
}

// NewMessageFlash -
func NewMessageFlash(c echo.Context) *messageFlash {
	return &messageFlash{
		EchoContext: c,
		Data:        make(flashDataValue),
	}
}

func (m *messageFlash) add(typeName string, msg string) {
	if m.Data[typeName] == nil {
		m.Data[typeName] = make([]string, 0)
	}
	m.Data[typeName] = append(m.Data[typeName], msg)
}

func (m *messageFlash) Save() {
	NewFlashData(flashMessageKeyName, m.EchoContext.Request()).
		Set(m.Data).
		Save(m.EchoContext.Response())
}

func (m *messageFlash) Read() flashDataValue {
	return NewFlashData(flashMessageKeyName, m.EchoContext.Request()).
		Read(m.EchoContext.Response())
}

func (m *messageFlash) AddSuccess(msg string) *messageFlash {
	m.add(flashMessageSuccess, msg)
	return m
}

func (m *messageFlash) AddInfo(msg string) *messageFlash {
	m.add(flashMessageInfo, msg)
	return m
}

func (m *messageFlash) AddWarning(msg string) *messageFlash {
	m.add(flashMessageWarning, msg)
	return m
}

func (m *messageFlash) AddDanger(msg string) *messageFlash {
	m.add(flashMessageDanger, msg)
	return m
}

// NewSuccessMessage -
func NewSuccessMessage(c echo.Context, msg string) {
	NewMessageFlash(c).AddSuccess(msg).Save()
}

// NewInfoMessage -
func NewInfoMessage(c echo.Context, msg string) {
	NewMessageFlash(c).AddInfo(msg).Save()
}

// NewWarningMessage -
func NewWarningMessage(c echo.Context, msg string) {
	NewMessageFlash(c).AddWarning(msg).Save()
}

// NewDangerMessage -
func NewDangerMessage(c echo.Context, msg string) {
	NewMessageFlash(c).AddDanger(msg).Save()
}