package dice

import "testing"

func TestD4(t *testing.T) {
	d4 := NewD4()
	for i := 0; i < 4000; i++ {
		r := d4.Roll()
		if r < 1 {
			t.Error(`value too low`)
		}
		if r > 4 {
			t.Error(`value too high`)
		}
	}
}

func TestD6(t *testing.T) {
	d6 := NewD6()
	for i := 0; i < 6000; i++ {
		r := d6.Roll()
		if r < 1 {
			t.Error(`value too low`)
		}
		if r > 6 {
			t.Error(`value too high`)
		}
	}
}
