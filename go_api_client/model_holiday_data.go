/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// Response data for holiday list
type HolidayData struct {
	Date time.Time `json:"date,omitempty"`
	Description string `json:"description,omitempty"`
	HolidayType string `json:"holiday_type,omitempty"`
	ClosedExchanges []string `json:"closed_exchanges,omitempty"`
	OpenExchanges []ExchangeTimingData `json:"open_exchanges,omitempty"`
}
