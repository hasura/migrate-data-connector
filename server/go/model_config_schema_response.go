/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ConfigSchemaResponse struct {

	ConfigSchema *OpenApiSchema `json:"config_schema"`

	OtherSchemas map[string]OpenApiSchema `json:"other_schemas"`
}