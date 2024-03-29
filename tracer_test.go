package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("New returns nil")
	} else {
		tracer.Trace("Hello, trace package")
		if buf.String() != "Hello, trace package\n" {
			t.Errorf("Output was incorrect string : '%s'", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var slientTracer Tracer = Off()
	slientTracer.Trace("data")
}
