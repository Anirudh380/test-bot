/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MultiOrderResponse struct {
	Status string `json:"status,omitempty"`
	// Response data for multi order request
	Data []MultiOrderData `json:"data,omitempty"`
	// Error details for multi order request
	Errors []MultiOrderError `json:"errors,omitempty"`
	Summary *MultiOrderSummary `json:"summary,omitempty"`
}
