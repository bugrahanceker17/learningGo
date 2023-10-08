package basket

import (
	"fmt"
	asrt "github.com/stretchr/testify/assert"
	"testing"
)

type TestCalculator struct {
	add      map[string]int
	subtract map[string]int
	multiply map[string]int
	divide   map[string]float64
}

func (c TestCalculator) Add(x, y int) int {
	return c.add[fmt.Sprintf("%d%d", x, y)]
}

func (c TestCalculator) Subtract(x, y int) int {
	return c.subtract[fmt.Sprintf("%d%d", x, y)]
}

func (c TestCalculator) Multiply(x, y int) int {
	return c.multiply[fmt.Sprintf("%d%d", x, y)]
}

func (c TestCalculator) Divide(x, y int) float64 {
	return c.divide[fmt.Sprintf("%d%d", x, y)]
}

func (c TestCalculator) On(fun string, x, y int, rt float64) {
	switch fun {
	case "Add":
		c.add[fmt.Sprintf("%d%d", x, y)] = int(rt)
	case "Subtract":
		c.subtract[fmt.Sprintf("%d%d", x, y)] = int(rt)
	case "Multiply":
		c.multiply[fmt.Sprintf("%d%d", x, y)] = int(rt)
	case "Divide":
		c.divide[fmt.Sprintf("%d%d", x, y)] = rt
	}
}

func TestAmount(t *testing.T) {

	assert := asrt.New(t)

	testCalculator := TestCalculator{
		add:      map[string]int{},
		subtract: map[string]int{},
		multiply: map[string]int{},
		divide:   map[string]float64{},
	}

	d := NewMinPriceDiscount(11, testCalculator)

	tables := []struct {
		amount     float64
		percentage float64
		expected   float64
	}{
		{100, 20, 80},
		{10, 20, 10},
		{100, 120, 100},
	}

	testCalculator.On("Subtract", 100, 20, 80)

	fmt.Println("testing...")
	for _, value := range tables {
		actual := d.Amount(value.amount, value.percentage)

		assert.Equal(value.expected, actual)
	}
}
