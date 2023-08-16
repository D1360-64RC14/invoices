package templates

import "time"

type DetailedPurchaseInvoice struct {
	Company       DetailedPurchaseInvoiceCompany
	Number        float32
	Date          time.Time
	Items         []DetailedPurchaseInvoiceItem
	Discounts     float32
	ShipmentCosts float32
	Total         float32
}

type DetailedPurchaseInvoiceCompany struct {
	Name    string
	Address string
	Phone   string
	Email   string
}

type DetailedPurchaseInvoiceItem struct {
	Description string
	UnitPrice   float32
	Ammount     int
	Price       float32
}
