/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ColumnInfo struct {
	// Column description
	Description string `json:"description,omitempty"`
	// Whether or not the column can be inserted into
	Insertable bool `json:"insertable,omitempty"`
	// Column name
	Name string `json:"name"`
	// Is column nullable
	Nullable bool `json:"nullable"`

	Type_ *ColumnType `json:"type"`
	// Whether or not the column can be updated
	Updatable bool `json:"updatable,omitempty"`

	ValueGenerated *ColumnValueGenerationStrategy `json:"value_generated,omitempty"`
}
