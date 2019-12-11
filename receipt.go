package tinkoff

const (
	TaxationOSN              = "osn"                // общая СН
	TaxationUSNIncome        = "usn_income"         // упрощенная СН (доходы)
	TaxationUSNIncomeOutcome = "usn_income_outcome" // упрощенная СН (доходы минус расходы)
	TaxationENVD             = "envd"               // единый налог на вмененный доход
	TaxationESN              = "esn"                // единый сельскохозяйственный налог
	TaxationPatent           = "patent"             // патентная СН
)

const (
	VATNone = "none"   // без НДС
	VAT0    = "vat0"   // НДС по ставке 0%
	VAT10   = "vat10"  // НДС чека по ставке 10%
	VAT110  = "vat110" // НДС чека по расчетной ставке 10/110

	// Deprecated: since 1 Jan 2019 vat18 will be replaced automatically to vat20 on the side of Tinkoff bank. Use VAT20 instead
	VAT18 = "vat18" // НДС чека по ставке 18%
	VAT20 = "vat20" // НДС чека по ставке 20%

	// Deprecated: since 1 Jan 2019 vat118 will be replaced automatically to vat120 on the side of Tinkoff bank. Use VAT120 instead
	VAT118 = "vat118" // НДС чека по расчетной ставке 18/118
	VAT120 = "vat120" // НДС чека по расчетной ставке 20/120
)

var taxationOptions = []string{
	TaxationOSN,
	TaxationUSNIncome,
	TaxationUSNIncomeOutcome,
	TaxationENVD,
	TaxationESN,
	TaxationPatent,
}

var vatOptions = []string{
	VATNone,
	VAT0,
	VAT10,
	VAT18,
	VAT20,
	VAT110,
	VAT118,
	VAT120,
}

type ReceiptItem struct {
	Name     string // НVAT18аименование товара. Максимальная длина строки – 64 символа
	Price    uint64 // Цена в копейках. *Целочисленное значение не более 10 знаков
	Quantity string // Количество/вес: целая часть не более 8 знаков; дробная часть не более 3 знаков
	Amount   uint64 // Сумма в копейках. Целочисленное значение не более 10 знаков
	Tax      string // Ставка налога
	Ean13    string // Штрих-код
	ShopCode string // Код магазина
}

func (i *ReceiptItem) IsValid() bool {
	if i.Name == "" || i.Price == 0 || i.Quantity == "" || i.Amount == 0 || i.Tax == "" {
		return false
	}

	for _, option := range vatOptions {
		if i.Tax == option {
			return true
		}
	}

	return false
}

type Receipt struct {
	Items    []*ReceiptItem
	Email    string
	Phone    string
	Taxation string
}

func (r *Receipt) IsValid() bool {
	if r.Email == "" && r.Phone == "" {
		return false
	}

	for _, option := range taxationOptions {
		if r.Taxation == option {
			return true
		}
	}

	return false
}
