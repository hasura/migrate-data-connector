/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RelationshipField struct {

	Query *Query `json:"query"`
	// The name of the relationship to follow for the subquery
	Relationship string `json:"relationship"`

	Type_ string `json:"type"`
}
