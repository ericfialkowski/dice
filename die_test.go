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
	"testing"
)

func TestD4(t *testing.T) {
	testRoll(NewD4(), 4, t)
}

func TestD6(t *testing.T) {
	testRoll(NewD6(), 6, t)
}

func TestD8(t *testing.T) {
	testRoll(NewD8(), 8, t)
}

func TestD10(t *testing.T) {
	testRoll(NewD10(), 10, t)
}

func TestD12(t *testing.T) {
	testRoll(NewD12(), 12, t)
}

func TestD20(t *testing.T) {
	testRoll(NewD20(), 20, t)
}

func Test5Sided(t *testing.T) {
	testSides(5, t)
}

func Test13Sided(t *testing.T) {
	testSides(13, t)
}

func Test12Sided(t *testing.T) {
	testSides(12, t)
}

func Test2D6(t *testing.T) {
	testRollMulti(NewD6(), 2, 12, t)
}

func Test3D6(t *testing.T) {
	testRollMulti(NewD6(), 3, 18, t)
}

func TestCustomRand(t *testing.T) {
	d1 := NewDieWithRandom(10, *rand.New(rand.NewSource(1)))
	d2 := NewDieWithRandom(10, *rand.New(rand.NewSource(1)))

	for i := 0; i < 1000000; i++ {
		if d1.Roll() != d2.Roll() {
			t.Error(`values don't match`)
			t.FailNow()
		}
	}
}

func testSides(sides int, t *testing.T) {
	testRoll(NewDie(sides), sides, t)
}

func testRoll(d *Die, high int, t *testing.T) {
	for i := 0; i < 1000000; i++ {
		r := d.Roll()
		if r < 1 {
			t.Error(`value too low`)
			t.FailNow()
		}
		if r > high {
			t.Error(`value too high`)
			t.FailNow()
		}
	}
}

func testRollMulti(d *Die, count int, high int, t *testing.T) {
	for i := 0; i < 1000000; i++ {
		r := d.RollMulti(count)
		if r < count {
			t.Error(`value too low`)
			t.FailNow()
		}
		if r > high {
			t.Error(`value too high`)
			t.FailNow()
		}
	}
}
