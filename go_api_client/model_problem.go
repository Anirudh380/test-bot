/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Problem struct {
	// Unique code for the error state
	ErrorCode string `json:"errorCode,omitempty"`
	// Verbose message for the error state
	Message string `json:"message,omitempty"`
	// Path to property failing validation
	PropertyPath string `json:"propertyPath,omitempty"`
	// Invalid value for the property failing validation
	InvalidValue *interface{} `json:"invalidValue,omitempty"`
}
