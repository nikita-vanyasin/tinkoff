
## Go client for Tinkoff Acquiring API (v2)

API Docs: https://oplata.tinkoff.ru/landing/develop/documentation

Based on some code from [koorgoo/tinkoff](https://github.com/koorgoo/tinkoff)

##### Differences:
1) support for API v2
1) 'reflect' package is not used
1) no additional error wrapping
1) not all methods are implemented :)

### Currently implemented features:
1) Init [(docs)](https://oplata.tinkoff.ru/landing/develop/documentation/Init)
1) Parse notification body [(docs)](https://oplata.tinkoff.ru/landing/develop/notifications/http)


### Installation
Use dep:
```bash
dep ensure -add github.com/nikita-vanyasin/tinkoff
``` 


### Usage example

Create and initialize client with API credentials:
```go

client := tinkoff.NewClient(terminalKey, terminalPassword)
```


Send Init request:
```go
	req := &tinkoff.InitRequest{
		Amount:      60000,
		OrderID:     "123456",
		CustomerKey: "123",
		Description: "some really useful product",
		Receipt: tinkoff.Receipt{
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
	status, paymentID, paymentURL, err := a.client.Init(req)
```

Handle HTTP-notification (example using [gin](https://github.com/gin-gonic/gin)):
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
