package templates

import "time"

type DetailedPurchaseInvoice struct {
	Name          string
	Number        float32
	Date          time.Time
	Items         []DetailedPurchaseInvoiceItem
	Discounts     float32
	ShipmentCosts float32
	Total         float32
}

type DetailedPurchaseInvoiceItem struct {
	Description string
	UnitPrice   float32
	Ammount     int
	Price       float32
}
