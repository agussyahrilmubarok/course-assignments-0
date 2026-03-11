package catalog

type Money struct {
	Amount   int64
	Currency string
}

func NewMoney(amount int64, currency string) Money {
	return Money{
		Amount:   amount,
		Currency: currency,
	}
}

func (m Money) Add(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, ErrCurrencyMismatch
	}

	return Money{
		Amount:   m.Amount + other.Amount,
		Currency: m.Currency,
	}, nil
}

func (m Money) Subtract(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, ErrCurrencyMismatch
	}

	if m.Amount < other.Amount {
		return Money{}, ErrInsufficientAmount
	}

	return Money{
		Amount:   m.Amount - other.Amount,
		Currency: m.Currency,
	}, nil
}

func (m Money) Multiply(qty int64) Money {
	return Money{
		Amount:   m.Amount * qty,
		Currency: m.Currency,
	}
}
