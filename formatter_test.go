package terminal

import (
	"bytes"
	"strings"

	. "gopkg.in/check.v1"
)

type StyleStackSuite struct{}

var _ = Suite(&StyleStackSuite{})

func (ts *StyleStackSuite) Testpush(c *C) {
	stack := styleStack{}
	s1 := &FormatterStyle{}
	s2 := &FormatterStyle{}
	stack.push(s1)
	stack.push(s2)

	c.Assert(s2, Equals, stack.current())

	s3 := &FormatterStyle{}
	stack.push(s3)
	c.Assert(s3, Equals, stack.current())
}

func (ts *StyleStackSuite) Testpop(c *C) {
	stack := styleStack{}
	s1 := NewFormatterStyle("white", "black", nil)
	s2 := NewFormatterStyle("yellow", "blue", nil)
	stack.push(s1)
	stack.push(s2)

	c.Assert(s2, Equals, stack.pop(nil))
	c.Assert(s1, Equals, stack.pop(nil))
}

func (ts *StyleStackSuite) TestPopCurrentEmpty(c *C) {
	stack := styleStack{}
	style := NewFormatterStyle("", "", nil)
	poppedStyle := stack.pop(nil)

	c.Assert(poppedStyle, NotNil)
	c.Assert(poppedStyle, DeepEquals, style)

	currentStyle := stack.current()

	c.Assert(currentStyle, NotNil)
	c.Assert(currentStyle, DeepEquals, style)
}

func (ts *StyleStackSuite) TestPopNotLast(c *C) {
	stack := styleStack{}
	s1 := NewFormatterStyle("white", "black", nil)
	s2 := NewFormatterStyle("yellow", "blue", nil)
	s3 := NewFormatterStyle("green", "red", nil)
	stack.push(s1)
	stack.push(s2)
	stack.push(s3)

	c.Assert(s2, Equals, stack.pop(s2))
	c.Assert(s1, Equals, stack.pop(nil))
}

func (ts *StyleStackSuite) TestInvalidpop(c *C) {
	defer func() {
		err := recover()
		c.Assert(err, Equals, "Incorrectly nested style tag found.")
	}()
	stack := styleStack{}
	s1 := NewFormatterStyle("white", "black", nil)
	s2 := NewFormatterStyle("yellow", "blue", nil)

	stack.push(s1)
	stack.pop(s2)

	// If we reach this point, test has failed
	c.FailNow()
}

type OutputFormatterSuite struct{}

var _ = Suite(&OutputFormatterSuite{})

func (formatter *Formatter) formatString(msg string) string {
	buf := bytes.NewBuffer([]byte(``))
	formatter.Format([]byte(msg), buf)

	return buf.String()
}

func (ts *OutputFormatterSuite) TestEmptyTag(c *C) {
	formatter := NewFormatter()
	c.Check(formatter.formatString("foo<>bar"), Equals, "foo<>bar")
}

func (ts *OutputFormatterSuite) TestWrittenBytesCount(c *C) {
	formatter := NewFormatter()
	buf := bytes.NewBuffer([]byte(``))

	n, _ := formatter.Format([]byte("foo<>bar"), buf)
	c.Check(n, Equals, 8)

	buf.Reset()
	n, _ = formatter.Format([]byte("<info>some info</info>"), buf)
	c.Check(n, Equals, 9)

	buf.Reset()
	n, _ = formatter.Format([]byte("<info></info>"), buf)
	c.Check(n, Equals, 0)

	buf.Reset()
	n, _ = formatter.Format([]byte("<a>some info</a>"), buf)
	c.Check(n, Equals, 16)

	buf.Reset()
	n, _ = formatter.Format([]byte("<error>error<info>info<comment>comment</info>error</error>"), buf)
	c.Check(n, Equals, 21)
}

func (ts *OutputFormatterSuite) TestLGCharEscaping(c *C) {
	formatter := NewFormatter()

	c.Check(formatter.formatString("foo\\<bar"), Equals, "foo<bar")
	c.Check(formatter.formatString("\\<info>some info\\</info>"), Equals, "<info>some info</info>")
	c.Check(string(Escape([]byte("<info>some info</info>"))), Equals, "\\<info>some info\\</info>")

	c.Check(
		formatter.formatString("<comment>github.com/symfonycorp/symfony-cli/console does work very well!</comment>"),
		Equals,
		"\033[33mgithub.com/symfonycorp/symfony-cli/console does work very well!\033[39m",
	)
}

