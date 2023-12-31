/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TableInfo struct {
	// The columns of the table
	Columns []ColumnInfo `json:"columns"`
	// Whether or not existing rows can be deleted in the table
	Deletable bool `json:"deletable,omitempty"`
	// Description of the table
	Description string `json:"description,omitempty"`
	// Foreign key constraints
	ForeignKeys map[string]Constraint `json:"foreign_keys,omitempty"`
	// Whether or not new rows can be inserted into the table
	Insertable bool `json:"insertable,omitempty"`

	Name *[]string `json:"name"`
	// The primary key of the table
	PrimaryKey []string `json:"primary_key,omitempty"`

	Type_ *TableType `json:"type,omitempty"`
	// Whether or not existing rows can be updated in the table
	Updatable bool `json:"updatable,omitempty"`
}
