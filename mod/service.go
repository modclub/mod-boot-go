package mod

import (
	"bytes"
	"strings"
	"text/template"
)

type Service interface {
	Name() string
	LocaleMessages() map[string]string
	Handler(*Context) *STDReply
}

type Context struct {
	LocaleMessages map[string]string
}

type STDReply struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"msg,omitempty"`
}

func (c Context) STD(data interface{}) *STDReply {
	return c.STDWithMessage(data, "")
}

func (c Context) STDWithMessage(data interface{}, messageId string, messageValues ...D) *STDReply {
	return c.std(0, data, messageId, messageValues...)
}

func (c Context) STDOK() *STDReply {
	return c.STD("ok")
}

func (c Context) STDErr(data interface{}) *STDReply {
	return c.STDWithMessage(data, "")
}

func (c Context) std(code int, data interface{}, messageId string, messageValues ...D) *STDReply {
	messageId = strings.TrimSpace(messageId)
	reply := &STDReply{
		Code: code,
		Data: data,
	}
	if messageId != "" && len(c.LocaleMessages) > 0 {
		if messageContent := c.LocaleMessages[messageId]; messageContent != "" {
			var (
				v      = template.Must(template.New("").Parse(messageContent))
				b      bytes.Buffer
				values D
			)
			if len(values) > 0 {
				values = messageValues[0]
			}
			if err := v.Execute(&b, values); err == nil {
				reply.Message = b.String()
			}
		}
	}
	return reply
}

func (c Context) FormatMessage(messageId string, values D) string {
	messageId = strings.TrimSpace(messageId)
	if messageId != "" && len(c.LocaleMessages) > 0 {
		messageContent := c.LocaleMessages[messageId]
		if messageContent != "" {
			v := template.Must(template.New("").Parse(messageContent))
			var b bytes.Buffer
			if err := v.Execute(&b, values); err == nil {
				return b.String()
			}
		}
	}
	return ""
}
