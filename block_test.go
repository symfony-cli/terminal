package terminal

import (
	. "gopkg.in/check.v1"
)

type OutputBlockSuite struct{}

var _ = Suite(&OutputBlockSuite{})

func (ts *OutputBlockSuite) TestSplitsBlockLines(c *C) {
	lines, maxLen := splitsBlockLines("Foo Bazz", 3)
	c.Assert(lines, DeepEquals, []string{"Foo", " Ba", "zz"})
	c.Assert(maxLen, Equals, 3)

	lines, maxLen = splitsBlockLines("<href=https://example.com>Foo</>", 3)
	c.Assert(lines, DeepEquals, []string{"<href=https://example.com>Foo</>"})
	c.Assert(maxLen, Equals, 3)

	lines, maxLen = splitsBlockLines("<href=https://example.com>Foo</>Bar", 3)
	c.Assert(lines, DeepEquals, []string{"<href=https://example.com>Foo</>", "Bar"})
	c.Assert(maxLen, Equals, 3)

	lines, maxLen = splitsBlockLines("<href=https://example.com>Foo Bar</>Baz", 3)
	c.Assert(lines, DeepEquals, []string{"<href=https://example.com>Foo Bar</>", "Baz"})
	c.Assert(maxLen, Equals, 3)
}

func (ts *OutputBlockSuite) TestSplitsBlockLinesDonotPanic(c *C) {
	lines, maxLen := splitsBlockLines("Foo Baz<", 4)
	c.Assert(lines, DeepEquals, []string{"Foo ", "Baz<"})
	c.Assert(maxLen, Equals, 4)
}
