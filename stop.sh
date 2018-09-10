docker-compose -f docker-compose-orderer.yml stop
docker-compose -f docker-compose-mitel-peers.yml stop
docker-compose -f docker-compose-biloxi-peers.yml stop
docker-compose -f docker-compose-atlanta-peers.yml stop
docker-compose -f ./mitel-node/docker-compose-api.yaml stop
docker-compose -f ./biloxi-node/docker-compose-api.yaml stop
docker-compose -f ./atlanta-node/docker-compose-api.yaml stop