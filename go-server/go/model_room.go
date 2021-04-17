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

type Room struct {

	Id int32 `json:"id,omitempty"`

	AccountName string `json:"account_name,omitempty"`

	AccountPassword string `json:"account_password,omitempty"`

	StartingTime time.Time `json:"starting_time,omitempty"`

	EndingTime time.Time `json:"ending_time,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	MaxCount int32 `json:"max_count,omitempty"`

	AdminId int32 `json:"admin_id,omitempty"`

	ServiceId int32 `json:"service_id,omitempty"`

	PlanName string `json:"plan_name,omitempty"`
}