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

type RoomInfoResponse struct {

	RoomId int32 `json:"room_id,omitempty"`

	IsPublic bool `json:"is_public,omitempty"`

	Announcement string `json:"announcement,omitempty"`

	MaxCount int32 `json:"max_count,omitempty"`

	PaymentPeriod int32 `json:"payment_period,omitempty"`

	RoomStatus *RoomStatus `json:"room_status,omitempty"`

	StartingTime time.Time `json:"starting_time,omitempty"`

	EndingTime time.Time `json:"ending_time,omitempty"`

	ServiceName string `json:"service_name,omitempty"`

	PlanName string `json:"plan_name,omitempty"`

	Role string `json:"role,omitempty"`

	PaymentFee int32 `json:"payment_fee,omitempty"`

	AdminName string `json:"admin_name,omitempty"`

	AdminRating float64 `json:"admin_rating,omitempty"`

	AdminEmail string `json:"admin_email,omitempty"`

	AdminPhone string `json:"admin_phone,omitempty"`

	Members []Member `json:"members,omitempty"`
}
