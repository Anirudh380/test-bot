/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ExchangeTimingData struct {
	Exchange string `json:"exchange,omitempty"`
	StartTime int64 `json:"start_time,omitempty"`
	EndTime int64 `json:"end_time,omitempty"`
}
