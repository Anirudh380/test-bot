/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TokenResponse struct {
	// E-mail address of the user
	Email string `json:"email,omitempty"`
	// Lists the exchanges to which the user has access
	Exchanges []string `json:"exchanges,omitempty"`
	// Lists the products types to which the user has access
	Products []string `json:"products,omitempty"`
	// The broker ID
	Broker string `json:"broker,omitempty"`
	// Uniquely identifies the user
	UserId string `json:"user_id,omitempty"`
	// Name of the user
	UserName string `json:"user_name,omitempty"`
	// Order types enabled for the user
	OrderTypes []string `json:"order_types,omitempty"`
	//   Identifies the user's registered role at the broker. This will be individual for all retail users
	UserType string `json:"user_type,omitempty"`
	//   To depict if the user has given power of attorney for transactions
	Poa bool `json:"poa,omitempty"`
	//   Indicates if DDPI is enabled for trading
	Ddpi bool `json:"ddpi,omitempty"`
	//   Whether the status of account is active or not
	IsActive bool `json:"is_active,omitempty"`
	// The authentication token that is to used with every subsequent API requests
	AccessToken string `json:"access_token,omitempty"`
	// An extended authentication token with a prolonged validity period, intended for specific API requests. Ensure you use this token only with the designated set of APIs.
	ExtendedToken string `json:"extended_token,omitempty"`
}
