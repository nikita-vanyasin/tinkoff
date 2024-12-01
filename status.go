package tinkoff

const (
	StatusNew             = "NEW"              // Создан
	StatusFormShowed      = "FORM_SHOWED"      // Платежная форма открыта покупателем
	StatusDeadlineExpired = "DEADLINE_EXPIRED" // Просрочен
	StatusCanceled        = "CANCELED"         // Отменен
	// not used. removal was not documented
	StatusPreauthorizing  = "PREAUTHORIZING"   // Проверка платежных данных
	StatusAuthorizing     = "AUTHORIZING"      // Резервируется
	StatusAuthorized      = "AUTHORIZED"       // Зарезервирован
	StatusAuthFail        = "AUTH_FAIL"        // Не прошел авторизацию
	StatusRejected        = "REJECTED"         // Отклонен
	Status3DSChecking     = "3DS_CHECKING"     // Проверяется по протоколу 3-D Secure
	Status3DSChecked      = "3DS_CHECKED"      // Проверен по протоколу 3-D Secure
	StatusReversing       = "REVERSING"        // Резервирование отменяется
	StatusPartialReversed = "PARTIAL_REVERSED" // Частичный возврат по авторизованному платежу завершился успешно.
	StatusReversed        = "REVERSED"         // Резервирование отменено
	StatusConfirming      = "CONFIRMING"       // Подтверждается
	StatusConfirmed       = "CONFIRMED"        // Подтвержден
	StatusRefunding       = "REFUNDING"        // Возвращается
	StatusQRRefunding     = "ASYNC_REFUNDING"  // Возврат QR
	StatusPartialRefunded = "PARTIAL_REFUNDED" // Возвращен частично
	StatusRefunded        = "REFUNDED"         // Возвращен полностью
)
