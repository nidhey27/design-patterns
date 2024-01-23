package main

type OrderExecutor struct {
}

func (oe *OrderExecutor) PlaceOrder(userID string, order *Order) {
	exchangeConnector, _ := GetExchangeConnector()
	exchangeConnector.SendOrderToExchange(userID, order)
}
