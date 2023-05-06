package strategy

type Cart struct {
	total float32
}

func (c *Cart) Pay(paymentStrategy PayementStrategy) {
	paymentStrategy.pay(int(c.total))
}

func (c *Cart) SetCart(amount float32) {
	c.total = amount
}