func (ts *OutputFormatterSuite) TestBundledStyles(c *C) {
	formatter := NewFormatter()

	c.Check(formatter.HasStyle("error"), Equals, true)
	c.Check(formatter.HasStyle("info"), Equals, true)
	c.Check(formatter.HasStyle("comment"), Equals, true)
	c.Check(formatter.HasStyle("question"), Equals, true)

	c.Check(
		formatter.formatString("<error>some error</error>"),
		Equals,
		"\033[37;41msome error\033[39;49m",
	)
	c.Check(
		formatter.formatString("<info>some info</info>"),
		Equals,
		"\033[32msome info\033[39m",
	)
	c.Check(
		formatter.formatString("<comment>some comment</comment>"),
		Equals,
		"\033[33msome comment\033[39m",
	)
	c.Check(
		formatter.formatString("<question>some question</question>"),
		Equals,
		"\033[30;46msome question\033[39;49m",
	)
}

func (ts *OutputFormatterSuite) TestNestedStyles(c *C) {
	formatter := NewFormatter()
	c.Check(
		formatter.formatString("<error>some <info>some info</info> error</error>"),
		Equals,
		"\033[37;41msome \033[32msome info\033[39m\033[37;41m error\033[39;49m",
	)
}

func (ts *OutputFormatterSuite) TestAdjacentStyles(c *C) {
	formatter := NewFormatter()
	c.Check(
		formatter.formatString("<error>some error</error><info>some info</info>"),
		Equals,
		"\033[37;41msome error\033[39;49m\033[32msome info\033[39m",
	)
}

func (ts *OutputFormatterSuite) TestStyleMatchingNotGreedy(c *C) {
	formatter := NewFormatter()
	c.Check(
		formatter.formatString("(<info>>=2.0,<2.3</info>)"),
		Equals,
		"(\033[32m>=2.0,<2.3\033[39m)",
	)
}

func (ts *OutputFormatterSuite) TestStyleEscaping(c *C) {
	formatter := NewFormatter()
	c.Check(
		formatter.formatString("(<info>"+string(Escape([]byte("z>=2.0,<\\<<a2.3\\")))+"</info>)"),
		Equals,
		"(\033[32mz>=2.0,<<<a2.3\\\033[39m)",
	)
	c.Check(
		formatter.formatString("<info>"+string(Escape([]byte("<error>some error</error>")))+"</info>"),
		Equals,
		"\033[32m<error>some error</error>\033[39m",
	)
}

func (ts *OutputFormatterSuite) TestDeepNestedStyles(c *C) {
	formatter := NewFormatter()
	c.Check(
		formatter.formatString("<error>error<info>info<comment>comment</info>error</error>"),
		Equals,
		"\033[37;41merror\033[32minfo\033[33mcomment\033[39m\033[37;41merror\033[39;49m",
	)
}

func (ts *OutputFormatterSuite) TestRedefineStyle(c *C) {
	formatter := NewFormatter()
	style := NewFormatterStyle("blue", "white", nil)
	formatter.SetStyle("info", style)

	c.Check(
		formatter.formatString("<info>some custom msg</info>"),
		Equals,
		"\033[34;47msome custom msg\033[39;49m",
	)
}

func (ts *OutputFormatterSuite) TestInlineStyle(c *C) {
	formatter := NewFormatter()

	c.Check(
		formatter.formatString("<fg=blue;bg=red>some text</>"),
		Equals,
		"\033[34;41msome text\033[39;49m",
	)

	c.Check(
		formatter.formatString("<fg=blue;bg=red>some text</fg=blue;bg=red>"),
		Equals,
		"\033[34;41msome text\033[39;49m",
	)
}

func (ts *OutputFormatterSuite) TestInlineStyleOptions(c *C) {
	formatter := NewFormatter()

	c.Check(formatter.formatString("<fg=green;>[test]</>"), Equals, "\033[32m[test]\033[39m")
	c.Check(formatter.formatString("<fg=green;bg=blue;>a</>"), Equals, "\033[32;44ma\033[39;49m")
	c.Check(formatter.formatString("<fg=green;options=bold>b</>"), Equals, "\033[32;1mb\033[39;22m")
	c.Check(formatter.formatString("<fg=green;options=reverse;>a</>"), Equals, "\033[32;7ma\033[39;27m")
	c.Check(formatter.formatString("<fg=green;options=reverse;><a></>"), Equals, "\033[32;7m<a>\033[39;27m")
	c.Check(formatter.formatString("<fg=green;options=bold,underscore>z</>"), Equals, "\033[32;1;4mz\033[39;22;24m")
	c.Check(formatter.formatString("<fg=green;options=bold,underscore,reverse;>d</>"), Equals, "\033[32;1;4;7md\033[39;22;24;27m")
	c.Check(formatter.formatString("<href=idea://open/?file=/path/somefile.php&line=12>some URL</>"), Equals, "\033]8;;idea://open/?file=/path/somefile.php&line=12\033\\some URL\033]8;;\033\\")
}

