package tinkoff_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/nikita-vanyasin/tinkoff"
)

func TestClient_ParseNotification(t *testing.T) {
	tests := []struct {
		name    string
		client  *tinkoff.Client
		body    []byte
		want    *tinkoff.Notification
		wantErr bool
	}{
		{
			name: "with_rebill",
			client: tinkoff.NewClientWithOptions(
				tinkoff.WithTerminalKey("1726675329428DEMO"),
				tinkoff.WithPassword("xxxxxxxxxxxxxxxx"),
			),
			body: []byte(`{
				"TerminalKey": "1726675329428DEMO",
				"OrderId": "9863471b-8d42-4617-9073-11339159aba9",
				"Success": true,
				"Status": "AUTHORIZED",
				"PaymentId": 5974256058,
				"ErrorCode": "0",
				"Amount": 19900,
				"CardId": 475738102,
				"Pan": "430000******0777",
				"ExpDate": "1230",
				"RebillId": 1741204285536,
				"Token": "16cfc8cd6e8f8fc1d404e3f9984c8058ecdf6eab622a9cc73c2e3824e47124ac"
			}`),
			want: &tinkoff.Notification{
				TerminalKey:    "1726675329428DEMO",
				OrderID:        "9863471b-8d42-4617-9073-11339159aba9",
				Success:        true,
				Status:         tinkoff.StatusAuthorized,
				PaymentID:      5974256058,
				ErrorCode:      "0",
				Amount:         19900,
				RebillID:       "1741204285536",
				RebillIDUInt64: 1741204285536,
				CardID:         475738102,
				PAN:            "430000******0777",
				DataStr:        "",
				Data:           nil,
				Token:          "16cfc8cd6e8f8fc1d404e3f9984c8058ecdf6eab622a9cc73c2e3824e47124ac",
				ExpirationDate: "1230",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.client.ParseNotification(bytes.NewReader(tt.body))
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ParseNotification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ParseNotification() = %v, want %v", got, tt.want)
			}
		})
	}
}
