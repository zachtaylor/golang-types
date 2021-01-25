package types_test

import (
	"strings"
	"testing"

	"taylz.io/types"
)

func err() error { return types.NewErr("test") }

func wrapped() types.Error { return types.WrapErr(err(), "wrap") }

func TestTraceErr(t *testing.T) {
	err := types.TraceErr(types.EOF)

	const expectMsg = "EOF"
	if actualMsg := err.Error(); expectMsg != actualMsg {
		t.Log("Expected", expectMsg)
		t.Log("Actual", actualMsg)
		t.Fail()
	}

	const expectSrc = "/taylz.io/types/error_test.go#15"
	if actualSrc := err.Source().String(); !strings.HasSuffix(actualSrc, expectSrc) {
		t.Log("Expected", expectSrc)
		t.Log("Actual", actualSrc)
		t.Fail()
	}
}

func TestErrorSource(t *testing.T) {
	wrapped := wrapped()
	actualOuter := wrapped.Source().String()
	actualInner := types.ErrorSource(wrapped.Unwrap()).String()

	const expectOuter = "/taylz.io/types/error_test.go#12"
	if !strings.HasSuffix(actualOuter, expectOuter) {
		t.Log("Expected", expectOuter)
		t.Log("Actual", actualOuter)
		t.Fail()
	}

	const expectInner = "/taylz.io/types/error_test.go#10"
	if !strings.HasSuffix(actualInner, expectInner) {
		t.Log("Expected", expectInner)
		t.Log("Actual", actualInner)
		t.Fail()
	}
}

func TestErrorTrace(t *testing.T) {
	chain := types.ErrorChain(wrapped())
	if len(chain) != 2 {
		t.Log("Expected len(chain)=2")
		t.Log("Actual", len(chain))
		t.Fail()
	}
	expect0 := "wrap: test"
	if err := chain[0]; err == nil {
		t.Log("Expected error")
		t.Log("Actual", nil)
		t.Fail()
	} else if actual0 := err.Error(); actual0 != expect0 {
		t.Log("Expected", expect0)
		t.Log("Actual", actual0)
		t.Fail()
	}
	expect1 := "test"
	if err := chain[1]; err == nil {
		t.Log("Expected error")
		t.Log("Actual", nil)
		t.Fail()
	} else if actual1 := err.Error(); actual1 != expect1 {
		t.Log("Expected", expect1)
		t.Log("Actual", actual1)
		t.Fail()
	}
}
