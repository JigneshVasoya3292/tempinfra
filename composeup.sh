docker-compose -f docker-compose-orderer.yml up -d
docker-compose -f docker-compose-mitel-peers.yml up -d
docker-compose -f docker-compose-biloxi-peers.yml up -d
docker-compose -f docker-compose-atlanta-peers.yml up -d
docker-compose -f ./mitel-node/docker-compose-api.yaml up -d
docker-compose -f ./biloxi-node/docker-compose-api.yaml up -d
docker-compose -f ./atlanta-node/docker-compose-api.yaml up -d