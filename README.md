
# Golang Tinkoff Acquiring API (v2) client

The package allows to send [token-signed](https://oplata.tinkoff.ru/develop/api/request-sign/) requests to Tinkoff Acquiring API and parse incoming HTTP notifications.

Acquiring API Docs: https://oplata.tinkoff.ru/develop/api/payments/


## Contents
- [Installation](#installation)
- [Usage](#usage)
  - [Create client](#create-client)
  - [Handle HTTP notification](#handle-http-notification)
  - [Create payment](#create-payment)
  - [Cancel or refund payment](#cancel-or-refund-payment)
  - [Get payment state](#get-payment-state)
  - [Resend notifications](#resend-notifications)
  - [Confirm two-step payment](#confirm-two-step-payment)
  - [Helper functions](#helper-functions)
- [References](#references)
- [Contribution](#contribution)


## Installation
Use **go mod** as usual or install the package with **dep**:
```bash
dep ensure -add github.com/nikita-vanyasin/tinkoff
```

## Usage

Automatically generated documentation can be found [here](https://pkg.go.dev/github.com/nikita-vanyasin/tinkoff).

Some examples of usage can be found in `*_test.go` files.


#### Create client
Provide terminal key and password from terminal settings page.
```go
client := tinkoff.NewClient(terminalKey, terminalPassword)
```

#### Handle HTTP notification
[Docs](https://oplata.tinkoff.ru/develop/api/notifications/setup-request/).
Example using [gin](https://github.com/gin-gonic/gin):
```go
router.POST("/payment/notification/tinkoff", func(c *gin.Context) {
    notification, err := client.ParseNotification(c.Request.Body)
    if err != nil {
        handleInternalError(c, err)
        return
    }

    // handle notification, e.g. update payment status in your DB

    // response well-formed body back on success. If you don't do this, the bank will send notification again later
    c.String(http.StatusOK, client.GetNotificationSuccessResponse())
}
```

#### Create payment
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
        Payments: &tinkoff.ReceiptPayments{
            Electronic: 60000,
        },
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

#### Create QR
// обязательные данные для инициализации платежной сессии
// required data to initialize the payment session
```go
req := &tinkoff.InitRequest{
    Amount:      1000,                 // минимум 1000 копеек 
    OrderID:     "123456",
    Data: map[string]string{"": "",},  // nil - недопустим.
res, err := client.Init(req)
gqr := &tinkoff.GetQrRequest{
    PaymentID: res.PayID,
}
resQR, errQ := client.GetQR(gqr)
```

#### Cancel or refund payment
[Cancel](https://oplata.tinkoff.ru/develop/api/payments/cancel-description/)
```go
req := &tinkoff.CancelRequest{
    PaymentID: "66623",
    Amount: 60000,
}
res, err := client.Cancel(req)
```

#### Get payment state
[GetState](https://oplata.tinkoff.ru/develop/api/payments/getstate-description/)
```go
res, err := client.GetState(&tinkoff.GetStateRequest{PaymentID: "3293"})
// ...
if res.Status == tinkoff.StatusConfirmed {
    fmt.Println("payment completed")
}
```

#### Confirm two-step payment
[Confirm](https://oplata.tinkoff.ru/develop/api/payments/confirm-description/)
```go
res, err := client.Confirm(&tinkoff.ConfirmRequest{PaymentID: "3294"})
// ...
if res.Status == tinkoff.StatusConfirmed {
    fmt.Println("payment completed")
}
```

#### Resend notifications
[Resend](https://oplata.tinkoff.ru/develop/api/payments/resend-description/)
```go
res, err := c.Resend()
// ...
fmt.Println("resend scheduled for %d notifications", res.Count)
```

#### Helper functions
`client.PostRequest` allows you to implement API requests which are not implemented in this package yet (e.g. when Tinkoff Bank adds new method to API).
Use BaseRequest type to implement any API request:
```go
type myCouponUpgradeRequest struct {
  tinkoff.BaseRequest
  PaymentID string `json:"PaymentId"`
  Coupon    string `json:"coupon"`
}
httpResp, err := client.PostRequest(&myCouponUpgradeRequest{PaymentID: "3293", Coupon: "whatever"})
```

## References
The code in this repo based on some code from [koorgoo/tinkoff](https://github.com/koorgoo/tinkoff). Differences:
- Support for API v2
- 'reflect' package is not used
- No additional error wrapping

More useful links:
- Official [Tinkoff Acquiring API SDK for Android (java)](https://github.com/TinkoffCreditSystems/tinkoff-asdk-android)
- Official [PHP client and integration examples](https://oplata.tinkoff.ru/develop/api/examples/)

## Contribution
All contributions are welcome! There are plenty of API methods that are not implemented yet due to their rare use-cases:
- FinishAuthorize
- Submit3DSAuthorization
- Charge
- AddCustomer / GetCustomer / RemoveCustomer
- GetCardList / RemoveCard 
 
