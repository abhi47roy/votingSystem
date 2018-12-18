package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type BlacklistMedia struct {
	MediaList    []Media `json:"mediaList"`
	UpdateTime   string  `json:"updateTime"`
	Organisation string  `json:"organisation"`
}
type Media struct {
	MediaNo string `json:"mediaNo"`
	Status  string `json:"status"`
	Reason  string `json:"reason"`
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
	if function == "pushBlacklistedCards" { //create a new marble
		return t.pushBlacklistedCards(stub, args)
	} else if function == "getBlacklistedCards" {
		return t.getBlacklistedCards(stub)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initMarble - create a new marble, store into chaincode state
// ============================================================
func (t *SimpleChaincode) pushBlacklistedCards(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments.")
	}

	// ==== Input sanitation ====
	var mediaListJson BlacklistMedia

	err = json.Unmarshal([]byte(args[0]), &mediaListJson)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", mediaListJson)
	}

	mediaListJson.Organisation = strings.ToLower(args[1])

	mediaListJSONasBytes, err := json.Marshal(mediaListJson)
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState("blacklistMedia", mediaListJSONasBytes)

	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init transaction")
	return shim.Success(nil)
}

// query all transactions
func (t *SimpleChaincode) getBlacklistedCards(stub shim.ChaincodeStubInterface) pb.Response {

	queryResults, err := stub.GetState("blacklistMedia")
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}
