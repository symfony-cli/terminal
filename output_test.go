package terminal

import (
	"bytes"

	"github.com/pkg/errors"
	. "gopkg.in/check.v1"
)

type ConsoleOutputSuite struct{}

var _ = Suite(&ConsoleOutputSuite{})

type fakeWriteCloser struct {
	*bytes.Buffer

	hasBeenClosed bool
}

func (wc *fakeWriteCloser) Close() error {
	wc.hasBeenClosed = true

	return errors.New("called")
}

func (ts *ConsoleOutputSuite) TestConsoleOutput(c *C) {
	buffer := new(bytes.Buffer)
	output := NewBufferedConsoleOutput(buffer, buffer)

	output.Write([]byte("test"))
	c.Assert(buffer.String(), Equals, "test")

	formatter := NewFormatter()
	c.Assert(output.GetFormatter(), Not(Equals), formatter)

	output.SetFormatter(formatter)
	c.Assert(output.GetFormatter(), Equals, formatter)
}

func (ts *ConsoleOutputSuite) TestClose(c *C) {
	buffer := new(bytes.Buffer)
	wc := new(fakeWriteCloser)
	output := NewBufferedConsoleOutput(wc, buffer)

	err := output.Stderr.Close()
	c.Assert(err, IsNil)

	c.Assert(wc.hasBeenClosed, Equals, false)
	err = output.Stdout.Close()
	c.Assert(wc.hasBeenClosed, Equals, true)
	c.Assert(err, ErrorMatches, "called")
}

func (ts *ConsoleOutputSuite) TestWrappers(c *C) {
	previousStdout, previousStderr := Stdout, Stderr
	defer func() {
		Stdout, Stderr = previousStdout, previousStderr
	}()

	bufferStdout := new(bytes.Buffer)
	bufferStderr := new(bytes.Buffer)
	Stdout = NewBufferedConsoleOutput(bufferStdout, bufferStderr)
	Stderr = Stdout.Stderr

	bufferStdout.Reset()
	Print("test")
	c.Check(bufferStdout.String(), Equals, "test")

	bufferStdout.Reset()
	Println("test")
	c.Check(bufferStdout.String(), Equals, "test\n")

	bufferStdout.Reset()
	Printf("test %d", 2)
	c.Check(bufferStdout.String(), Equals, "test 2")

	bufferStdout.Reset()
	Printfln("test %d", 3)
	c.Check(bufferStdout.String(), Equals, "test 3\n")

	bufferStderr.Reset()
	Eprint("test")
	c.Check(bufferStderr.String(), Equals, "test")

	bufferStderr.Reset()
	Eprintln("test")
	c.Check(bufferStderr.String(), Equals, "test\n")

	bufferStderr.Reset()
	Eprintf("test %d", 2)
	c.Check(bufferStderr.String(), Equals, "test 2")

	bufferStderr.Reset()
	Eprintfln("test %d", 3)
	c.Check(bufferStderr.String(), Equals, "test 3\n")
}
