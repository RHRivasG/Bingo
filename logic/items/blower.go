package items

import (
	"math/rand"
)

//Blower .
type Blower struct {
	Balls []Ball
}

//NewBlower .
func NewBlower() Blower {
	vingu := []string{"V", "I", "N", "G", "U"}
	var balls []Ball
	for i, letter := range vingu {
		for j := i*15 + 1; j <= i*15+15; j++ {
			balls = append(balls, Ball{letter, j})
		}
	}
	return Blower{balls}
}

//GetBallOut .
func (b *Blower) GetBallOut() Ball {
	n := rand.Intn(len(b.Balls))
	ball := b.Balls[n]
	b.Balls = remove(b.Balls, n)
	return ball
}

func remove(balls []Ball, i int) []Ball {
	return append(balls[:i], balls[i+1:]...)
}