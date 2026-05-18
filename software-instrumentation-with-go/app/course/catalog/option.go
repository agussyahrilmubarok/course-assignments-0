package catalog

import "errors"

var (
	ErrCurrencyMismatch    = errors.New("currency mismatch")
	ErrInsufficientAmount  = errors.New("insufficient amount")
	ErrInvalidCourseStatus = errors.New("invalid course status transition")
	ErrInvalidBatchStatus  = errors.New("invalid batch status transition")
	ErrBatchNotOpen        = errors.New("batch not open for booking")
	ErrNotEnoughSeats      = errors.New("not enough seats available")
	ErrBatchAlreadyStarted = errors.New("batch already started")
	ErrBatchNotFound       = errors.New("batch not found")
	ErrBatchFull           = errors.New("batch full")
)
