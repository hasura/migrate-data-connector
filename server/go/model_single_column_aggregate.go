/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SingleColumnAggregate struct {
	// The column to apply the aggregation function to
	Column string `json:"column"`

	Function string `json:"function"`

	ResultType string `json:"result_type"`

	Type_ string `json:"type"`
}
