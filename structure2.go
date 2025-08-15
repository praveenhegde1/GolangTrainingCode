package main

import (
	"fmt"
)

// Product represents a manufactured product
type Product struct {
	ID       string
	Name     string
	Quantity int
	Weight   float64
}

// Inventory represents the current stock levels
type Inventory struct {
	Products map[string]Product
}

// ProductionOrder represents a production order
type ProductionOrder struct {
	ProductID string
	Quantity  int
}

// ManufacturingSystem manages inventory and production
type ManufacturingSystem struct {
	Inventory        Inventory
	ProductionOrders []ProductionOrder
}

func (ms *ManufacturingSystem) AddProduct(product Product) {
	ms.Inventory.Products[] = product
}

func (ms *ManufacturingSystem) PlaceOrder(order ProductionOrder) {
	product, ok := ms.Inventory.Products[order.ProductID]
	if !ok {
		fmt.Println("Product not found")
		return
	}
	if product.Quantity < order.Quantity {
		fmt.Println("Insufficient inventory")
		return
	}
	product.Quantity -= order.Quantity
	ms.ProductionOrders = append(ms.ProductionOrders, order)
}

func main() {
	ms := &ManufacturingSystem{
		Inventory: Inventory{
			Products: map[string]Product{},
		},
	}

	// Add products to inventory
	ms.AddProduct(Product{ID: "P1", Name: "Widget", Quantity: 100, Weight: 0.5})
	ms.AddProduct(Product{ID: "P2", Name: "Gadget", Quantity: 50, Weight: 1.0})

	// Place production orders
	ms.PlaceOrder(ProductionOrder{ProductID: "P1", Quantity: 20})
	ms.PlaceOrder(ProductionOrder{ProductID: "P2", Quantity: 30})

	fmt.Println("Inventory:")
	for _, product := range ms.Inventory.Products {
		fmt.Println(product)
	}

	fmt.Println("Production Orders:")
	for _, order := range ms.ProductionOrders {
		fmt.Println(order)
	}
}
