// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ProfitDetailsColumns holds the columns for the "profit_details" table.
	ProfitDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "benefit_date", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// ProfitDetailsTable holds the schema information for the "profit_details" table.
	ProfitDetailsTable = &schema.Table{
		Name:       "profit_details",
		Columns:    ProfitDetailsColumns,
		PrimaryKey: []*schema.Column{ProfitDetailsColumns[0]},
	}
	// ProfitGeneralsColumns holds the columns for the "profit_generals" table.
	ProfitGeneralsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "to_platform", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "to_user", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "transferred_to_platform", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "transferred_to_user", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
	}
	// ProfitGeneralsTable holds the schema information for the "profit_generals" table.
	ProfitGeneralsTable = &schema.Table{
		Name:       "profit_generals",
		Columns:    ProfitGeneralsColumns,
		PrimaryKey: []*schema.Column{ProfitGeneralsColumns[0]},
	}
	// ProfitUnsoldsColumns holds the columns for the "profit_unsolds" table.
	ProfitUnsoldsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "benefit_date", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// ProfitUnsoldsTable holds the schema information for the "profit_unsolds" table.
	ProfitUnsoldsTable = &schema.Table{
		Name:       "profit_unsolds",
		Columns:    ProfitUnsoldsColumns,
		PrimaryKey: []*schema.Column{ProfitUnsoldsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProfitDetailsTable,
		ProfitGeneralsTable,
		ProfitUnsoldsTable,
	}
)

func init() {
}
