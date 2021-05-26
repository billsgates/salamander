/*
 * Salamander api server
 *
 * Salamander is the backend api server for Bills Gate.
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type User struct {

	Id int32 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Email string `json:"email,omitempty"`

	Phone string `json:"phone,omitempty"`

	Rating float64 `json:"rating,omitempty"`

	RatingCount int32 `json:"rating_count,omitempty"`
}
