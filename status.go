package tinkoff

const (
	StatusNew             = "NEW"              // Создан
	StatusFormShowed      = "FORMSHOWED"       // Платежная форма открыта покупателем
	StatusDeadlineExpired = "DEADLINE_EXPIRED" // Просрочен
	StatusCanceled        = "CANCELED"         // Отменен
	StatusPreauthorizing  = "PREAUTHORIZING"   // Проверка платежных данных
	StatusAuthorizing     = "AUTHORIZING"      // Резервируется
	StatusAuthorized      = "AUTHORIZED"       // Зарезервирован
	StatusAuthFail        = "AUTH_FAIL"        // Не прошел авторизацию
	StatusRejected        = "REJECTED"         // Отклонен
	Status3DSChecking     = "3DS_CHECKING"     // Проверяется по протоколу 3-D Secure
	Status3DSChecked      = "3DS_CHECKED"      // Проверен по протоколу 3-D Secure
	StatusReversing       = "REVERSING"        // Резервирование отменяется
	StatusReversed        = "REVERSED"         // Резервирование отменено
	StatusConfirming      = "CONFIRMING"       // Подтверждается
	StatusConfirmed       = "CONFIRMED"        // Подтвержден
	StatusRefunding       = "REFUNDING"        // Возвращается
	StatusPartialRefunded = "PARTIAL_REFUNDED" // Возвращен частично
	StatusRefunded        = "REFUNDED"         // Возвращен полностью
)
