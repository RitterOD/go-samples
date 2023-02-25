package interface_exercises

import (
	"fmt"
	"io"
)

type CountedWriter struct {
	counter *int
	writer  io.Writer
}

func (c *CountedWriter) Write(p []byte) (int, error) {
	rv, err := c.writer.Write(p)
	*c.counter += rv
	fmt.Printf("\ncounter %d\n", *c.counter)
	return rv, err
}

func (c *CountedWriter) setWriter(writer io.Writer) *CountedWriter {
	c.writer = writer
	return c
}

func Wrap(writer io.Writer, counter *int) io.Writer {
	var rv = CountedWriter{writer: writer, counter: counter}
	return &rv
}
