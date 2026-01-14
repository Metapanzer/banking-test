package validator

import (
	"regexp"
	"slices"
	"task-8/model"
)

func ValidateResponse(product model.Product) (bool, []string) {
	isValid := true
	var errorMessages []string
	skuValid, message := isValidSKU(product.Sku)
	if !skuValid {
		isValid = false
		errorMessages = append(errorMessages, message...)
	}
	nameValid, message := isValidProductName(product.ProductName)
	if !nameValid {
		isValid = false
		errorMessages = append(errorMessages, message...)
	}
	quantityValid, message := isQuantityInStock(product.QuantityInStock)
	if !quantityValid {
		isValid = false
		errorMessages = append(errorMessages, message...)
	}
	priceValid, message := isValidPrice(product.Price)
	if !priceValid {
		isValid = false
		errorMessages = append(errorMessages, message...)
	}
	categoryValid, message := isValidCategory(product.Category)
	if !categoryValid {
		isValid = false
		errorMessages = append(errorMessages, message...)
	}

	return isValid, errorMessages
}

func isValidSKU(sku string) (bool, []string) {
	isValid := true
	var message []string
	if sku == "" || len(sku) == 0 {
		message = append(message, "The sku is a mandatory field")
		isValid = false
	}
	if len(sku) < 12 || len(sku) > 12 {
		message = append(message, "The sku must be in the format SKU-XXXXXXXX")
		isValid = false
	}

	re := regexp.MustCompile(`^SKU-\d{8}$`)
	if !re.MatchString(sku) {
		message = append(message, "The sku must be in the format SKU-XXXXXXXX")
		isValid = false
	}
	return isValid, message
}

func isValidProductName(name string) (bool, []string) {
	isValid := true
	var message []string
	if name == "" || len(name) == 0 {
		message = append(message, "The productName is a mandatory field")
		isValid = false
	}
	return isValid, message
}

func isQuantityInStock(quantity int) (bool, []string) {
	isValid := true
	var message []string
	if quantity < 0 {
		message = append(message, "The quantityInStock cannot be negative")
		isValid = false
	}
	return isValid, message
}

func isValidPrice(price float64) (bool, []string) {
	isValid := true
	var message []string
	if price <= 0 {
		message = append(message, "The price must be greater than zero")
		isValid = false
	}
	return isValid, message
}

func isValidCategory(category string) (bool, []string) {
	allowedCategories := []string{"Electronics", "Books", "Apparel", "Home Goods"}
	isValid := true
	var message []string
	if category == "" || len(category) == 0 {
		message = append(message, "The category is a mandatory field")
		isValid = false
	}

	if !slices.Contains(allowedCategories, category) {
		message = append(message, "Invalid product category")
		isValid = false
	}
	return isValid, message
}
