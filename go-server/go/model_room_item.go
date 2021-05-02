/*
 * Salamander api server
 *
 * Salamander is the backend api server for HermitCrab.
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RoomItem struct {

	RoomId int32 `json:"room_id,omitempty"`

	ServiceName string `json:"service_name,omitempty"`

	PlanName string `json:"plan_name,omitempty"`

	Role string `json:"role,omitempty"`

	PaymentFee int32 `json:"payment_fee,omitempty"`

	PaymentStatus *PaymentStatus `json:"payment_status,omitempty"`

	RoomStatus *RoomStatus `json:"room_status,omitempty"`
}
