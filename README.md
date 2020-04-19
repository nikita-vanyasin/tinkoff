
## Go client for Tinkoff Acquiring API (v2)

### Warning: package API has no stable version yet.
Feel free to contribute. [Roadmap to v1.0.0](#roadmap-to-v100)

API Docs: https://oplata.tinkoff.ru/develop/api/payments/

Based on some code from [koorgoo/tinkoff](https://github.com/koorgoo/tinkoff). Differences:
- support for API v2
- 'reflect' package is not used
- no additional error wrapping
- not all methods are implemented yet :)

## Installation
Use go mod as usual or install the package with dep:
```bash
dep ensure -add github.com/nikita-vanyasin/tinkoff
```

## Package doc
https://pkg.go.dev/github.com/nikita-vanyasin/tinkoff

## Usage example

##### Create and initialize client with API credentials:
```go
client := tinkoff.NewClient(terminalKey, terminalPassword)
```

##### Send Init request:
Init [docs](https://oplata.tinkoff.ru/develop/api/payments/init-description/)
```go
req := &tinkoff.InitRequest{
    Amount:      60000,
    OrderID:     "123456",
    CustomerKey: "123",
    Description: "some really useful product",
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
res, err := a.client.Init(req)
```

##### Handle HTTP-notification:
Parse notification body [(docs)](https://oplata.tinkoff.ru/develop/api/notifications/setup-request/).
Example using [gin](https://github.com/gin-gonic/gin):
```go
router.POST("/payment/notification/tinkoff", func(c *gin.Context) {
    notification, err := a.client.ParseNotification(c.Request.Body)
    if err != nil {
        handleInternalError(c, err)
        return
    }
    
    c.String(http.StatusOK, a.client.GetNotificationSuccessResponse())
}
```

##### Cancel (refund):
Cancel [docs](https://oplata.tinkoff.ru/develop/api/payments/cancel-description/)
```go
req := &tinkoff.CancelRequest{
    PaymentID: "66623",
    Amount: 60000,
}
res, err := a.client.Cancel(req)
```

### Roadmap to v1.0.0
- Accept time.Time object instead of formatted string
- Add more validation (not sure about that one yet)
- GetStatus
- Resend
- Improve usage documentation

### Also planned:
- Add support for more API methods
