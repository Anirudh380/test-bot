/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type GetMarketStatusResponse struct {
	Status string `json:"status,omitempty"`
	Data *MarketStatusData `json:"data,omitempty"`
}
