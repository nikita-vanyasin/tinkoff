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
				tinkoff.WithPassword("0x6sliXzI4fMbxxx"),
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
				"Token": "c0b3457bd4822b3cc823a9eb8d6b33733fb3225e4fb7b7039e5807987502648f"
			}`),
			want: &tinkoff.Notification{
				TerminalKey:    "1726675329428DEMO",
				OrderID:        "9863471b-8d42-4617-9073-11339159aba9",
				Success:        true,
				Status:         tinkoff.StatusAuthorized,
				PaymentID:      5974256058,
				ErrorCode:      "0",
				Amount:         19900,
				RebillID:       1741204285536,
				CardID:         475738102,
				PAN:            "430000******0777",
				DataStr:        "",
				Data:           nil,
				Token:          "c0b3457bd4822b3cc823a9eb8d6b33733fb3225e4fb7b7039e5807987502648f",
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
