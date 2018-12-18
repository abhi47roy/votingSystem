angular.module("myDapp", [])
	.controller("dappDemoCtrl", function($scope) {
		
		//---------- Covert Hexadecimal to ASCII -------------------------------
		function hex2ascii(hexx) {
			var hex = hexx.toString();//force conversion
			var str = '';
			for (var i = 0; i < hex.length; i += 2)
				str += String.fromCharCode(parseInt(hex.substr(i, 2), 16));
			return str.replace(/\0/g, '');
		}
		
		$scope.userType = "";
		
		//---------- Contract invocation using web3 -------------------------------
		var web3Client = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
		
		
		// #################### REPLACE it with your ABI #############################
		//var VoterContractABI = JSON.parse('<Your ABI>');
		
		
		// #################### REPLACE it with your Address #############################
		//var VoterContractAddress = '<your_Address>';
		
		
		//var VoterContract = web3Client.eth.contract(VoterContractABI);
		//var VoterContractInstance = VoterContract.at(VoterContractAddress)
		$scope.user
		
		//---------- Get Candidate List ----------------------------------------------
		
		// function getCandidateList() {
		// 	$scope.candidates = VoterContractInstance.getCandidateList();
		// 	var addresses = $scope.candidates[0];
		// 	var firstNames = $scope.candidates[1];
		// 	var lastNames = $scope.candidates[2];
		// 	var voteCounts = $scope.candidates[3];
		// 	//var voteFlags = $scope.candidates[5];
		// 	var tmpArr = [];
			
		// 	for(var i=0; i<addresses.length; i++) {
		// 		var tmpObj = {};
		// 		tmpObj.addresses = hex2ascii(addresses[i]);
		// 		tmpObj.firstNames = hex2ascii(firstNames[i]);
		// 		tmpObj.lastNames = hex2ascii(lastNames[i]);
		// 		tmpObj.voteCounts = voteCounts[i].c[0];
		// 		//tmpObj.voteFlags = voteFlags[i];
		// 		tmpArr.push(tmpObj);
		// 	}
			
		// 	$scope.candidateJsonData = tmpArr;
		
		// }
		// getCandidateList();
		
		// //---------- Create Candidate/Voter -------------------------------
		
		// $scope.createCandidateObj = {
		// 	candidateAddress : "",
		// 	firstName : "",
		// 	lastName : "",
		// 	candidateType : ""
		// };
		// $scope.createCandidateObj.candidateType = "Candidate";
		// $scope.createCandidate = function() {
		// 	try {
		// 		VoterContractInstance.createCandidate($scope.createCandidateObj.candidateAddress, $scope.createCandidateObj.firstName, $scope.createCandidateObj.lastName, $scope.createCandidateObj.candidateType, {from:web3Client.eth.accounts[0], gas: 2000000});
		// 		//var createCandidateStatus = VoterContractInstance.createCandidate("POC3", "Davy", "Jones", "Candidate", {from:web3Client.eth.accounts[0], gas: 2000000});
		// 		updateCandidateDetails($scope.createCandidateObj.candidateAddress, "add");
		// 	} 
		// 	catch(err) {
		// 		console.log(err.message);
		// 	}
		// }
		
		// //---------- Create Candidate/Voter (Batch Process) for testing -------------------------------
		// var createCandidateObjBatch = {
		// 	candidateAddress : "",
		// 	firstName : "",
		// 	lastName : "",
		// 	candidateType : ""
		// };
		
		// function createCandidateBatch(_candidateType, _startCount, _endCount) {
		// 	for(var i=_startCount; i< _endCount; i++) {
		// 		createCandidateObjBatch.candidateAddress = _candidateType + i.toString();
		// 		createCandidateObjBatch.firstName = _candidateType + "_firstName_" + i.toString();
		// 		createCandidateObjBatch.lastName = _candidateType + "_lastName_" + i.toString();
		// 		if(_candidateType == "VOTER") {
		// 			createCandidateObjBatch.candidateType = "Voter";
		// 		} else if(_candidateType == "CANDIDATE") {
		// 			createCandidateObjBatch.candidateType = "Candidate";
		// 		}
		// 		console.log(createCandidateObjBatch);
		// 		try {
		// 			//VoterContractInstance.createCandidate(createCandidateObjBatch.candidateAddress, createCandidateObjBatch.firstName, createCandidateObjBatch.lastName, createCandidateObjBatch.candidateType, {from:web3Client.eth.accounts[0], gas: 2000000});
		// 		}
		// 		catch(err) {
		// 			console.log(err.message);
		// 		}
		// 	}
		// }
		
		// //createCandidateBatch("VOTER");
		
		
		// //---------- Update candidateJsonData -------------------------------
		
		// function updateJSON(src, newRecord) {
		// 	return src.map(function(item) {
		// 	  return (item.addresses === newRecord.addresses) ? newRecord : item;
		// 	});
		// }
		
		// function updateCandidateDetails(_address, activity){
		// 	var tmpCandidateDetail = VoterContractInstance.getCandidateDetail(_address, {from:web3Client.eth.accounts[0], gas: 200000});
					
		// 	var tmpObj = {};
		// 	tmpObj.addresses = hex2ascii(tmpCandidateDetail[0]);
		// 	tmpObj.firstNames = hex2ascii(tmpCandidateDetail[1]);
		// 	tmpObj.lastNames = hex2ascii(tmpCandidateDetail[2]);
		// 	tmpObj.voteCounts = tmpCandidateDetail[3].c[0];
		// 	//tmpObj.voteFlags = tmpCandidateDetail[5];
			
		// 	if(activity == "update") {
		// 		$scope.candidateJsonData = updateJSON($scope.candidateJsonData, tmpObj);			//Update candidateJsonData
		// 	} else if(activity == "add") {
		// 		$scope.candidateJsonData.push(tmpObj);
		// 	}
		// }
		
		// //---------- Vote for candidate -------------------------------
		// // Variables like voterIndex , voterAddress are for future development
		
		// var voterIndex = 0;
		
		// $scope.vote = function(instance, indexVal) {
		// 	var voterAddress = "VOTER" + voterIndex.toString();
		// 	console.log(voterAddress);
		// 	voterIndex += 1;
		// 	try {
		// 		VoterContractInstance.voteForCandidate(instance.addresses, voterAddress, {from:web3Client.eth.accounts[0], gas: 2000000});
		// 		updateCandidateDetails(instance.addresses, "update");
		// 	}
		// 	catch(err) {
		// 		console.log(err.message);
		// 	}
		// }
		
		$scope.showUserEntryFlag = false;
		$scope.userType = function(value) {
			$scope.showUserEntryFlag = true;
			console.log("Test");
			console.log(value);
			if(value == "registered"){
			$scope.BorrowerFlag  = true;
			$scope.LenderFlag  = false;}
			if(value == "newuser"){
				$scope.BorrowerFlag  = false;
				$scope.LenderFlag  = true;}
			
		}

		$scope.return = function() {
			$scope.showUserEntryFlag = false;
			$scope.BorrowerFlag  = false;
			$scope.LenderFlag  = false;
			
			
		}

	});

	//"POC1","Jack","Sparrow","Candidate"