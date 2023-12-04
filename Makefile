.PHONY: grpc-go build test docker docker-build docker-push proto

GIT_COMMIT   ?= $(shell git rev-parse HEAD)
GO_VERSION   ?= $(shell go version | awk {'print $$3'})
DOCKER_IMAGE ?= figmentnetworks/oasis-rpc-proxy
DOCKER_TAG   ?= latest

# alias
proto: grpc-go

# Generate proto buffs
grpc-go:
	@protoc -I ./ grpc/account/accountpb/account.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
	@protoc -I ./ grpc/block/blockpb/block.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
	@protoc -I ./ grpc/chain/chainpb/chain.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
	@protoc -I ./ grpc/event/eventpb/event.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
	@protoc -I ./ grpc/state/statepb/state.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
	@protoc -I ./ grpc/transaction/transactionpb/transaction.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
	@protoc -I ./ grpc/validator/validatorpb/validator.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
	@protoc -I ./ grpc/delegation/delegationpb/delegation.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
	@protoc -I ./ grpc/debondingdelegation/debondingdelegationpb/debonding_delegation.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
	@protoc -I ./ grpc/rawdata/rawdatapb/rawdata.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative

# Build the binary
build:
	go build \
		-ldflags "\
			-X github.com/jjangg96/oasis-rpc-proxy/cli.gitCommit=${GIT_COMMIT} \
			-X github.com/jjangg96/oasis-rpc-proxy/cli.goVersion=${GO_VERSION}"

# Run tests
test:
	go test -race -cover ./...

test-integration:
	go test -v -tags=integration ./test/integration

# Build a local docker image for testing
docker:
	docker build -t oasis-rpc-proxy -f Dockerfile .

# Build a public docker image
docker-build:
	docker build \
		-t ${DOCKER_IMAGE}:${DOCKER_TAG} \
		-f Dockerfile \
		.

# Push docker images
docker-push:
	docker tag ${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_IMAGE}:latest
	docker push ${DOCKER_IMAGE}:${DOCKER_TAG}
	docker push ${DOCKER_IMAGE}:latest
