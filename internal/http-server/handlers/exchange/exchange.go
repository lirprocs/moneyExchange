package exchange

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type Request struct {
	Amount    int   `json:"amount" validate:"required"`
	Banknotes []int `json:"banknotes" validate:"required"`
}

type Response struct {
	Exchanges [][]int `json:"exchanges"`
}

func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.exchange.New"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", err)

			render.JSON(w, r, fmt.Errorf("failed to decode request"))

			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			http.Error(w, "missing a required field", http.StatusBadRequest)
			return
		}

		var exchange [][]int
		calculateExchange(req.Amount, req.Banknotes, []int{}, &exchange)

		log.Info("The count has been completed")
		render.JSON(w, r, Response{Exchanges: exchange})
		//log.Info("", exchange)
	}
}

func calculateExchange(amount int, banknotes []int, currentExchange []int, result *[][]int) {
	if amount == 0 {
		temp := make([]int, len(currentExchange))
		copy(temp, currentExchange)
		*result = append(*result, temp)
		return
	}

	for i, note := range banknotes {
		if note <= amount {
			calculateExchange(amount-note, banknotes[i:], append(currentExchange, note), result)
		}
	}
}
