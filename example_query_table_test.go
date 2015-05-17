package gorethink

import (
	"fmt"
)

func ExampleTerm_TableCreate() {
	// Setup database
	DB("test").TableDrop("table").Run(sess)

	response, err := DB("test").TableCreate("table").RunWrite(sess)
	if err != nil {
		log.Fatalf("Error creating table: %s", err)
	}

	fmt.Printf("%d table created", response.TablesCreated)

	// Output:
	// 1 table created
}

func ExampleTerm_IndexCreate() {
	// Setup database
	DB("test").TableDrop("table").Run(sess)
	DB("test").TableCreate("table").Run(sess)

	response, err := DB("test").Table("table").IndexCreate("name").RunWrite(sess)
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}

	fmt.Printf("%d index created", response.Created)

	// Output:
	// 1 index created
}

func ExampleTerm_IndexCreate_compound() {
	// Setup database
	DB("test").TableDrop("table").Run(sess)
	DB("test").TableCreate("table").Run(sess)

	response, err := DB("test").Table("table").IndexCreateFunc("full_name", func(row Term) interface{} {
		return []interface{}{row.Field("first_name"), row.Field("last_name")}
	}).RunWrite(sess)
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}

	fmt.Printf("%d index created", response.Created)

	// Output:
	// 1 index created
}
