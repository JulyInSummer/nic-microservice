// Package classification for Product API
//
// Documentation for product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger: meta
package handlers

// Generic error message returned as a string
// swagger: response errorResponse
type errorResponse struct {
	// Description of the error
	// in: body
	Body genericError
}

// Validation errors defined as an array of strings
// swagger: response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body validationError
}

// swagger:parameters deleteProduct
type productIDParametersWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:response noContent
type productsNoContent struct {
}

type genericError struct {
	Message string `json:"message"`
}

type validationError struct {
	Messages []string `json:"messages"`
}
