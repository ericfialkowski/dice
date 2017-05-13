package dice

/*
MIT License

Copyright (c) 2017 Eric Fialkowski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

import (
	"math/rand"
	"time"
)

// Die represents an n sided Die
type Die struct {
	sides int
	r     *rand.Rand
}

// NewDie returns a new Die of sides
func NewDie(sides int) *Die {
	return &Die{sides, rand.New(rand.NewSource(time.Now().UnixNano()))}
}

// NewDieWithRandom returns a new Die of sides
func NewDieWithRandom(sides int, r rand.Rand) *Die {
	return &Die{sides, &r}
}

// NewD4 returns a new 4 sided Die
func NewD4() *Die {
	return NewDie(4)
}

// NewD6 returns a new 6 sided Die
func NewD6() *Die {
	return NewDie(6)
}

// NewD8 returns a new 4 sided Die
func NewD8() *Die {
	return NewDie(8)
}

// NewD10 returns a new 6 sided Die
func NewD10() *Die {
	return NewDie(10)
}

// NewD12 returns a new 12 sided Die
func NewD12() *Die {
	return NewDie(12)
}

// NewD20 returns a new 6 sided Die
func NewD20() *Die {
	return NewDie(20)
}

// Roll returns a random Die roll
func (d *Die) Roll() int {
	return d.r.Intn(d.sides) + 1
}

// RollMulti returns the sum of a count of random Die rolls
func (d *Die) RollMulti(count int) int {
	var rtn int
	for i := 0; i < count; i++ {
		rtn += d.Roll()
	}
	return rtn
}
