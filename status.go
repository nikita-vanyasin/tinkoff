package tinkoff

const (
	StatusNew             = "NEW"
	StatusAuthorized      = "AUTHORIZED"       // Деньги захолдированы на карте клиента. Ожидается подтверждение операции
	StatusConfirmed       = "CONFIRMED"        // Операция подтверждена
	StatusReversed        = "REVERSED"         // Операция отменена
	StatusRefunded        = "REFUNDED"         // Произведён возврат
	StatusPartialRefunded = "PARTIAL_REFUNDED" // Произведён частичный возврат
	StatusRejected        = "REJECTED"         // Списание денежных средств закончилась ошибкой
)
