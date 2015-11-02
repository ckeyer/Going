package hello_test

import (
	"io"
	"testing"

	//"fmt"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

type A struct {
	ID   int
	Name string
}

var aa = A{
	ID:   12,
	Name: "ckeyer",
}
var bb = A{
	ID:   12,
	Name: "ckeyer",
}

func (s *MySuite) TestHelloWorld(c *C) {
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
	//	c.Assert(fmt.Sprintf("%#v", aa), Equals, fmt.Sprintf("%#v", bb))
	c.Assert(aa, Equals, &bb)
}
