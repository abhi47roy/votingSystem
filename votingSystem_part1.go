package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type VoteSystem struct {
}

type VotingStruct struct {
	VotingTitle    string         `json:"voting_title"`
	VotingQuestion string         `json:"voting_question"`
	Options        []string       `json:"voting_options"`
	VoteCounts     []VoteCount    `json:"voting_counts"`
	PollCounts     map[string]int `json:"polls_counts"`
}

type VoteCount struct {
	VoteOption string `json:"vote_option"`
	VoterName  string `json:"voter_name"`
}

func (s *VoteSystem) Init(APIstub shim.ChaincodeStubInterface) sc.Response {

	return shim.Success(nil)

}

func (s *VoteSystem) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	fmt.Println("Invoke is running " + args[0])
	// Handle different functions

	if args[0] == "initLedger" {
		fmt.Println("initLedger is running ")
		return s.InitLedger(APIstub)
	} else if args[0] == "createPoll" {
		fmt.Println("createPoll is running ")
		return s.createPoll(APIstub, args)
	} else if args[0] == "updatePollCount" {
		fmt.Println("updatePollCount is running ")
		return s.updatePollCount(APIstub, args)
	} else if function == "query" {
		switch args[0] {
		case "queryPoll":
			fmt.Println("queryPoll is running ")
			return s.queryPoll(APIstub, args)
		case "getAllPollTitle":
			fmt.Println("getAllPollTitle is running ")
			return s.getAllPollTitle(APIstub)
		case "getPollQuestion":
			fmt.Println("getPollQuestion is running ")
			return s.getPollQuestion(APIstub, args)
		default:
			fmt.Println("Invalid function")
			return shim.Error("Invalid function")
		}
	}

	fmt.Println("invoke did not find func: " + args[0]) //error

	return shim.Error("Received unknown function invocation")
}

// Dump Questions
func (s *VoteSystem) InitLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

	quesOneOptions := []string{"Option1", "Option2"}

	pollCountMap1 := make(map[string]int)
	fmt.Println("The Map is before Inserting  pollCountMap1: ", pollCountMap1)
	//var quesOptions1 []string

	for i := 0; i < len(quesOneOptions); i++ {
		fmt.Println("The value of ", quesOneOptions[i])
		pollCountMap1[quesOneOptions[i]] = 0
	}

	uniqueHash1 := s.getPackageHash("Title 1")

	voteQuestions := []VotingStruct{

		VotingStruct{VotingTitle: "Title 1", VotingQuestion: "Question 1", Options: quesOneOptions, PollCounts: pollCountMap1},
	}

	fmt.Println("********************")
	fmt.Println("The voteQuestion Array ", voteQuestions)

	i := 0

	for i < len(voteQuestions) {
		fmt.Printf("Inside the Ledger Method")
		fmt.Println(" i is ", i)
		voteQuesAsBytes, _ := json.Marshal(voteQuestions[i])
		APIstub.PutState(uniqueHash1, voteQuesAsBytes)
		i = i + 1
	}

	return shim.Success([]byte("The Dump Question has been successfully saved on the blockchain"))

}

// Add new poll on blockchain
// **************************************************************************************
func (s *VoteSystem) createPoll(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) <= 2 {
		return shim.Error("Number of Argument are incorrect ")
	}
	fmt.Println("We are creating a new poll ")
	fmt.Println("The title of vote Poll ", args[1])
	fmt.Println("The question of vote Poll ", args[2])

	var quesOptions []string
	pollCountMap := make(map[string]int)
	fmt.Println("The Map is before Inserting : ", pollCountMap)

	for i := 3; i < len(args); i++ {
		fmt.Println("The value of ", args[i])
		quesOptions = append(quesOptions, args[i])
		pollCountMap[args[i]] = 0
	}

	fmt.Println("The Map after inserting the element ", pollCountMap)

	var vote = VotingStruct{VotingTitle: args[1], VotingQuestion: args[2], Options: quesOptions, PollCounts: pollCountMap}

	uniqueKey := s.getPackageHash(args[1])
	voteAsBytes, _ := json.Marshal(vote)
	APIstub.PutState(uniqueKey, voteAsBytes)

	fmt.Printf("We have created the POLL:  %s", uniqueKey)

	return shim.Success([]byte("Your Poll Question has been successfully saved on the blockchain"))

}

