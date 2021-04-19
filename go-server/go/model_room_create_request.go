/*
 * Salamander api server
 *
 * Salamander is the backend api server for HermitCrab.
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RoomCreateRequest struct {

	MaxCount int32 `json:"max_count,omitempty"`

	AdminId int32 `json:"admin_id,omitempty"`

	ServiceId int32 `json:"service_id,omitempty"`

	PlanName string `json:"plan_name,omitempty"`
}
