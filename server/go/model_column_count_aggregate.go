/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ColumnCountAggregate struct {
	// The column to apply the count aggregate function to
	Column string `json:"column"`
	// Whether or not only distinct items should be counted
	Distinct bool `json:"distinct"`

	Type_ string `json:"type"`
}
