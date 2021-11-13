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
