/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Asks
type Depth struct {
	// quantity
	Quantity int32 `json:"quantity,omitempty"`
	// price
	Price float64 `json:"price,omitempty"`
	// orders
	Orders int32 `json:"orders,omitempty"`
}