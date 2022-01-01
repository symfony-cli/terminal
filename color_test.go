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
