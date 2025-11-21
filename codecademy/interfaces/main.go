package main

import "fmt"

// Define the interface
type Discount interface {
	Apply(amount float64) float64
}

// Implementation 3: NoDiscount
type NoDiscount struct {}

func (d NoDiscount) Apply(amount float64) float64 {
	return amount
}

// Store struct that delegates to Discount
type Store struct {
	discountType Discount
}

// Function to use the Discount interface
func (s Store) printDiscountedPrice(amount float64) {
	fmt.Printf("Discounted Price: %.2f\n", s.discountType.Apply(amount))
}

func main() {
	discount := NoDiscount{}
	//discount := PercentageDiscount {percent 10}
	// discount := FixedAmountDiscount {amount: 20}
	store := Store{discountType: discount}

	initialPrice := 100.0
	store.printDiscountedPrice(initialPrice)

}