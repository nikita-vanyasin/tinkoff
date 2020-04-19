package tinkoff

const (
	StatusNew             = "NEW"              // Платеж создан в системе банка
	StatusAuthorized      = "AUTHORIZED"       // Деньги захолдированы на карте клиента. Ожидается подтверждение операции
	StatusConfirmed       = "CONFIRMED"        // Операция подтверждена
	StatusReversed        = "REVERSED"         // Операция отменена
	StatusRefunded        = "REFUNDED"         // Произведён возврат
	StatusPartialRefunded = "PARTIAL_REFUNDED" // Произведён частичный возврат
	StatusRejected        = "REJECTED"         // Списание денежных средств закончилась ошибкой
	StatusCanceled        = "CANCELED"         // Неоплаченный платеж отменен
)

func IsRefundableStatus(status string) bool {
	switch status {
	case StatusNew:
	case StatusAuthorized:
	case StatusConfirmed:
		return true
	}
	return false
}
