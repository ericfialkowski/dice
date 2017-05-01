package die

import (
	"math/rand"
	"time"
)

/*
Die represents an n sided Die
*/
type Die struct {
	sides int
	r     *rand.Rand
}

/*
NewDie returns a new Die of sides
*/
func NewDie(sides int) *Die {
	return &Die{sides, rand.New(rand.NewSource(time.Now().UnixNano()))}
}

/*
NewDieWithRandom returns a new Die of sides
*/
func NewDieWithRandom(sides int, r rand.Rand) *Die {
	return &Die{sides, &r}
}

/*
NewD4 returns a new 4 sided Die
*/
func NewD4() *Die {
	return NewDie(4)
}

/*
NewD6 returns a new 6 sided Die
*/
func NewD6() *Die {
	return NewDie(6)
}

/*
NewD8 returns a new 4 sided Die
*/
func NewD8() *Die {
	return NewDie(8)
}

/*
NewD10 returns a new 6 sided Die
*/
func NewD10() *Die {
	return NewDie(10)
}

/*
NewD12 returns a new 12 sided Die
*/
func NewD12() *Die {
	return NewDie(12)
}

/*
NewD20 returns a new 6 sided Die
*/
func NewD20() *Die {
	return NewDie(20)
}

/*
Roll returns a random Die roll
*/
func (d *Die) Roll() int {
	return (d.r.Int() % d.sides) + 1
}

/*
RollMulti returns the sum of a count of random Die rolls
*/
func (d *Die) RollMulti(count int) int {
	var rtn int
	for i := 0; i < count; i++ {
		rtn += d.Roll()
	}
	return rtn
}

// func main() {
// 	d := NewD6()
// 	for i := 0; i < 1000; i++ {
// 		fmt.Printf("%d\t%d\n", i, d.RollMulti(2))
// 	}
// }
