package tests

import (
	"github.com/gavv/httpexpect/v2"
	"math/rand"
	"moneyExchangeTest/internal/http-server/handlers/exchange"
	"net/url"
	"testing"
)

const (
	host = "localhost:8088"
)

func TestURLShortener_HappyPath(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
	}
	e := httpexpect.Default(t, u.String())

	desiredNumbers := []int{50, 100, 150, 200, 300, 400, 500}
	randomIndex := rand.Intn(len(desiredNumbers))
	randomNumber := desiredNumbers[randomIndex]

	e.POST("/exchange").
		WithJSON(exchange.Request{
			Amount:    randomNumber,
			Banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50}}).
		Expect().
		Status(200).
		JSON().Object().
		ContainsKey("exchanges")
}
