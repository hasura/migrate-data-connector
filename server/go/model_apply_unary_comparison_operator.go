/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ApplyUnaryComparisonOperator struct {

	Column *ComparisonColumn `json:"column"`

	Operator *UnaryComparisonOperator `json:"operator"`

	Type_ string `json:"type"`
}
