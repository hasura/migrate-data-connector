/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ApplyBinaryComparisonOperator struct {

	Column *ComparisonColumn `json:"column"`

	Operator *BinaryComparisonOperator `json:"operator"`

	Type_ string `json:"type"`

	Value *ComparisonValue `json:"value"`
}
