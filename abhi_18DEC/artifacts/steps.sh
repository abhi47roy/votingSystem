./bin/cryptogen generate --config=./crypto-config.yaml

export FABRIC_CFG_PATH=$PWD

./bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block

export CHANNEL_NAME=commonchannel

./bin/configtxgen -profile CommonChannel -outputCreateChannelTx ./channel-artifacts/commonchannel.tx -channelID $CHANNEL_NAME

./bin/configtxgen -profile CommonChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP

./bin/configtxgen -profile CommonChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP

docker-compose -f docker-compose-cli.yaml up

################## Second Terminal  ###################

#******************************************************************************************************
#To interact with the blockchain open docker cli :
#******************************************************************************************************           
docker exec -it cli /bin/bash		

peer channel create -o orderer.nec.com:7050 -c commonchannel -f ./channel-artifacts/commonchannel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/nec.com/orderers/orderer.nec.com/msp/tlscacerts/tlsca.nec.com-cert.pem

#******************************************************************************************************
#Setting Peer Environment Variables and joining the peers nodes to the channel
#******************************************************************************************************
CORE_PEER_ID=peer0.org1.nec.com
CORE_PEER_ADDRESS=peer0.org1.nec.com:7051
CORE_PEER_LOCALMSPID=Org1MSP
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.nec.com/users/Admin@org1.nec.com/msp
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.nec.com/peers/peer0.org1.nec.com/tls/ca.crt
peer channel join -b commonchannel.block

CORE_PEER_ID=peer0.org2.nec.com
CORE_PEER_ADDRESS=peer0.org2.nec.com:7051
CORE_PEER_LOCALMSPID=Org2MSP
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.nec.com/users/Admin@org2.nec.com/msp
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.nec.com/peers/peer0.org2.nec.com/tls/ca.crt
peer channel join -b commonchannel.block



#******************************************************************************************************
#Update anchor Peers (for all peers in the network)
#******************************************************************************************************
CORE_PEER_ID=peer0.org1.nec.com
CORE_PEER_ADDRESS=peer0.org1.nec.com:7051
CORE_PEER_LOCALMSPID=Org1MSP
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.nec.com/users/Admin@org1.nec.com/msp
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.nec.com/peers/peer0.org1.nec.com/tls/ca.crt
peer channel update -o orderer.nec.com:7050 -c commonchannel -f ./channel-artifacts/Org1MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/nec.com/orderers/orderer.nec.com/msp/tlscacerts/tlsca.nec.com-cert.pem

CORE_PEER_ID=peer0.org2.nec.com
CORE_PEER_ADDRESS=peer0.org2.nec.com:7051
CORE_PEER_LOCALMSPID=Org2MSP
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.nec.com/users/Admin@org2.nec.com/msp
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.nec.com/peers/peer0.org2.nec.com/tls/ca.crt
peer channel update -o orderer.nec.com:7050 -c commonchannel -f ./channel-artifacts/Org2MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/nec.com/orderers/orderer.nec.com/msp/tlscacerts/tlsca.nec.com-cert.pem


#******************************************************************************************************
#Install chaincode (for all peers in the network)
#******************************************************************************************************
CORE_PEER_ID=peer0.org1.nec.com
CORE_PEER_ADDRESS=peer0.org1.nec.com:7051
CORE_PEER_LOCALMSPID=Org1MSP
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.nec.com/users/Admin@org1.nec.com/msp
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.nec.com/peers/peer0.org1.nec.com/tls/ca.crt
peer chaincode install -n fabcar1 -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/fabcar/go

CORE_PEER_ID=peer0.org2.nec.com
CORE_PEER_ADDRESS=peer0.org2.nec.com:7051
CORE_PEER_LOCALMSPID=Org2MSP
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.nec.com/users/Admin@org2.nec.com/msp
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.nec.com/peers/peer0.org2.nec.com/tls/ca.crt
peer chaincode install -n fabcar1 -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/fabcar/go

#******************************************************************************************************
#Instantiate chaincode (only on one peer)
#******************************************************************************************************

peer chaincode instantiate -o orderer.nec.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/nec.com/orderers/orderer.nec.com/msp/tlscacerts/tlsca.nec.com-cert.pem -C commonchannel -n fabcar1 -v 1.0 -c '{"Args":["init"]}' -P "OR ('Org1MSP.member','Org2MSP.member')"

peer chaincode invoke -o orderer.nec.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/nec.com/orderers/orderer.nec.com/msp/tlscacerts/tlsca.nec.com-cert.pem -C commonchannel -n fabcar1 -c '{"Args":["initLedger"]}'

peer chaincode query -C commonchannel -n fabcar1 -c '{"Args":["queryAllCars"]}'



#******************************************************************************************************
#Install chaincode (for all peers in the network)
#******************************************************************************************************
CORE_PEER_ID=peer0.org1.nec.com
CORE_PEER_ADDRESS=peer0.org1.nec.com:7051
CORE_PEER_LOCALMSPID=Org1MSP
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.nec.com/users/Admin@org1.nec.com/msp
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.nec.com/peers/peer0.org1.nec.com/tls/ca.crt
peer chaincode install -n abtcc7 -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/AbtChainCode

CORE_PEER_ID=peer0.org2.nec.com
CORE_PEER_ADDRESS=peer0.org2.nec.com:7051
CORE_PEER_LOCALMSPID=Org2MSP
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.nec.com/users/Admin@org2.nec.com/msp
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.nec.com/peers/peer0.org2.nec.com/tls/ca.crt
peer chaincode install -n abtcc7 -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/AbtChainCode

peer chaincode instantiate -o orderer.nec.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/nec.com/orderers/orderer.nec.com/msp/tlscacerts/tlsca.nec.com-cert.pem -C commonchannel -n abtcc7 -v 1.0 -c '{"Args":["init"]}' -P "OR ('Org1MSP.member','Org2MSP.member')"