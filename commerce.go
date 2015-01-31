package faker

import (
	"math"
	"math/rand"
	"strings"
)

type Commerce struct{}

// Example:
//  Commerce{}.Color() // lime
func (c Commerce) Color() string {
	return Fetch("commerce.color")
}

// Example:
//  Commerce{}.Department() // Electronics, Health & Baby
func (c Commerce) Department() string {
	n := RandomInt(1, 3)

	deps := make([]string, n)
	for i, _ := range deps {
		d := Fetch("commerce.department")
		for includesString(deps, d) {
			d = Fetch("commerce.department")
		}
		deps[i] = d
	}

	if n > 1 {
		sep := Fetch("separator")
		res := strings.Join([]string{
			strings.Join(deps[0:len(deps)-1], ", "),
			deps[len(deps)-1],
		}, sep)
		return res
	} else {
		return deps[0]
	}
}

// Example:
//  Commerce{}.ProductName() // Ergonomic Granite Shoes
func (c Commerce) ProductName() string {
	words := []string{
		Fetch("commerce.product_name.adjective"),
		Fetch("commerce.product_name.material"),
		Fetch("commerce.product_name.product"),
	}
	return strings.Join(words, " ")
}

// Example:
//  Commerce{}.Price() // 97.79
func (c Commerce) Price() float32 {
	return float32(math.Floor(rand.Float64()*10000.0)/100.0 + 0.01)
}
