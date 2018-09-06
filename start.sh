docker-compose -f docker-compose-orderer.yml start
docker-compose -f docker-compose-mitel-peers.yml start
docker-compose -f docker-compose-org1-peers.yml start
docker-compose -f docker-compose-org2-peers.yml start
docker-compose -f ./mitel-node/docker-compose-api.yaml start
docker-compose -f ./org1-node/docker-compose-api.yaml start
docker-compose -f ./org2-node/docker-compose-api.yaml start