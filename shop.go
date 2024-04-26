package tinkoff

type Shop struct {
	ShopCode string `json:"ShopCode,omitempty"` // Код магазина. Для параметра ShopСode необходимо использовать значение параметра Submerchant_ID, полученного при регистрации через xml.
	Amount   uint64 `json:"Amount,omitempty"`   // Сумма перечисления в копейках по реквизитам ShopCode за вычетом Fee
	Name     string `json:"Name,omitempty"`     // Наименование позиции
	Fee      string `json:"Fee,omitempty"`      // Часть суммы Операции оплаты или % от суммы Операции оплаты. Fee удерживается из возмещения третьего лица (ShopCode) в пользу Предприятия по операциям оплаты.
}
