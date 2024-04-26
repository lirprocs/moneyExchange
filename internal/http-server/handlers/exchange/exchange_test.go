package exchange

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"moneyExchangeTest/internal/log/testlog"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestExchangeHandler(t *testing.T) {
	cases := []struct {
		Name      string
		Amount    int
		Banknotes []int
		Expected  [][]int
	}{
		{
			Name:      "400",
			Amount:    400,
			Banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			Expected: [][]int{
				{200, 200},
				{200, 100, 100},
				{200, 100, 50, 50},
				{200, 50, 50, 50, 50},
				{100, 100, 100, 100},
				{100, 100, 100, 50, 50},
				{100, 100, 50, 50, 50, 50},
				{100, 50, 50, 50, 50, 50, 50},
				{50, 50, 50, 50, 50, 50, 50, 50},
			},
		},
		{
			Name:      "200",
			Amount:    200,
			Banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			Expected: [][]int{
				{200},
				{100, 100},
				{100, 50, 50},
				{50, 50, 50, 50},
			},
		},
		{
			Name:      "100",
			Amount:    100,
			Banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			Expected: [][]int{
				{100},
				{50, 50},
			},
		},
		{
			Name:      "Пустой список банкнот",
			Amount:    200,
			Banknotes: []int{},
			Expected:  nil,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			reqBody, err := json.Marshal(Request{Amount: tc.Amount, Banknotes: tc.Banknotes})
			require.NoError(t, err)

			req := httptest.NewRequest("POST", "/exchange", bytes.NewBuffer(reqBody))
			w := httptest.NewRecorder()

			New(testlog.NewDiscardLogger())(w, req)

			var resp Response
			err = json.Unmarshal(w.Body.Bytes(), &resp)
			require.NoError(t, err)

			require.True(t, reflect.DeepEqual(tc.Expected, resp.Exchanges))
		})
	}
}
