/*
 * Salamander api server
 *
 * Salamander is the backend api server for HermitCrab.
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Body3 struct {

	UserId int32 `json:"userId,omitempty"`

	RoomId int32 `json:"roomId,omitempty"`

	Status *PaymentStatus `json:"status,omitempty"`
}
