/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MutationCapabilities struct {

	AtomicitySupportLevel *AtomicitySupportLevel `json:"atomicity_support_level,omitempty"`

	Delete *Object `json:"delete,omitempty"`

	Insert *InsertCapabilities `json:"insert,omitempty"`

	Returning *Object `json:"returning,omitempty"`

	Update *Object `json:"update,omitempty"`
}
