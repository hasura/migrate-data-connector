/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DeleteMutationOperation struct {
	// The fields to return for the rows affected by this delete operation
	ReturningFields map[string]Field `json:"returning_fields,omitempty"`

	Table *[]string `json:"table"`

	Type_ string `json:"type"`

	Where *Expression `json:"where,omitempty"`
}
