/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CapabilitiesResponse struct {

	Capabilities *Capabilities `json:"capabilities"`

	ConfigSchemas *ConfigSchemaResponse `json:"config_schemas"`

	DisplayName string `json:"display_name,omitempty"`

	ReleaseName string `json:"release_name,omitempty"`
}
