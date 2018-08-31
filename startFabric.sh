#!/usr/bin/env bash

#docker exec mitel-cli peer channel create -o orderer.mitel.com:7050 -c webinarchannel -f /etc/hyperledger/fabric/peer/channel-artifacts/channel.tx
#docker exec -e "CORE_PEER_ADDRESS=peer0.mitel.com:7051" mitel-cli  peer channel join -b webinarchannel.block
docker exec -e "CORE_PEER_ADDRESS=peer1.org1.com:7051" mitel-cli peer channel join -b webinarchannel.block
docker exec -e "CORE_PEER_ADDRESS=peer1.org1.com:7051" mitel-cli peer channel join -b webinarchannel.block

#docker exec -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" cli peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/go/chaincode_example02
#docker exec -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051"  cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n mycc -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "OR ('Org1MSP.member','Org2MSP.member')"
#docker exec -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051"  cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n mycc -v 1.0 -c '{"Args":["invoke","a","b","20"]}'
#docker exec -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" cli peer chaincode query -C mychannel -n mycc -c '{"Args":["query","a"]}'