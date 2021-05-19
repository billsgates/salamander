/*
 * Salamander api server
 *
 * Salamander is the backend api server for HermitCrab.
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RoomInfoResponse struct {

	RoomId int32 `json:"room_id,omitempty"`

	IsPublic bool `json:"is_public,omitempty"`

	Announcement string `json:"announcement,omitempty"`

	MaxCount int32 `json:"max_count,omitempty"`

	RoomStatus *RoomStatus `json:"room_status,omitempty"`

	ServiceName string `json:"service_name,omitempty"`

	PlanName string `json:"plan_name,omitempty"`

	Role string `json:"role,omitempty"`

	PaymentFee int32 `json:"payment_fee,omitempty"`

	RoundInfo *RoundInfo `json:"round_info,omitempty"`

	Admin *Admin `json:"admin,omitempty"`

	Members []Member `json:"members,omitempty"`
}
