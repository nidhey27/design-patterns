package main

type Order struct {
	Type      TXN_TYPE
	OrderType ORDER_TYPE
	Quantity  int
	Price     float64
	Stock     *Stock
	Exch      EXCH
	// transactions
	// status ORDER_TYP
	// time
}

func (o *Order) GetTxnType() TXN_TYPE {
	return o.Type
}
