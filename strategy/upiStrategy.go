package strategy

import "fmt"

type UPIStrategy struct {
	upiID string
}

func (u *UPIStrategy) SetUPIStrategy(UPIid string) {
	u.upiID = UPIid

}

func (u *UPIStrategy) pay(amount int) {
	fmt.Printf("Paying $%v by UPI %s\n", amount, u.upiID)
}
