var config = require('config.json');
var _ = require('lodash');
var jwt = require('jsonwebtoken');
var bcrypt = require('bcryptjs');
var Q = require('q');
var mongo = require('mongoskin');
var Hyperledger = require('./hyperledger.service');
//var mongoUtil = require('mongoPoolUtil');
//var db = mongoUtil.getDb();
var chainUser = require('./utils/chainUser.service')
var chainUtils = require('./utils/chainUtils.service');
var db = mongo.db(config.connectionString, { native_parser: true });
var hfc = require('fabric-client');
var fs = require('fs');
var path = require('path');
var eh;
hfc.setConfigSetting('request-timeout', 60000);
hfc.addConfigFile(path.join(__dirname, './../network-config.json'));
var ORGS = hfc.getConfigSetting('network-config');
chainUser.setORGS(ORGS);
var client = new hfc();
var bcrypt = require('bcryptjs'); 
db.bind('users');

var service = {};

service.getAllTransactions = getAllTransactions;
service.addTransactionsRecords = addTransactionsRecords;

module.exports = service;

function getAllTransactions(username) {
    var deferred = Q.defer();
    console.log("Node service called");
    console.log(username);
    let parseResponse;
    var req = {
        chaincodeId: config.chaincode_id,
        argList: ["queryTransactionsByOrganisation", "BANK"],
        organisation: "org1",
        targets: ["org1", "org2"],
        channel: config.channel_name,
        userName: "ABT"
    };

    Hyperledger.queryChaincode(req)
        .then(function (response) {
            console.log("Test ",response);
            deferred.resolve(response);
        })
        .catch(function (err) {
            console.log("Error on BlockChain", err)
            deferred.reject(err);
        });

    return deferred.promise;
}

function addTransactionsRecords(){
    var deferred = Q.defer();
    let args = JSON.parse("{\"transactionDate\":\"20181217 19:00:60\",\"transactionType\":\"T0002\",\"mediaType\":\"0001\",\"mediaNo\":\"11131231372\",\"equipmentType\":\"0001\",\"equipmentId\":\"146\",\"stopId\":\"153\",\"routeId\":\"4\",\"locationLat\":\"9628.983583\",\"locationLong\":\"7138.784599\",\"companyId\":\"3\",\"operatorId\":\"176\",\"isDeleted\":\"N\",\"createdBy\":\"150\",\"createdDate\":\"20180601 10:30:00.62\",\"updatedBy\":\"\",\"updatedDate\":\"20180605 10:30:00.62\",\"pgmId\":\"AWPR0032\",\"version\":\"1\",\"vehicleNo\":\"9876543738\",\"deviceTransId\":\"9876543745\",\"shiftId\":\"1\",\"tripId\":\"129\",\"subShiftId\":\"4\"}");
	let fcn = "initTransaction";
    var recordCount = 20;

	for(var i = 10; i < recordCount; i++){
        args.mediaType= i.toString().padStart(4, "0");
        // FOR ABT
		console.log("ABT invoked for media type --> " + args.mediaType);

        var abtReq = {
            chaincodeId: config.chaincode_id,
            argList: ["initTransaction",JSON.stringify(args),"ABT","Test"],
            organisation: "org1",
            targets: ["org1", "org2"],
            channel: config.channel_name,
            userName: "ABT"
        };
        //invoke.invokeChaincode(peers, channelName, chaincodeName, fcn, [JSON.stringify(args)], "abtUser", "ABT"); 
        Hyperledger.invokeChaincode(abtReq)
            .then(function (response) {
                console.log("ABT Response :", response)
                deferred.resolve(response);
            })
            .catch(function (err) {
                console.log("Error on BlockChain ABT Method", err)
                deferred.reject(err);
            });

        //FOR BANK
		console.log("Bank invoked for media type --> " + args.mediaType);
		//invoke.invokeChaincode(peers, channelName, chaincodeName, fcn, [JSON.stringify(args)], "bankUser", "Bank");
        var bankReq = {
            chaincodeId: config.chaincode_id,
            argList: ["initTransaction",JSON.stringify(args),"BANK","Test"],
            organisation: "org1",
            targets: ["org1", "org2"],
            channel: config.channel_name,
            userName: "ABT"
        };
        //invoke.invokeChaincode(peers, channelName, chaincodeName, fcn, [JSON.stringify(args)], "abtUser", "ABT"); 
        Hyperledger.invokeChaincode(bankReq)
            .then(function (response) {
                console.log("BANK Response :", response)
                deferred.resolve(response);
            })
            .catch(function (err) {
                console.log("Error on BlockChain BANK Method", err)
                deferred.reject(err);
            });
    
    }
	return deferred.promise;
}