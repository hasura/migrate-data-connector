/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CustomUpdateColumnOperatorRowUpdate struct {
	// The name of the column in the row
	Column string `json:"column"`

	OperatorName string `json:"operator_name"`

	Type_ string `json:"type"`
	// The value to use with the column operator
	Value ModelMap `json:"value"`

	ValueType string `json:"value_type"`
}
