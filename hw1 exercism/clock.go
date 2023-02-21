package main

import (
	"fmt"
	"math"
)

// Define the Clock type here.
type Clock struct {
	Hour   int
	Minute int
}

func New(h, m int) Clock {

	return Clock{h, m}.Add(0)
}

func standardization(c Clock) Clock {
	c.Minute = (c.Minute%60 + 60) % 60
	c.Hour = (c.Hour%24 + 24) % 24

	return c
}

func (c Clock) Add(m int) Clock {
	c.Minute += m
	c.Hour += int(math.Floor(float64(c.Minute) / 60))

	return standardization(c)
}

func (c Clock) Subtract(m int) Clock {

	return c.Add(-m)
}

func (c Clock) String() string {

	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

func main() {

	c := New(1, 20)
	fmt.Println(
		c.Add(80),
		c.Add(140),
		c.Subtract(70),
		c.Subtract(80),
		c.Subtract(90),
		c.String(),
		Clock{0, 0}.Add(0),
		-3%2,
		3%2,
	)
}
