// Package classification of Product API
//
// Documentation for Product API
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
// swagger:meta

package handlers

import "github.com/codepnw/microservices/api/data"

// swagger:response productsResponse
type productsResponseWrapper struct {
	// in:body
	Body []data.Product
}

// swagger:parameters deleteProduct
type productIdParameterWrapper struct {
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:response noContent
type productsNoContent struct{}
