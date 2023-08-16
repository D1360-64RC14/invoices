package templates

import "time"

type PurchaseInvoice struct {
	Name   string
	Number float32
	Date   time.Time
	Items  []PurchaseInvoiceItem
	Total  float32
}

type PurchaseInvoiceItem struct {
	Description string
	Ammount     int
	Price       float32
}
