/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("hello_world", []byte(args[0]))
	if err != nil {
			return nil, err
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
			return t.Init(stub, "init", args)
	} else if function == "write" {
			return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" {                            //read a variable
			return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)


	return nil, errors.New("Received unknown function query: " + function)
}

//write function

func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var fileName, fileNameValue, signiture, signitureValue string
    var err error
    fmt.Println("running write()")

    if len(args) != 4 {
        return nil, errors.New("Incorrect number of arguments. Expecting 4. name of the variable and value to set")
    }

    fileName = args[0]                            //rename for fun
    fileNameValue = args[1]
		signiture = args[2]                            //rename for fun
		signitureValue = args[3]
		fmt.Println("writing file name...")
    err = stub.PutState(fileName, []byte(fileNameValue))  //write the variable into the chaincode state
		fmt.Println("writing digital signiture ...")
    err = stub.PutState(signiture, []byte(signitureValue))  //write the variable into the chaincode state
    if err != nil {
        return nil, err
    }
    return nil, nil
}

//read function.

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var fileName, fileNameValue, signiture, signitureValue, results, jsonResp string
    var err error

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
    }

    fileName = args[0]
    valAsbytes1, err := stub.GetState(fileName)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + fileName + "\"}"
        return nil, errors.New(jsonResp)
    }
		fileNameValue = string(valAsbytes1)

		signiture = args[1]
    valAsbytes2, err := stub.GetState(signiture)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + signiture + "\"}"
        return nil, errors.New(jsonResp)
    }
		signitureValue = string(valAsbytes2)

		results = "FileName: " + fileNameValue + " Digital signiture: " + signitureValue
    //return valAsbytes, nil
		return []byte(results), nil
}
