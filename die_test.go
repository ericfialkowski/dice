package dice

import "testing"

func TestD4(t *testing.T) {
	test(NewD4(), 4, t)
}

func TestD6(t *testing.T) {
	test(NewD6(), 6, t)
}

func TestD8(t *testing.T) {
	test(NewD8(), 8, t)
}

func TestD10(t *testing.T) {
	test(NewD10(), 10, t)
}

func TestD12(t *testing.T) {
	test(NewD12(), 12, t)
}

func TestD20(t *testing.T) {
	test(NewD20(), 20, t)
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

func testSides(sides int, t *testing.T) {
	test(NewDie(sides), sides, t)
}

func test(d *Die, high int, t *testing.T) {
	for i := 0; i < 1000000; i++ {
		r := d.Roll()
		if r < 1 {
			t.Error(`value too low`)
		}
		if r > high {
			t.Error(`value too high`)
		}
	}
}
