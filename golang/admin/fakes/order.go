package fakes

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/satori/go.uuid"
)

func GenOrder() string {
	order_no := uuid.NewV4().String()
	orderTime := RandTime(365*86400, false).Unix() * 1000
	salesman := fmt.Sprintf("%s %s", faker.FirstName(), faker.LastName())
	price := RandFloat(500, 90000, true)
	status := RandItem([]string{"pending", "payed", "finish"})
	return ReduceBlanks(fmt.Sprintf(`{
	"order_no": "%s",
	"timestamp": %d,
	"username": "%s",
	"price": %.1f,
	"status": "%s"
}`, order_no, orderTime, salesman, price, status))
}
