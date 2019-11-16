package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t tracer) Trace(a ...interface{}) {
	_, _ = t.out.Write([]byte(fmt.Sprint(a...)))
	_, _ = t.out.Write([]byte("\n"))
}
