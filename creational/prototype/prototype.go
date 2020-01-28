package creational

import (
	"errors"
	"fmt"
)

type shirtCloner interface {
	Clone(m int) (ItemInfoGetter, error)
}

// ShirtsCache is exposed to be used globally
type ShirtsCache struct{}

// Clone is exposed to be used globally
func (s *ShirtsCache) Clone(color string) (ItemInfoGetter, error) {
	switch color {
	case "white":
		item := *white
		return &item, nil
	case "black":
		item := *black
		return &item, nil
	case "blue":
		item := *blue
		return &item, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

// ItemInfoGetter is exposed to be used globally
type ItemInfoGetter interface {
	Info() string
}

// Shirt is exposed to be used globally
type Shirt struct {
	Price float32
	SKU   string
	Color string
}

var white = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: "white",
}

var black = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: "black",
}

var blue = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: "blue",
}

// Info is exposed to be used globally
func (s *Shirt) Info() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %s that costs %f\n", s.SKU, s.Color, s.Price)
}
