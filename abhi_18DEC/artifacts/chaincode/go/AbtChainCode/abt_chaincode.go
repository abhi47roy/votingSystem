package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)
var transactionReqNo int = 00000
// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type transaction struct {
	TransactionDate string `json:"transactionDate"`
	TransactionType string `json:"transactionType"`
	MediaType       string `json:"mediaType"`
	MediaNo         string `json:"mediaNo"`
	EquipmentType   string `json:"equipmentType"`
	EquipmentID     string `json:"equipmentId"`
	StopID          string `json:"stopId"`
	RouteID         string `json:"routeId"`
	CompanyID       string `json:"companyId"`
	OperatorID      string `json:"operatorId"`
	VehicleNo       string `json:"vehicleNo"`
	ShiftID         string `json:"shiftId"`
	TripID          string `json:"tripId"`
	SubShiftID      string `json:"subShiftId"`
	Organisation    string `json:"organisation"`
	MatchStatus     string `json:"matchStatus"`
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
	fmt.Println("invoke is running " + args[0])

	// Handle different functions
	if args[0] == "initTransaction" { //create a new marble
		return t.initTransaction(stub, args)
	}else if args[0] == "reconcileTransactions" {
		return t.reconcileTransactions(stub, args)
	}else if function == "query" {
		switch args[0] {
		case "queryTransactionsByOrganisation":
			fmt.Println("queryTransactionsByOrganisation is running ")
			return t.queryTransactionsByOrganisation(stub, args)

		case "getTransactionByRange":
			fmt.Println("getTransactionByRange is running ")
			return t.getTransactionByRange(stub, args)
		
		case "queryTransactions":
			fmt.Println("queryTransactions is running ")
			return t.queryTransactions(stub, args)
		
		case "getAllTransactions":
			fmt.Println("getAllTransactions is running ")
			return t.getAllTransactions(stub, args)			
		
		default:
			fmt.Println("Invalid function")
			return shim.Error("Invalid function")
		}
	} 
    fmt.Println("invoke did not find func: " + args[0]) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initTransaction - create a new transaction, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initTransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
     // ==== Input sanitation ====
	var transactionJSON transaction

	err = json.Unmarshal([]byte(args[1]), &transactionJSON)
	if err != nil {
		fmt.Println("Error while unmarhsilling")
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", transactionJSON)
	}

	transactionJSON.Organisation = strings.ToLower(args[2])

	transactionJSON.MatchStatus = "0001"
    fmt.Println("Before Marshilling")
    fmt.Println(transactionJSON)
	transactionJSONasBytes, err := json.Marshal(transactionJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	var indexName string
	if transactionJSON.Organisation == "abt" {
		indexName = "transaction~abt"
	} else if transactionJSON.Organisation == "bank" {
		indexName = "transaction~bank"
	}

	fmt.Println(transactionJSON)
    transactionReqNo = transactionReqNo + 1
	transactionIndexKey, err := stub.CreateCompositeKey(indexName,
		[]string{
			transactionJSON.TransactionDate,
			transactionJSON.TransactionType,
			transactionJSON.MediaType,
			transactionJSON.EquipmentID,
			transactionJSON.EquipmentType,
			transactionJSON.Organisation})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("After the Key Line")

	fmt.Println(transactionIndexKey)
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	//value := []byte{0x00}
	stub.PutState(transactionIndexKey, transactionJSONasBytes)

	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init transaction")
	return shim.Success(nil)
}

// ===========================================================================================
// getTransactionByRange performs a range query based on the start and end keys provided.

// Read-only function results are not typically submitted to ordering. If the read-only
// results are submitted to ordering, or if the query is used in an update transaction
// and submitted to ordering, then the committing peers will re-execute to guarantee that
// result sets are stable between endorsement time and commit time. The transaction is
// invalidated by the committing peers if the result set has changed between endorsement
// time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
// ===========================================================================================
func (t *SimpleChaincode) getTransactionByRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments.")
	}

	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getMarblesByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


func (t *SimpleChaincode) genrateApplicationNumber(indexName string,TransactionType string,Organisation string,TransactionDate string) string {
	//Increse the count of the application number
	requestNumber := indexName + "/" + TransactionType + "/" + Organisation + "/" + TransactionDate
	return requestNumber
}



// ===== Example: Parameterized rich query =================================================
// queryTransactionsByOrganisation queries for transactions based on a passed in organization.
// Only available on state databases that support rich query (e.g. CouchDB)
// =========================================================================================
func (t *SimpleChaincode) queryTransactionsByOrganisation(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("****** Hello World called")

	fmt.Println(args[1])
    
    organisation := strings.ToLower(args[1])

	queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"organisation\":{\"$eq\": \"%s\"}}]}}", organisation )

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// query all transactions
func (t *SimpleChaincode) getAllTransactions(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	startDate := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"transactionDate\": {\"$gte\": \"" + startDate + "\"}}}")

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// queryTransactions uses a query string to perform a query for transactions.
// Only available on state databases that support rich query (e.g. CouchDB)

func (t *SimpleChaincode) queryTransactions(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// if len(args) < 1 {
	// 	return shim.Error("Incorrect number of arguments. Expecting 1")
	// }

	queryString := args[1]

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	for resultsIterator.HasNext() {

		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")

		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// reconcileTransactions to match data pushed by bank and abt organizations
func (t *SimpleChaincode) reconcileTransactions(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// query transaction which are yet to be matched

	startDate := args[len(args)-1]
	fmt.Println(startDate)
	queryAbtString := fmt.Sprintf("{\"selector\": {\"$and\": [{\"matchStatus\": \"0001\"},{\"organisation\": \"abt\"},{\"transactionDate\": {\"$gte\": \"" + startDate + "\"}}]}}")
	queryBankString := fmt.Sprintf("{\"selector\": {\"$and\": [{\"matchStatus\": \"0001\"},{\"organisation\": \"bank\"},{\"transactionDate\": {\"$gte\": \"" + startDate + "\"}}]}}")
	fmt.Println(queryAbtString)
	resultsAbtIterator, err := stub.GetQueryResult(queryAbtString)
	resultsBankIterator, err := stub.GetQueryResult(queryBankString)
	fmt.Println(resultsAbtIterator)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsAbtIterator.Close()
	defer resultsBankIterator.Close()

	var abtRecord transaction
	var bankRecord transaction

	for resultsAbtIterator.HasNext() {
		queryAbtResponse, err := resultsAbtIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		for resultsBankIterator.HasNext() {

			queryBankResponse, err := resultsBankIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}
			err = json.Unmarshal(queryAbtResponse.Value, &abtRecord)
			err = json.Unmarshal(queryBankResponse.Value, &bankRecord)
			if err != nil {
				return shim.Error("Error while unmarshal response")
			}

			if abtRecord.TransactionDate == bankRecord.TransactionDate &&
				abtRecord.MediaNo == bankRecord.MediaNo &&
				abtRecord.EquipmentID == bankRecord.EquipmentID &&
				abtRecord.EquipmentType == bankRecord.EquipmentType &&
				abtRecord.MediaType == bankRecord.MediaType {
				// change status to matched
				abtRecord.MatchStatus = "0002"
				bankRecord.MatchStatus = "0002"

				abtRecordBytes, err := json.Marshal(abtRecord)
				bankRecordBytes, err := json.Marshal(bankRecord)
				if err != nil {
					return shim.Error("Error while marshal struct")
				}
				stub.PutState(queryAbtResponse.Key, abtRecordBytes)
				stub.PutState(queryBankResponse.Key, bankRecordBytes)

				break
			}
		}

	}
	return shim.Success(nil)
}
