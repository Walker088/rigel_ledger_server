package exchangerate

type ExchangeRateManager struct {
	source string
}

func New() *ExchangeRateManager {
	return &ExchangeRateManager{
		source: "https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies",
	}
}

func (er *ExchangeRateManager) GetExchange(from string, to string) {

}