func (ts *OutputFormatterSuite) TestNonStyleTag(c *C) {
	formatter := NewFormatter()

	c.Check(
		formatter.formatString("<info>some <tag> <setting=value> styled <p>single-char tag</p></info>"),
		Equals,
		"\033[32msome \033[32m<tag>\033[32m \033[32m<setting=value>\033[32m styled \033[32m<p>\033[32msingle-char tag</p>\033[39m",
	)
}

func (ts *OutputFormatterSuite) TestFormatLongString(c *C) {
	formatter := NewFormatter()
	long := strings.Repeat("\\", 14000)

	c.Check(
		formatter.formatString("<error>some error</error>"+long),
		Equals,
		"\033[37;41msome error\033[39;49m"+long,
	)
}

func (ts *OutputFormatterSuite) TestNotDecoratedFormatter(c *C) {
	formatter := NewFormatter()
	formatter.Decorated = false
	formatter.SupportsAdvancedDecoration = false

	c.Check(formatter.HasStyle("error"), Equals, true)
	c.Check(formatter.HasStyle("info"), Equals, true)
	c.Check(formatter.HasStyle("comment"), Equals, true)
	c.Check(formatter.HasStyle("question"), Equals, true)

	c.Check(
		formatter.formatString("<error>some error</error>"),
		Equals,
		"some error",
	)
	c.Check(
		formatter.formatString("<info>some info</info>"),
		Equals,
		"some info",
	)
	c.Check(
		formatter.formatString("<comment>some comment</comment>"),
		Equals,
		"some comment",
	)
	c.Check(
		formatter.formatString("<question>some question</question>"),
		Equals,
		"some question",
	)
	c.Check(
		formatter.formatString("<href=idea://open/?file=/path/somefile.php&line=12>some URL</>"),
		Equals,
		"some URL",
	)

	formatter.Decorated = true

	c.Check(
		formatter.formatString("<error>some error</error>"),
		Equals,
		"\033[37;41msome error\033[39;49m",
	)
	c.Check(
		formatter.formatString("<info>some info</info>"),
		Equals,
		"\033[32msome info\033[39m",
	)
	c.Check(
		formatter.formatString("<comment>some comment</comment>"),
		Equals,
		"\033[33msome comment\033[39m",
	)
	c.Check(
		formatter.formatString("<question>some question</question>"),
		Equals,
		"\033[30;46msome question\033[39;49m",
	)

	c.Check(
		formatter.formatString("<href=idea://open/?file=/path/somefile.php&line=12>some URL</>"),
		Equals,
		"some URL",
	)
	formatter.SupportsAdvancedDecoration = true
	c.Check(
		formatter.formatString("<href=idea://open/?file=/path/somefile.php&line=12>some URL</>"),
		Not(Equals),
		"some URL",
	)
}

func (ts *OutputFormatterSuite) TestFormatterHrefOptions(c *C) {
	formatter := NewFormatter()
	formatter.Decorated = true
	formatter.SupportsAdvancedDecoration = true

	c.Check(
		formatter.formatString("<href=idea://open/?file=/path/somefile.php&line=12>some URL</>"),
		Not(Equals),
		"some URL",
	)

	c.Check(
		formatter.formatString("<href=https://foo.bar/XXyy>some URL</>"),
		Equals,
		"\033]8;;https://foo.bar/XXyy\033\\some URL\033]8;;\033\\",
	)

	formatter.SupportsAdvancedDecoration = false
	c.Check(
		formatter.formatString("<href=idea://open/?file=/path/somefile.php&line=12>some URL</>"),
		Equals,
		"some URL",
	)

	formatter.Decorated = false
	formatter.SupportsAdvancedDecoration = true
	c.Check(
		formatter.formatString("<href=idea://open/?file=/path/somefile.php&line=12>some URL</>"),
		Equals,
		"some URL",
	)

	formatter.Decorated = false
	formatter.SupportsAdvancedDecoration = false
	c.Check(
		formatter.formatString("<href=idea://open/?file=/path/somefile.php&line=12>some URL</>"),
		Equals,
		"some URL",
	)
}
