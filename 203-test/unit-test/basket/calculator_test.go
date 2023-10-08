package basket

import "testing"
import asrt "github.com/stretchr/testify/assert"

func TestAdd(t *testing.T) {
	x, y := 2, 5

	c := Calculate{}

	actual := c.Add(x, y)
	expected := 7

	if actual != expected {
		t.Errorf("no success")
	}
}

func TestSubtract(t *testing.T) {
	c := Calculate{}
	tables := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 2, 0},
		{6, 2, 4},
		{19, 8, 11},
	}

	for _, value := range tables {
		actual := c.Subtract(value.x, value.y)

		if actual != value.expected {
			t.Errorf("no success")
		}
	}
}

func TestMultiply(t *testing.T) {
	c := Calculate{}

	assert := asrt.New(t)

	tables := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 2, 4},
		{6, 2, 12},
		{10, 8, 80},
	}

	for _, value := range tables {
		actual := c.Multiply(value.x, value.y)
		assert.Equal(value.expected, actual)
	}
}

func TestDivide(t *testing.T) {

}
