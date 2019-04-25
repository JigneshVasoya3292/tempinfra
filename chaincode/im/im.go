/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	/* "bytes"
	"encoding/gob" */
)

type Identity struct {
	ID     string `json:"id"`
	Data    string `json:"data"`
}

type AllIdentities struct {
	items     []string `json:"items"`
}

var ALLIDENTITIES_KEY = "allidentities"

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {	
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "registerUser" { //create a new user
		return t.registerUser(stub, args)
	} else if function == "query" {
		return t.query(stub, args) // to get value for key
	} else if function == "queryAll" {
		return t.queryAll(stub, args) // to get value for key
	} else {
		fmt.Println("invoke did not find func: " + function) //error
		return shim.Error("Received unknown function invocation")
	}

}

// ============================================================
// registerUser - create a new user, store into chaincode state
// ============================================================
func (t *SimpleChaincode) registerUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0       1
	// "Id", "Data"
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("Organisation Name must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("user Id argument must be a non-empty string")
	}

	id := args[0]
	data := args[1]

	// ==== Check if user already exists ====
	userAsBytes, err := stub.GetState(id)
	if err != nil {
		return shim.Error("Error to get world state: " + err.Error())
	} else if userAsBytes != nil {
		return shim.Error("This user already exists: " + id)
	}

	// ==== Create user object and marshal to JSON ====
	user := Identity{ID: id, Data: data}
	userJSONasBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save identity to state ===
	err = stub.PutState(id, userJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	allIdentitiesAsbytes, err := stub.GetState(ALLIDENTITIES_KEY)
	allIdentitiesToUpdate := []string{}
	if err != nil {
		fmt.Println("allidentities key not found")
		allIdentitiesToUpdate = append(allIdentitiesToUpdate, data)
	} else {
		errb := json.Unmarshal(allIdentitiesAsbytes, &allIdentitiesToUpdate) //unmarshal it aka JSON.parse()
		if errb != nil {
			fmt.Println("allidentities key not JSON parsable")
			allIdentitiesToUpdate = append(allIdentitiesToUpdate, data)
		} else {
			fmt.Println("allidentities length before: ", len(allIdentitiesToUpdate))
			allIdentitiesToUpdate = append(allIdentitiesToUpdate, data)
			fmt.Println("allidentities length after: ", len(allIdentitiesToUpdate))
		}		
	}
	
	identitiesJSONBytes, err := json.Marshal(allIdentitiesToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save identity to state ===
	err = stub.PutState(ALLIDENTITIES_KEY, identitiesJSONBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Identity saved. Return success ====
	return shim.Success(nil)
}

func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting id to query")
	}

	id := args[0]
	// Get the state from the ledger
	Avalbytes, err := stub.GetState(id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil Value for " + id + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(Avalbytes)
}

func (t *SimpleChaincode) queryAll(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	/* fmt.Println("queryAll is called ");
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(allIdentities)
	bs := buf.Bytes() */
	
	allIdentitiesAsbytes, err := stub.GetState(ALLIDENTITIES_KEY)
	if err != nil {
		shim.Error("Error to get allidentities:" + err.Error())
	}

	return shim.Success(allIdentitiesAsbytes)
}
