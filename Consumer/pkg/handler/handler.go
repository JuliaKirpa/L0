package handler

import (
	"NatsMC/Consumer/pkg/service"
	"NatsMC/models"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	s := sse.NewServer(nil)
	defer s.Shutdown()

	router := gin.Default()
	router.StaticFile("/", "static/index.html")
	//что-то абсолютно непонятное, но работает, но вообще не на этом пути ааааааааа
	router.GET("/events/:channel", func(c *gin.Context) {
		s.ServeHTTP(c.Writer, c.Request)
	})
	router.GET("/:id", h.getById)
	//подключаемся к nats-str
	sc, err := stan.Connect("prod", "sub-2")
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("static", func(m *stan.Msg) {
		go func() {
			for {
				s.SendMessage("", sse.SimpleMessage(string(m.Data)))
				time.Sleep(5 * time.Second)
			}
		}()
	})
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

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
