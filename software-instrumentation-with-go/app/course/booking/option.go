package booking

import "errors"

var (
	ErrBookingNotFound      = errors.New("booking not found")
	ErrBookingExpired       = errors.New("booking expired")
	ErrBookingAlreadyPaid   = errors.New("booking already paid")
	ErrBookingCanceled      = errors.New("booking canceled")
	ErrInvalidBookingStatus = errors.New("invalid booking status")
	ErrBookingCodeExists    = errors.New("booking code already exists")
	ErrInvalidCustomerName  = errors.New("invalid customer name")
)
