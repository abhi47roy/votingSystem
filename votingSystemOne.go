package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

func (s *VoteSystem) getPollQuestion(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	fmt.Println("Get the poll Question")
	uniqueKey := s.getPackageHash(args[1])
	voteQuestionAsBytes, _ := APIstub.GetState(uniqueKey)

	fmt.Println("We have the vote Questions with the options :", voteQuestionAsBytes)

	return shim.Success(voteQuestionAsBytes)

}
