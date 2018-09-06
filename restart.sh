docker-compose -f docker-compose-orderer.yml restart
docker-compose -f docker-compose-mitel-peers.yml restart
docker-compose -f docker-compose-org1-peers.yml restart
docker-compose -f docker-compose-org2-peers.yml restart
docker-compose -f ./mitel-node/docker-compose-api.yaml restart
docker-compose -f ./org1-node/docker-compose-api.yaml restart
docker-compose -f ./org2-node/docker-compose-api.yaml restart