/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Response data for market live status
type MarketStatusData struct {
	Exchange string `json:"exchange,omitempty"`
	Status string `json:"status,omitempty"`
	LastUpdated int64 `json:"last_updated,omitempty"`
}
