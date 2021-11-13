package terminal

import (
	"bytes"
	"os"
	"sync"
)

var (
	hrefSupport     bool
	hrefSupportOnce sync.Once
)

type FormatterStyle struct {
	color *Color
	href  string
}

func NewFormatterStyle(foreground, background string, options []string) *FormatterStyle {
	color, err := NewColor(foreground, background, options)
	if err != nil {
		panic(err)
	}
	return &FormatterStyle{color: color}
}

func (style *FormatterStyle) GetHref() string {
	return style.href
}

func (style *FormatterStyle) SetHref(href string) {
	style.href = href
}

func (style *FormatterStyle) apply(msg []byte) []byte {
	buf := bytes.NewBuffer([]byte{})
	if hasHrefSupport() && len(style.href) > 0 {
		buf.WriteString("\033]8;;")
		buf.WriteString(style.href)
		buf.WriteString("\033\\")
	}
	buf.Write(style.color.Set())
	buf.Write(msg)
	return buf.Bytes()
}

func (style *FormatterStyle) unapply() []byte {
	buf := bytes.NewBuffer([]byte{})
	buf.Write(style.color.Unset())
	if hasHrefSupport() && len(style.href) > 0 {
		buf.WriteString("\033]8;;\033\\")
	}
	return buf.Bytes()
}

func hasHrefSupport() bool {
	hrefSupportOnce.Do(func() {
		hrefSupport = os.Getenv("TERMINAL_EMULATOR") != "JetBrains-JediTerm" && os.Getenv("KONSOLE_VERSION") == ""
	})
	return hrefSupport
}
