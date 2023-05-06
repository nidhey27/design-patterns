package strategy

type PayementStrategy interface {
	pay(amount int)
}