// Add a new Poll Question Ends
// **************************************************************************************

func (s *VoteSystem) queryPoll(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	fmt.Println("The QueryPoll has been called with args :", args[1])
	uniqueKey := s.getPackageHash(args[1])
	voteAsBytes, _ := APIstub.GetState(uniqueKey)

	return shim.Success(voteAsBytes)
}

func (s *VoteSystem) getPackageHash(pollTitle string) string {
	uniqueHash := pollTitle

	hash256 := sha256.New()
	hash256.Write([]byte(uniqueHash))

	packageHash := fmt.Sprintf("%x", hash256.Sum(nil))
	return packageHash
}

// Update the poll votes
// **************************************************************************************

func (s *VoteSystem) updatePollCount(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	var voteQuestion VotingStruct
	uniqueKey := s.getPackageHash(args[1])
	voteAsBytes, _ := APIstub.GetState(uniqueKey)
	json.Unmarshal(voteAsBytes, &voteQuestion)

	fmt.Println("Orginal data before :", voteQuestion)
	fmt.Println("The args[0] :", args[0])
	fmt.Println("The args[1] :", args[1])
	fmt.Println("The args[2] :", args[2])
	fmt.Println("The args[3] :", args[3])

	fmt.Println("The Options :", voteQuestion.Options)

	// Updating the User Details with Option Selected Starts
	var voteSelected VoteCount
	voteSelected.VoteOption = args[2]
	voteSelected.VoterName = args[3]
	voteQuestion.VoteCounts = append(voteQuestion.VoteCounts, voteSelected)
	// Updating the User Details with Option Selected Ends

	fmt.Println(voteQuestion.PollCounts)
	//myValue, _ := voteQuestion.PollCounts[args[2]]

	fmt.Println("The Vote Option selected Count Before Poll ", voteQuestion.PollCounts[args[2]])
	voteQuestion.PollCounts[args[2]] = voteQuestion.PollCounts[args[2]] + 1
	fmt.Println("The Vote Option selected Count After Poll ", voteQuestion.PollCounts[args[2]])
	fmt.Println("The question  after: ", voteQuestion)
	voteQuesAsBytes, _ := json.Marshal(voteQuestion)
	APIstub.PutState(uniqueKey, voteQuesAsBytes)
	fmt.Println("The voting has been done.")
	return shim.Success([]byte("Your vote for this poll has been stored on blockchain"))
}

// **************************************************************************************

func (s *VoteSystem) getHistoryForQuestion(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	uniqueKey := s.getPackageHash(args[1])

	fmt.Printf("- start Get Poll History: %s\n", uniqueKey)

	resultsIterator, err := stub.GetHistoryForKey(uniqueKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- Poll History returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

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
		buffer.WriteString("{\"key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return buffer.Bytes(), nil
}

// Get All Poll Titles
// **************************************************************************************

func (s *VoteSystem) getAllPollTitle(APIstub shim.ChaincodeStubInterface) sc.Response {

	var queryString string
	queryString = fmt.Sprintf("{\"selector\":{\"$and\":[{\"voting_title\":{\"$ne\": \"%s\"}}]}}", "")

	fmt.Println("The query string formed")
	fmt.Println(queryString)
	resultsIterator, err := getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		fmt.Println("The error is :", err)
		return shim.Error("Error while geting ")
	}
	fmt.Println("The sucess all request :")

	return shim.Success(resultsIterator)

}

func main() {
	// Create a new Smart Contract
	err := shim.Start(new(VoteSystem))
	if err != nil {
		fmt.Printf("Error creating new Vote Poll Contract : %s", err)
	}

}
