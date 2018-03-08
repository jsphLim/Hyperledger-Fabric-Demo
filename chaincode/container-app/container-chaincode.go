
package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)


type SmartContract struct {
}


type Container struct {
	X string `json:"x"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}


func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	
	function, args := APIstub.GetFunctionAndParameters()

	if function == "set" {
		return s.set(APIstub, args)
	} else if function == "get" {
		return s.get(APIstub,args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}


func (s *SmartContract) get(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
		
	containerAsBytes, _ := APIstub.GetState(args[0])
	if containerAsBytes == nil {
		return shim.Error("Could not locate container")
	}
	return shim.Success(containerAsBytes)
}



func (s *SmartContract) set(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	containerAsBytes, _ := json.Marshal(args[1])

	err:=APIstub.PutState(args[0],containerAsBytes)
	if err != nil{
		return shim.Error("wirteIn error")
	}

	return shim.Success(nil)
	
}

 
func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
