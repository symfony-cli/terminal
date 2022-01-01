/*
 * Copyright (c) 2021-present Fabien Potencier <fabien@symfony.com>
 *
 * This file is part of Symfony CLI project
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

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
