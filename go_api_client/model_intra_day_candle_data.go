/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Get OHLC values for all instruments across various timeframes
type IntraDayCandleData struct {
	Candles [][]interface{} `json:"candles,omitempty"`
}
