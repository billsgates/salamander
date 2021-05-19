/*
 * Salamander api server
 *
 * Salamander is the backend api server for HermitCrab.
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RoundInfo struct {

	PaymentDeadline string `json:"payment_deadline,omitempty"`

	RoundInterval int32 `json:"round_interval,omitempty"`

	StartingTime string `json:"starting_time,omitempty"`

	EndingTime string `json:"ending_time,omitempty"`
}
