/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ColumnInsertSchema struct {
	// The name of the column that this field should be inserted into
	Column string `json:"column"`

	ColumnType *ColumnType `json:"column_type"`
	// Is the column nullable
	Nullable bool `json:"nullable"`

	Type_ string `json:"type"`

	ValueGenerated *ColumnValueGenerationStrategy `json:"value_generated,omitempty"`
}
