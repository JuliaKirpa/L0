package main

import (
	"github.com/nats-io/stan.go"
	"strconv"
	"time"
)

func main() {
	sc, err := stan.Connect("prod", "static")
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	for i := 0; i <= 100; i++ {
		sc.Publish("static", []byte("Hello from iteration "+strconv.Itoa(i)))
		time.Sleep(2 * time.Second)
	}
}

//var order = &models.Order{
//	OrderUid:    "b563feb7b2b84b6test",
//	TrackNumber: "WBILMTESTTRACK",
//	Entry:       "WBIL",
//	Delivery: models.Delivery{
//		Name:    "Test Testov",
//		Phone:   "+9720000000",
//		Zip:     "2639809",
//		City:    "Kiryat Mozkin",
//		Address: "Ploshad Mira 15",
//		Region:  "Kraiot",
//		Email:   "test@gmail.com",
//	},
//	Payment: models.Payment{
//		Transaction:  "b563feb7b2b84b6test",
//		RequestId:    "",
//		Currency:     "USD",
//		Provider:     "wbpay",
//		Amount:       1817,
//		PaymentDt:    1637907727,
//		Bank:         "alpha",
//		DeliveryCost: 1500,
//		GoodsTotal:   315,
//		CustomFee:    0,
//	},
//	Items: []models.Items{{
//		ChrtId:      9934930,
//		TrackNumber: "WBILMTESTTRACK",
//		Price:       453,
//		Rid:         "ab4219087a764ae0btest",
//		Name:        "Mascaras",
//		Sale:        30,
//		Size:        "",
//		TotalPrice:  317,
//		NmId:        2389212,
//		Brand:       "Vivienne Sabo",
//		Status:      202,
//	}},
//	Locale:            "en",
//	InternalSignature: "",
//	CustomerId:        "test",
//	DeliveryService:   "meest",
//	Shardkey:          "9",
//	SmId:              99,
//	DateCreated:       time.Now(),
//	OofShard:          "1",
//}
