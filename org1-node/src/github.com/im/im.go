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
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Identity struct {
	ID     string `json:"id"`
	Org    string `json:"name"`
	UserID string `json:"user_id"`
}

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
	// "Org", "name"
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

 	org := args[0]
	userId := args[1]
	id := strings.Join([]string{org, userId}, " ")
	// ==== Check if user already exists ====
	userAsBytes, err := stub.GetState(id)
	if err != nil {
		return shim.Error("Error to get world state: " + err.Error())
	} else if userAsBytes != nil {
		return shim.Error("This user already exists: " + id)
	}

 	// ==== Create user object and marshal to JSON ====
	user := Identity{ID: id, Org: org, UserID: userId}
	userJSONasBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save identity to state ===
	err = stub.PutState(id, userJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

 	// ==== Identity saved. Return success ====
	return shim.Success(nil)
}

 func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

 	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting name of org and userId to query")
	}

 	org := args[0]
	userId := args[1]
	id := strings.Join([]string{org, userId}, " ")
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
