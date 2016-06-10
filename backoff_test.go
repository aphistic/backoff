package backoff

import (
	"time"

	. "gopkg.in/check.v1"
)

func (s *BackoffSuite) TestZeroBackOff(c *C) {
	b := NewZeroBackOff()
	testSequence(c, b, time.Millisecond, []uint{0, 0, 0, 0})
	b.Reset()
	testSequence(c, b, time.Millisecond, []uint{0, 0, 0, 0})
}

func (s *BackoffSuite) TestConstantBackOff(c *C) {
	b1 := NewConstantBackOff(25 * time.Second)
	b2 := NewConstantBackOff(50 * time.Minute)

	testSequence(c, b1, time.Second, []uint{25, 25, 25, 25})
	b2.Reset()
	testSequence(c, b2, time.Minute, []uint{50, 50, 50, 50})

	testSequence(c, b1, time.Second, []uint{25, 25, 25, 25})
	b1.Reset()
	testSequence(c, b2, time.Minute, []uint{50, 50, 50, 50})
}

func (s *BackoffSuite) TestLinearBackOffMax(c *C) {
	b := NewLinearBackOff(time.Millisecond, time.Millisecond, time.Millisecond*4)

	testSequence(c, b, time.Millisecond, []uint{1, 2, 3, 4, 4, 4})
	b.Reset()
	testSequence(c, b, time.Millisecond, []uint{1, 2, 3, 4, 4, 4})
}