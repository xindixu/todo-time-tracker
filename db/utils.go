package db

import (
	"reflect"

	"github.com/doug-martin/goqu/v9"
)

// ModelToRecord converts a model struct to a goqu.Record using reflection
// It reads the `db` tags from struct fields to map them to database columns
func ModelToRecord(model interface{}) goqu.Record {
	record := goqu.Record{}

	v := reflect.ValueOf(model)
	t := reflect.TypeOf(model)

	// Handle pointer types
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	// Iterate through struct fields
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Get the db tag
		dbTag := fieldType.Tag.Get("db")
		if dbTag == "" || dbTag == "-" {
			continue // Skip fields without db tags or with "-"
		}

		// Skip ID field for inserts (usually auto-generated)
		if dbTag == "id" {
			continue
		}

		// Get the field value
		if field.IsValid() && field.CanInterface() {
			record[dbTag] = field.Interface()
		}
	}

	return record
}
