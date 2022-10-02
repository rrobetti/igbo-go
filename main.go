package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"igbo-go/client"
	api "igbo-go/grpc"
)

const defaultURL = "localhost:8080"

func main() {
	create := flag.Bool("create", false, "Create object")
	update := flag.Bool("update", false, "Update object")
	delete := flag.Bool("delete", false, "Delete object")
	retrieve := flag.Bool("retrieve", false, "Retrieve object")
	list := flag.Bool("list", false, "List activities")

	flag.Parse()

	igboDBClient := client.NewIgboDbClient(defaultURL)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch {
	case *retrieve:
		aIds := []string{"bfe056c8-41c9-11ed-b878-0242ac120003"}
		ids := api.Ids{Values: aIds}
		objects, err := igboDBClient.Retrieve(ctx, &ids)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		json, _ := json.Marshal(objects)
		fmt.Println(string(json))
	case *create:
		if len(os.Args) != 3 {
			fmt.Fprintln(os.Stderr, `Usage: --create "message"`)
			os.Exit(1)
		}

		attribute := api.Attribute{
			Name:  "Attribute1",
			Type:  api.AttributeType_STRING,
			Value: "Attr value 1",
		}
		object := api.Object{
			Id:         "bfe056c8-41c9-11ed-b878-0242ac120003",
			Type:       "MyObjectType",
			Attributes: []*api.Attribute{&attribute},
		}
		aObjects := []*api.Object{&object}
		objects := api.Objects{
			Items: aObjects,
		}

		id, err := igboDBClient.Create(ctx, &objects)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		json, _ := json.Marshal(objects)
		fmt.Printf("Added: %s as %d\n", string(json), id)
	case *update:
		aIds := []string{"bfe056c8-41c9-11ed-b878-0242ac120003"}
		ids := api.Ids{Values: aIds}
		objects, err := igboDBClient.Retrieve(ctx, &ids)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		objects.Items[0].Attributes[0].Value = "Attribute Updated"
		json, _ := json.Marshal(objects)
		fmt.Printf("Objects updated: %s \n", string(json))
		results, err := igboDBClient.Update(ctx, objects)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		for _, result := range results.Results {
			fmt.Printf("Update result: %v %v \n", result.Type, result.Message)
		}
	case *delete:
		aIds := []string{"bfe056c8-41c9-11ed-b878-0242ac120003"}
		ids := api.Ids{Values: aIds}
		results, err := igboDBClient.Delete(ctx, &ids)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		for _, result := range results.Results {
			fmt.Printf("Deleted result: %v %v \n", result.Type, result.Message)
		}
	case *list:
	default:
		flag.Usage()
		os.Exit(1)
	}
}
