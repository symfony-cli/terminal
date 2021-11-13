package terminal

import (
	. "gopkg.in/check.v1"
)

type ColorSuite struct{}

var _ = Suite(&ColorSuite{})

func (ts *ColorSuite) TestForeground(c *C) {
	_, err := NewColor("undefined-color", "", nil)
	c.Check(err, ErrorMatches, "invalid \"undefined-color\" color")
}

func (ts *ColorSuite) TestBackground(c *C) {
	_, err := NewColor("", "undefined-color", nil)
	c.Check(err, ErrorMatches, "invalid \"undefined-color\" color")
}

func (ts *ColorSuite) TestOptions(c *C) {
	_, err := NewColor("", "", []string{"foo"})
	c.Check(err, ErrorMatches, "invalid option specified: foo.")
}
