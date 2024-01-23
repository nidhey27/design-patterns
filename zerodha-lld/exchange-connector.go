package main

import (
	"fmt"
	"sync"
)

var (
	exchangeConnector *ExchangeConnector
	lock              = &sync.Mutex{}
)

type ExchangeConnector struct {
}

func GetExchangeConnector() (*ExchangeConnector, error) {
	err := fmt.Errorf("")
	if exchangeConnector == nil {
		lock.Lock()
		defer lock.Unlock()
		if exchangeConnector == nil {
			exchangeConnector = &ExchangeConnector{}
		} else {
			err = fmt.Errorf("ExchangeConnector instance already created.")
		}
	} else {
		err = fmt.Errorf("ExchangeConnector instance already created.")
	}
	return exchangeConnector, err
}

func (e *ExchangeConnector) SendOrderToExchange(userID string, order *Order) {
	fmt.Println("Order is sent to EXCHANGE. Congratulations!")
}
