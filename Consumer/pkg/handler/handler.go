package handler

import (
	"NatsMC/models"
	"encoding/json"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	s := sse.NewServer(nil)
	defer s.Shutdown()

	router := gin.Default()
	router.StaticFile("/", "static/index.html")

	router.GET("/events/:channel", func(c *gin.Context) {
		s.ServeHTTP(c.Writer, c.Request)
	})
	router.GET("/events/:id", h.getById)

	byte, _ := json.Marshal(order)

	go func() {
		for {
			s.SendMessage("", sse.SimpleMessage(string(byte)))
			time.Sleep(5 * time.Second)
		}
	}()

	log.Println("Listening at :3000")
	router.Run("localhost:3000")

	return router
}

var order = &models.Order{
	OrderUid:    "b563feb7b2b84b6test",
	TrackNumber: "WBILMTESTTRACK",
	Entry:       "WBIL",
	Delivery: models.Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	},
	Payment: models.Payment{
		Transaction:  "b563feb7b2b84b6test",
		RequestId:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDt:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   315,
		CustomFee:    0,
	},
	Items: []models.Items{{
		ChrtId:      9934930,
		TrackNumber: "WBILMTESTTRACK",
		Price:       453,
		Rid:         "ab4219087a764ae0btest",
		Name:        "Mascaras",
		Sale:        30,
		Size:        "",
		TotalPrice:  317,
		NmId:        2389212,
		Brand:       "Vivienne Sabo",
		Status:      202,
	}},
	Locale:            "en",
	InternalSignature: "",
	CustomerId:        "test",
	DeliveryService:   "meest",
	Shardkey:          "9",
	SmId:              99,
	DateCreated:       time.Now(),
	OofShard:          "1",
}
