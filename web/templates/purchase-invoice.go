package templates

import "time"

type PurchaseInvoice struct {
	Name    string
	Number  float32
	Date    time.Time
	Items   []PurchaseInvoiceItem
	Total   float32
	Message string
}

type PurchaseInvoiceItem struct {
	Description string
	Ammount     int
	Price       float32
}
