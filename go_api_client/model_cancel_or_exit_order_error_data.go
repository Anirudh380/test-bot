/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Error data for cancel or exit order request
type CancelOrExitOrderErrorData struct {
	// Unique code for the error state
	ErrorCode string `json:"errorCode,omitempty"`
	// Verbose message for the error state
	Message string `json:"message,omitempty"`
	// Path to property failing validation
	PropertyPath string `json:"propertyPath,omitempty"`
	// Invalid value for the property failing validation
	InvalidValue *interface{} `json:"invalidValue,omitempty"`
	// Key of instrument
	InstrumentKey string `json:"instrument_key,omitempty"`
	// Reference order ID
	OrderId string `json:"order_id,omitempty"`
}
