/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type FunctionRelationships struct {
	// A map of relationships from the source table to target tables. The key of the map is the relationship name
	Relationships map[string]Relationship `json:"relationships"`

	SourceFunction *[]string `json:"source_function"`

	Type_ string `json:"type"`
}
