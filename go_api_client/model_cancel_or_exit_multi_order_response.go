/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CancelOrExitMultiOrderResponse struct {
	Status string `json:"status,omitempty"`
	Data *CancelOrExitMultiOrderData `json:"data,omitempty"`
	// Error data for cancel or exit order request
	Errors []CancelOrExitOrderErrorData `json:"errors,omitempty"`
	Summary *BatchExecutionSummary `json:"summary,omitempty"`
}
