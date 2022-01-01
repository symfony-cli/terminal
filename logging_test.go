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
	"github.com/rs/zerolog"
	. "gopkg.in/check.v1"
)

type LoggingSuite struct{}

var _ = Suite(&LoggingSuite{})

func (ts *LoggingSuite) TestSetLogLevel(c *C) {
	defer SetLogLevel(1)
	var err error
	c.Assert(Logger.GetLevel(), Equals, zerolog.ErrorLevel)

	err = SetLogLevel(3)
	c.Assert(err, IsNil)
	c.Assert(Logger.GetLevel(), Equals, zerolog.InfoLevel)

	err = SetLogLevel(5)
	c.Assert(err, IsNil)
	c.Assert(Logger.GetLevel(), Equals, zerolog.TraceLevel)

	err = SetLogLevel(1)
	c.Assert(err, IsNil)
	c.Assert(Logger.GetLevel(), Equals, zerolog.ErrorLevel)

	err = SetLogLevel(9)
	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "The provided verbosity level '9' is not in the range [1,4]")
	c.Assert(Logger.GetLevel(), Equals, zerolog.ErrorLevel)
}
