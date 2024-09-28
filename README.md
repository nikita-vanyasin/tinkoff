
# Golang Tinkoff Acquiring API (v2) client

The package allows to send token-signed requests to Tinkoff Acquiring API and parse incoming HTTP notifications.

Acquiring API Docs: https://www.tinkoff.ru/kassa/dev/payments/


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
```bash
go get github.com/rentifly/tinkoff@latest
```

## Usage

Automatically generated documentation can be found [here](https://pkg.go.dev/github.com/rentifly/tinkoff).

Some examples of usage can be found in `*_test.go` files.


#### Create client
Provide terminal key and password from terminal settings page.
```go
client := tinkoff.NewClientWithOptions(
  WithTerminalKey(terminalKey),
  WithPassword(password),
  // Optional override HTTP client:
  // WithHTTPClient(myClient),
  // Optional override base Tinkoff Acquiring API URL:
  // WithBaseURL(myURL),
)
```

#### Handle HTTP notification
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

// Set timeout for Init request:
ctx, cancel := context.WithTimeout(ctx, time.Second*10)
defer cancel()

// Execute:
res, err := client.InitWithContext(ctx, req)
// ...
fmt.Println("payment form url: %s", res.PaymentPageURL)
```

#### Create QR
```go
req := &tinkoff.InitRequest{
    Amount:      1000,                 // минимум 1000 копеек 
    OrderID:     "123456",
    Data: map[string]string{"": "",},  // nil - недопустим.
res, err := client.Init(req)
// ...
gqr := &tinkoff.GetQrRequest{
    PaymentID: res.PaymentID,
}
resQR, err := client.GetQRWithContext(ctx, gqr)
```

#### Cancel or refund payment
```go
req := &tinkoff.CancelRequest{
    PaymentID: "66623",
    Amount: 60000,
}
res, err := client.CancelWithContext(ctx, req)
```

#### Get payment state
```go
res, err := client.GetStateWithContext(ctx, &tinkoff.GetStateRequest{PaymentID: "3293"})
// ...
if res.Status == tinkoff.StatusConfirmed {
    fmt.Println("payment completed")
}
```

#### Confirm two-step payment
```go
res, err := client.ConfirmWithContext(ctx, &tinkoff.ConfirmRequest{PaymentID: "3294"})
// ...
if res.Status == tinkoff.StatusConfirmed {
    fmt.Println("payment completed")
}
```

#### Resend notifications
```go
res, err := c.ResendWithContext(ctx)
// ...
fmt.Println("resend has been scheduled for %d notifications", res.Count)
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
httpResp, err := client.PostRequestWithContext(ctx, &myCouponUpgradeRequest{PaymentID: "3293", Coupon: "whatever"})
```

## References
The code in this repo based on some code from [koorgoo/tinkoff](https://github.com/koorgoo/tinkoff). Differences:
- Support for API v2
- 'reflect' package is not used. Zero dependencies.
- No additional error wrapping

## Contribution
All contributions are welcome! There are plenty of API methods that are not implemented yet due to their rare use-cases:
- FinishAuthorize
- Submit3DSAuthorization
- Charge
- AddCustomer / GetCustomer / RemoveCustomer
- GetCardList / RemoveCard 
 
