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
	group.GET("/purchase-invoice", r.postPurchaseInvoice)
	group.GET("/detailed-purchase-invoice", r.postDetailedPurchaseInvoice)
}

func (r *DefaultRouter) getMainPage(c *gin.Context) {
	c.File("web/app/index.html")
}

func (r *DefaultRouter) postPurchaseInvoice(c *gin.Context) {
	c.HTML(200, "purchase-invoice.html", templates.PurchaseInvoice{
		Name:   "Diego Garcia",
		Number: 256,
		Date:   time.Now(),
		Items: []templates.PurchaseInvoiceItem{
			{Description: "Par de botas", Price: 124.99},
			{Description: "Travesseiro", Price: 28.50},
			{Description: "Capa para Travesseiro", Price: 32.00},
			{Description: "5 Cabides", Price: 25.03},
		},
		Total: 124.99 + 28.50 + 32.00 + 25.03,
	})
}

func (r *DefaultRouter) postDetailedPurchaseInvoice(c *gin.Context) {
	c.HTML(200, "detailed-purchase-invoice.html", templates.DetailedPurchaseInvoice{
		Name:   "Diego Garcia",
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
	})
}