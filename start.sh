docker-compose -f docker-compose-orderer.yml start
docker-compose -f docker-compose-mitel-peers.yml start
docker-compose -f docker-compose-biloxi-peers.yml start
docker-compose -f docker-compose-atlanta-peers.yml start
docker-compose -f ./mitel-node/docker-compose-api.yaml start
docker-compose -f ./biloxi-node/docker-compose-api.yaml start
docker-compose -f ./atlanta-node/docker-compose-api.yaml start