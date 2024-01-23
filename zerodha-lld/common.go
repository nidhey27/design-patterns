package main

type TXN_TYPE int

const (
	BUY TXN_TYPE = iota
	SELL
)

type ORDER_TYPE int

const (
	MARKET ORDER_TYPE = iota
	LIMIT
)

type EXCH int

const (
	NSE EXCH = iota
	BSE
)

type ORDER_STATUS int

const (
	OPEN ORDER_STATUS = iota
	PARTIALLY_DONE
	DONE
	CANCELLED
)
