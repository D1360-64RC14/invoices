package routers

import (
	"time"

	"github.com/D1360-64RC14/invoices/web/templates"
	"github.com/gin-gonic/gin"
)

type DefaultRouter struct{}

func NewDefaultRouter() *DefaultRouter {
	return &DefaultRouter{}
}

func (r *DefaultRouter) AttachTo(group *gin.RouterGroup) {
	group.GET("/", r.getMainPage)
	group.GET("/purchase-invoice", r.getPurchaseInvoice)
	group.GET("/detailed-purchase-invoice", r.getDetailedPurchaseInvoice)
}

func (r *DefaultRouter) getMainPage(c *gin.Context) {
	c.File("web/app/index.html")
}

func (r *DefaultRouter) getPurchaseInvoice(c *gin.Context) {
	c.HTML(200, "invoice-view.html", templates.InvoiceView{
		Template: "purchase-invoice",
		Content: templates.PurchaseInvoice{
			Name:   "Diego Garcia",
			Number: 256,
			Date:   time.Now(),
			Items: []templates.PurchaseInvoiceItem{
				{Description: "Par de botas", Ammount: 1, Price: 124.99},
				{Description: "Travesseiro", Ammount: 1, Price: 28.50},
				{Description: "Capa para Travesseiro", Ammount: 1, Price: 32.00},
				{Description: "Cabide", Ammount: 5, Price: 25.03},
			},
			Total:   124.99 + 28.50 + 32.00 + 25.03,
			Message: "Dolore officia nisi mollit exercitation ex duis ad. Dolor aute enim duis non veniam dolore. Nulla qui tempor est consectetur elit tempor minim consequat consectetur.",
		},
	})
}

func (r *DefaultRouter) getDetailedPurchaseInvoice(c *gin.Context) {
	c.HTML(200, "invoice-view.html", templates.InvoiceView{
		Template: "detailed-purchase-invoice",
		Content: templates.DetailedPurchaseInvoice{
			Company: templates.DetailedPurchaseInvoiceCompany{
				Name:    "Company",
				Address: "Av. Lorem Ipsum, 42",
				Phone:   "00 9 9100-0000",
				Email:   "contact@company.com",
			},
			Number: 256,
			Date:   time.Now(),
			Items: []templates.DetailedPurchaseInvoiceItem{
				{Description: "Par de botas", UnitPrice: 124.99, Ammount: 1, Price: 124.99},
				{Description: "Travesseiro", UnitPrice: 28.50, Ammount: 2, Price: 57.00},
				{Description: "Capa para Travesseiro", UnitPrice: 32.00, Ammount: 2, Price: 64.00},
				{Description: "Cabides", UnitPrice: 5.01, Ammount: 5, Price: 25.05},
				{Description: "Vazo de Flor", UnitPrice: 2.50, Ammount: 3, Price: 7.50},
			},
			Discounts:     10.00,
			ShipmentCosts: 25.00,
			Total:         124.99 + 57.00 + 64.00 + 25.05 + 7.50 - 10.00 + 25.00,
		},
	})
}
