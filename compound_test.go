package tinkoff_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/nikita-vanyasin/tinkoff"
)

func TestInitThenCancel(t *testing.T) {
	c := helperCreateClient(t)

	// create new payment
	orderID := strconv.FormatInt(time.Now().UnixNano(), 10)
	initReq := &tinkoff.InitRequest{
		Amount:      60000,
		OrderID:     orderID,
		CustomerKey: "123",
		Description: "some really useful product",
		PayType:     tinkoff.PayTypeTwoSteps,
		Receipt: &tinkoff.Receipt{
			Email: "user@example.com",
			Items: []*tinkoff.ReceiptItem{
				{
					Price:    60000,
					Quantity: "1",
					Amount:   60000,
					Name:     "Product #1",
					Tax:      tinkoff.VATNone,
				},
			},
			Taxation: tinkoff.TaxationUSNIncome,
		},
		Data: map[string]string{
			"custom data field 1": "aasd6da78dasd9",
			"custom data field 2": "0",
		},
	}
	initRes, err := c.Init(initReq)
	assertNotError(t, err)

	assertEq(t, true, initRes.Success)

	assertEq(t, tinkoff.StatusNew, initRes.Status)
	assertNotEmptyString(t, initRes.OrderID)
	assertNotEmptyString(t, initRes.PaymentURL)

	// cancel it
	req := &tinkoff.CancelRequest{
		PaymentID: initRes.PaymentID,
		Amount:    60000,
		Receipt: &tinkoff.Receipt{
			Email: "user@example.com",
			Items: []*tinkoff.ReceiptItem{
				{
					Price:    60000,
					Quantity: "1",
					Amount:   60000,
					Name:     "Product #1",
					Tax:      tinkoff.VATNone,
				},
			},
			Taxation: tinkoff.TaxationUSNIncome,
		},
	}
	cancelRes, err := c.Cancel(req)
	assertNotError(t, err)

	assertEq(t, true, cancelRes.Success)

	assertEq(t, tinkoff.StatusCanceled, cancelRes.Status)
	assertEq(t, initRes.PaymentID, cancelRes.PaymentID)
	assertEq(t, orderID, cancelRes.OrderID)

}
