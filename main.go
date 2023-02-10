package main

import (
	"fmt"

	"log"

	"github.com/shallowravine/Workspace_1/helper"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	a, b := 2, 3
	c := helper.Add(a, b)
	fmt.Println(c)

	ctx, cancel, client, err := client.createClient

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	tD := client.Database("TestDatabase")

	theCollection := tD.Collection("testCollection")

	insertResult, err := theCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"tenant_id", "34e89ac3-9a3a-4a7e-9ea3-3aa3aac52840"},
			{"user", "slathery"},
			{"password", "cheese"},
			{"tags", bson.A{"meat", "yup", "mhmm"}},
		},
		bson.D{
			{"tenant_id", "009498b1-a78a-4d03-a39f-c72e5b3c6351"},
			{"user", "chicken"},
			{"password", "run"},
			{"tags", bson.A{"Old", "CodeButter", "Metal"}},
		},
	})

	if err != nil {
		log.Println("There was an err in trying to migrate the data into the database")

	}
	fmt.Println(insertResult.InsertedIDs)

}
