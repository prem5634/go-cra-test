package orders

import "fmt"

// Item is a line item in an order.
type Item struct {
	Name     string
	Price    float64
	Quantity int
}

// ProcessOrderA computes an order total applying a tier discount and tax.
// NOTE: this body is intentionally duplicated in ProcessOrderB (process_b.go)
// to exercise copy-paste detection (duplication pair #1).
func ProcessOrderA(items []Item, customerTier string) (float64, error) {
	if len(items) == 0 {
		return 0, fmt.Errorf("order has no items")
	}
	var subtotal float64
	for _, item := range items {
		if item.Price < 0 {
			return 0, fmt.Errorf("invalid price for %s", item.Name)
		}
		if item.Quantity <= 0 {
			return 0, fmt.Errorf("invalid quantity for %s", item.Name)
		}
		subtotal += item.Price * float64(item.Quantity)
	}
	discount := 0.0
	if customerTier == "gold" {
		discount = 0.15
	} else if customerTier == "silver" {
		discount = 0.10
	} else if subtotal > 1000 {
		discount = 0.05
	}
	discounted := subtotal * (1 - discount)
	tax := discounted * 0.08
	total := discounted + tax
	if total < 0 {
		return 0, fmt.Errorf("computed negative total")
	}
	return total, nil
}
