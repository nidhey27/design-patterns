package main

type OrderManager struct {
	OrderValidator *OrderValidator
	OrderExecutor  *OrderExecutor
}

func (om *OrderManager) PlaceOrder(userID string, order *Order) {
	if om.OrderValidator.ValidateOrder(userID, order) {
		om.OrderExecutor.PlaceOrder(userID, order)
	}
}
