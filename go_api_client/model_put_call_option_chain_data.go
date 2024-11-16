/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PutCallOptionChainData struct {
	InstrumentKey string `json:"instrument_key,omitempty"`
	MarketData *MarketData `json:"market_data,omitempty"`
	OptionGreeks *AnalyticsData `json:"option_greeks,omitempty"`
}
