package payments

import "fmt"

type CardPayment struct {
	CardNumber string
}

func (c *CardPayment) Pay(amount float64) bool {
	fmt.Printf("Paid ₹%.2f via Card ending in %s\n", amount, c.CardNumber[len(c.CardNumber)-4:])
	return true
}
