package terminal

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type OutputFormatterStyleSuite struct{}

var _ = Suite(&OutputFormatterStyleSuite{})

func (style *FormatterStyle) applyStr(str string) string {
	return string(style.apply([]byte(str))) + string(style.unapply())
}

func (ts *OutputFormatterStyleSuite) TestConstructor(c *C) {
	style := NewFormatterStyle("green", "black", []string{"bold", "underscore"})
	c.Check(style.applyStr("foo"), Equals, "\033[32;40;1;4mfoo\033[39;49;22;24m")

	style = NewFormatterStyle("red", "", []string{"blink"})
	c.Check(style.applyStr("foo"), Equals, "\033[31;5mfoo\033[39;25m")

	style = NewFormatterStyle("", "white", nil)
	c.Check(style.applyStr("foo"), Equals, "\033[47mfoo\033[49m")
}

func (ts *OutputFormatterStyleSuite) TestForeground(c *C) {
	style := NewFormatterStyle("black", "", nil)
	c.Check(style.applyStr("foo"), Equals, "\033[30mfoo\033[39m")

	style = NewFormatterStyle("blue", "", nil)
	c.Check(style.applyStr("foo"), Equals, "\033[34mfoo\033[39m")

	style = NewFormatterStyle("default", "", nil)
	c.Check(style.applyStr("foo"), Equals, "\033[39mfoo\033[39m")
}

func (ts *OutputFormatterStyleSuite) TestBackground(c *C) {
	style := NewFormatterStyle("", "black", nil)
	c.Check(style.applyStr("foo"), Equals, "\033[40mfoo\033[49m")

	style = NewFormatterStyle("", "yellow", nil)
	c.Check(style.applyStr("foo"), Equals, "\033[43mfoo\033[49m")

	style = NewFormatterStyle("", "default", nil)
	c.Check(style.applyStr("foo"), Equals, "\033[49mfoo\033[49m")
}

func (ts *OutputFormatterStyleSuite) TestOptions(c *C) {
	style := NewFormatterStyle("", "", []string{"reverse", "conceal"})
	c.Check(style.applyStr("foo"), Equals, "\033[7;8mfoo\033[27;28m")

	style = NewFormatterStyle("", "", []string{"conceal", "reverse"})
	c.Check(style.applyStr("foo"), Equals, "\033[7;8mfoo\033[27;28m")

	style = NewFormatterStyle("", "", []string{"bold"})
	c.Check(style.applyStr("foo"), Equals, "\033[1mfoo\033[22m")

	style = NewFormatterStyle("", "", []string{"bold", "bold"})
	c.Check(style.applyStr("foo"), Equals, "\033[1mfoo\033[22m")
}
