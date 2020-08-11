.PHONY: grpc-go build test docker docker-build docker-push

GIT_COMMIT   ?= $(shell git rev-parse HEAD)
GO_VERSION   ?= $(shell go version | awk {'print $$3'})
DOCKER_IMAGE ?= figmentnetworks/oasis-rpc-proxy
DOCKER_TAG   ?= latest

# Generate proto buffs
grpc-go:
	@protoc -I ./ grpc/account/accountpb/account.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/account/accountpb/account.pb.go grpc/account/accountpb/account.pb.go
	@protoc -I ./ grpc/block/blockpb/block.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockpb/block.pb.go grpc/block/blockpb/block.pb.go
	@protoc -I ./ grpc/chain/chainpb/chain.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/chain/chainpb/chain.pb.go grpc/chain/chainpb/chain.pb.go
	@protoc -I ./ grpc/event/eventpb/event.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/event/eventpb/event.pb.go grpc/event/eventpb/event.pb.go
	@protoc -I ./ grpc/state/statepb/state.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb/state.pb.go grpc/state/statepb/state.pb.go
	@protoc -I ./ grpc/transaction/transactionpb/transaction.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionpb/transaction.pb.go grpc/transaction/transactionpb/transaction.pb.go
	@protoc -I ./ grpc/validator/validatorpb/validator.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/validator/validatorpb/validator.pb.go grpc/validator/validatorpb/validator.pb.go
	@protoc -I ./ grpc/delegation/delegationpb/delegation.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/delegation/delegationpb/delegation.pb.go grpc/delegation/delegationpb/delegation.pb.go
	@protoc -I ./ grpc/debondingdelegation/debondingdelegationpb/debonding_delegation.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/debondingdelegation/debondingdelegationpb/debonding_delegation.pb.go grpc/debondingdelegation/debondingdelegationpb/debonding_delegation.pb.go
	@rm -rvf github.com

# Build the binary
build:
	go build \
		-ldflags "\
			-X github.com/figment-networks/oasis-rpc-proxy/cli.gitCommit=${GIT_COMMIT} \
			-X github.com/figment-networks/oasis-rpc-proxy/cli.goVersion=${GO_VERSION}"

# Run tests
test:
	go test -race -cover ./...

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