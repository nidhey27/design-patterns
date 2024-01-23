package main

func main() {
	user := &User{
		UserID: "Nidhey",
	}

	system := &System{
		Users: map[string]*User{
			"Nidhey": user,
		},
	}

	stock := &Stock{
		Exch:  NSE,
		Name:  "Reliance",
		Price: 1299.00,
	}

	order := &Order{
		Type:      BUY,
		OrderType: MARKET,
		Quantity:  10,
		Price:     1299.00,
		Stock:     stock,
		Exch:      NSE,
	}

	orderManager := &OrderManager{
		OrderValidator: &OrderValidator{
			System: system,
		},
		OrderExecutor: &OrderExecutor{},
	}

	orderManager.PlaceOrder(user.UserID, order)
}
