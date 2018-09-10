docker-compose -f docker-compose-orderer.yml restart
docker-compose -f docker-compose-mitel-peers.yml restart
docker-compose -f docker-compose-biloxi-peers.yml restart
docker-compose -f docker-compose-atlanta-peers.yml restart
docker-compose -f ./mitel-node/docker-compose-api.yaml restart
docker-compose -f ./biloxi-node/docker-compose-api.yaml restart
docker-compose -f ./atlanta-node/docker-compose-api.yaml restart