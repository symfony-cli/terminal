package terminal

import (
	"fmt"
	"io"
)

type Cursor struct {
	Writer io.Writer
}

func NewCursor(w io.Writer) Cursor {
	return Cursor{Writer: w}
}

func (c Cursor) MoveUp(lines int) Cursor {
	fmt.Fprintf(c.Writer, "\x1b[%dA", lines)
	return c
}

func (c Cursor) MoveDown(lines int) Cursor {
	fmt.Fprintf(c.Writer, "\x1b[%dB", lines)
	return c
}

func (c Cursor) MoveRight(columns int) Cursor {
	fmt.Fprintf(c.Writer, "\x1b[%dC", columns)
	return c
}

func (c Cursor) MoveLeft(columns int) Cursor {
	fmt.Fprintf(c.Writer, "\x1b[%dD", columns)
	return c
}

func (c Cursor) MoveToColumn(column int) Cursor {
	fmt.Fprintf(c.Writer, "\x1b[%dG", column)
	return c
}

func (c Cursor) MoveToPosition(column, row int) Cursor {
	fmt.Fprintf(c.Writer, "\x1b[%d;%dH", row+1, column)
	return c
}

func (c Cursor) SavePosition() Cursor {
	fmt.Fprint(c.Writer, "\x1b7\x1b[s")
	return c
}

func (c Cursor) RestorePosition() Cursor {
	fmt.Fprint(c.Writer, "\x1b[u\x1b8")
	return c
}

func (c Cursor) Hide() Cursor {
	fmt.Fprint(c.Writer, "\x1b[?25l")
	return c
}

func (c Cursor) Show() Cursor {
	fmt.Fprint(c.Writer, "\x1b[?25h\x1b[?0c")
	return c
}

// ClearLine clears all the output from the current line.
func (c Cursor) ClearLine() Cursor {
	fmt.Fprint(c.Writer, "\r\x1b[2K")
	return c
}

// ClearLine clears all the output from the current line after the current position.
func (c Cursor) ClearLineAfter() Cursor {
	fmt.Fprint(c.Writer, "\x1b[K")
	return c
}

// ClearOutput clears all the output from the cursors' current position to the end of the screen.
func (c Cursor) ClearOutput() Cursor {
	fmt.Fprint(c.Writer, "\x1b[0J")
	return c
}

// ClearOutput clears the entire screen.
func (c Cursor) ClearScreen() Cursor {
	fmt.Fprint(c.Writer, "\x1b[0J")
	return c
}
