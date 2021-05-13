/*
 * Salamander api server
 *
 * Salamander is the backend api server for HermitCrab.
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type Body1 struct {

	StartDate time.Time `json:"start_date,omitempty"`

	PlanInterval int32 `json:"plan_interval,omitempty"`

	PaymentDeadline int32 `json:"payment_deadline,omitempty"`

	IsAddCanlendar bool `json:"is_add_canlendar,omitempty"`
}
