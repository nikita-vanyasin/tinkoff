
## Go client for Tinkoff Acquiring API (v2)

### Warning: package has no stable version yet.
Feel free to contribute.

The package allows to send [token-signed](https://oplata.tinkoff.ru/develop/api/request-sign/) requests to Tinkoff Acquiring API and parse incoming HTTP notifications.
Acquiring API Docs: https://oplata.tinkoff.ru/develop/api/payments/

### Contents
- [Installation](#installation)
- [Usage](#usage)
  - [Create client](#create-client)
  - [Handle HTTP notification](#handle-http-notification)
  - [Create payment](#create-payment)
  - [Cancel or refund payment](#cancel-or-refund-payment)
  - [Get payment state](#get-payment-state)
  - [Helper functions](#helper-functions)
- [Roadmap to v1.0.0](#roadmap-to-v100)
- [References](#references)


### Installation
Use **go mod** as usual or install the package with dep:
```bash
dep ensure -add github.com/nikita-vanyasin/tinkoff
```

### Usage

Automatically generated documentation can be found [here](https://pkg.go.dev/github.com/nikita-vanyasin/tinkoff).
Some examples of usage can be found in package `*_test.go` files.


##### Create client
Provide terminal key and password from terminal settings page.
```go
client := tinkoff.NewClient(terminalKey, terminalPassword)
```

##### Handle HTTP notification
[Docs](https://oplata.tinkoff.ru/develop/api/notifications/setup-request/).
Example using [gin](https://github.com/gin-gonic/gin):
```go
router.POST("/payment/notification/tinkoff", func(c *gin.Context) {
    notification, err := client.ParseNotification(c.Request.Body)
    if err != nil {
        handleInternalError(c, err)
        return
    }
    
    // response well-formed body back on success. If you don't do this, the bank will send notification again later
    c.String(http.StatusOK, client.GetNotificationSuccessResponse())
}
```

##### Create payment
[Init](https://oplata.tinkoff.ru/develop/api/payments/init-description/)
```go
req := &tinkoff.InitRequest{
    Amount:      60000,
    OrderID:     "123456",
    CustomerKey: "123",
    Description: "some really useful product",
    RedirectDueDate: tinkoff.Time(time.Now().Add(4 * 24 * time.Hour)),
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
res, err := client.Init(req)
// ...
fmt.Println("payment form url: %s", res.PaymentPageURL)
```

##### Cancel or refund payment
[Cancel](https://oplata.tinkoff.ru/develop/api/payments/cancel-description/)
```go
req := &tinkoff.CancelRequest{
    PaymentID: "66623",
    Amount: 60000,
}
res, err := client.Cancel(req)
```

#### Get payment state
[GetState](https://oplata.tinkoff.ru/develop/api/payments/cancel-description/)
```go
res, err := client.GetState(&tinkoff.GetStateRequest{PaymentID: "3293"})
// ..
if res.Status == tinkoff.StatusConfirmed {
    fmt.Println("payment completed")
}
```

#### Helper functions
- `client.PostRequest` allows you to implement API requests which are not implemented in this package yet (e.g. when Tinkoff Bank adds new method to API).
  Use BaseRequest type to implement any API request:
  ```go
  type myCouponUpgradeRequest struct {
      tinkoff.BaseRequest
      PaymentID string `json:"PaymentId"`
      Coupon    string `json:"coupon"`
  }
  httpResp, err := client.PostRequest(&myCouponUpgradeRequest{PaymentID: "3293", Coupon: "whatever"})
  ```

- `tinkoff.IsRefundableStatus(s)` allows you to check if the payment can be refunded according to policies provided by Tinkoff API doc:
  ```go
  if !tinkoff.IsRefundableStatus(payment.Status) {
      return errors.New("payment can't be refunded")
  }
  ``` 

### Roadmap to v1.0.0
- Resend
- FinishAuthorize
- Confirm
- Submit3DSAuthorization


### References
The code in this repo based on some code from [koorgoo/tinkoff](https://github.com/koorgoo/tinkoff). Differences:
- support for API v2
- 'reflect' package is not used
- no additional error wrapping

More useful links:
- Official [Tinkoff Acquiring API SDK for Android (java)](https://github.com/TinkoffCreditSystems/tinkoff-asdk-android)
- Official [PHP client and integration examples](https://oplata.tinkoff.ru/develop/api/examples/)
