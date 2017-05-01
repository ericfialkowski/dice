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
	"testing"
)

func TestNewBag(t *testing.T) {
	b := NewBag()
	if b.bag == nil {
		t.Error(`Internal bag not instantiated`)
	}
}

func TestBagNew(t *testing.T) {
	b := new(Bag)
	if b.bag != nil {
		t.Error(`Internal bag somehow instantiated`)
	}
}

func TestGet(t *testing.T) {
	b := NewBag()
	d1 := b.Get(6)
	d2 := b.Get(6)
	d3 := b.Get(4)
	if d1 != d2 {
		t.Error("Dice aren't same")
	}
	if d1 == d3 {
		t.Error("Dice shouldn't be the same")
	}
}
