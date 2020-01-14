# go-enum


```go
package enum

import "testing"

type Human struct {
	gender int    `enum:"[1,2,3]"`
	a      string `enum:"[li,qq]"`
}

func TestVerify(t *testing.T) {
	h := Human{
		4,
		"li",
	}
	v := New(h)
	if err := v.Verify(); err != nil {
		t.Log(err)
	} else {
		t.Fail()
	}
}

```