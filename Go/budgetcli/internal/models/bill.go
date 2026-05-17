package models

type PayPeriod string

const (
	PayPeriodA PayPeriod = "A"
	PayPeriodB PayPeriod = "B"
)

type Bill struct {
	BillName 				 string
	DueRecurringDate int
	DueAmount        float64
	DueBalance       *float64
	PayPeriodPaid    PayPeriod
	IsAutoPay        bool
	Notes            string
}
