package postgres

import (
	"crystal-cli/schema"
	"fmt"
)

func GenerateSQL(schema *schema.Schema) {
	var sql string = ``
	for table, cols := range schema.Models {
		t := fmt.Sprintf("CREATE TABLE IF NOT EXIST %s (", table)
		for col, colData := range cols {
			t += fmt.Sprintf("%s %s,", col, colData.Fields.Type)
		}
		sql += t + ")\n"
	}
}
