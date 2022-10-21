package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"igbo-go/client"
	api "igbo-go/grpc"
)

const defaultURL = "localhost:1234"

func main() {

	igboDBClient := client.NewIgboDbClient(defaultURL)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for {
		attribute := new(api.Attribute)
		attribute.Name = "Attribute1"
		attribute.Type = api.AttributeType_STRING
		attribute.Value = "Attr value 1"

		oKey := new(api.ObjectKey)
		oKey.Type = "MyObjectType"
		oKey.Id = "bfe056c8-41c9-11ed-b878-0242ac120003"
		object := api.Object{
			Key:        oKey,
			Attributes: []*api.Attribute{attribute},
		}
		aObjects := []*api.Object{&object}
		objects := api.Objects{
			Items: aObjects,
		}

		requestId := new(api.RequestId)
		requestId.Type = api.OperationType_CREATE
		requestId.Id = uuid.NewString()

		operationRequest := new(api.OperationRequest)
		operationRequest.RequestId = requestId
		objectKeysRequest2 := new(api.OperationRequest_Objects)
		objectKeysRequest2.Objects = &objects
		operationRequest.Payload = objectKeysRequest2

		fmt.Println("Sending CREATE operation.")
		err := igboDBClient.OperationsStream(*operationRequest)
		fmt.Println("Sleep for 3 sec.")
		time.Sleep(3 * time.Second)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}

		oKey = new(api.ObjectKey)
		oKey.Type = "MyObjectType"
		oKey.Id = "bfe056c8-41c9-11ed-b878-0242ac120003"
		keys := []*api.ObjectKey{oKey}

		objectKeys := new(api.ObjectKeys)
		objectKeys.Keys = keys

		requestId = new(api.RequestId)
		requestId.Type = api.OperationType_READ
		requestId.Id = uuid.NewString()

		operationRequest = new(api.OperationRequest)
		operationRequest.RequestId = requestId
		objectKeysRequest := new(api.OperationRequest_ObjectKeys)
		objectKeysRequest.ObjectKeys = objectKeys
		operationRequest.Payload = objectKeysRequest

		fmt.Println("Sending READ operation.")
		err = igboDBClient.OperationsStream(*operationRequest)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Sleep for 3 sec.")
		time.Sleep(3 * time.Second)

		oKey = new(api.ObjectKey)
		oKey.Type = "MyObjectType"
		oKey.Id = "bfe056c8-41c9-11ed-b878-0242ac120003"
		keys = []*api.ObjectKey{oKey}

		objectKeys = new(api.ObjectKeys)
		objectKeys.Keys = keys

		requestId = new(api.RequestId)
		requestId.Type = api.OperationType_DELETE
		requestId.Id = uuid.NewString()

		operationRequest = new(api.OperationRequest)
		operationRequest.RequestId = requestId
		objectKeysRequest = new(api.OperationRequest_ObjectKeys)
		objectKeysRequest.ObjectKeys = objectKeys
		operationRequest.Payload = objectKeysRequest

		fmt.Println("Sending DELETE operation.")
		err = igboDBClient.OperationsStream(*operationRequest)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Sleep for 3 sec.")
		time.Sleep(3 * time.Second)
	}

}
