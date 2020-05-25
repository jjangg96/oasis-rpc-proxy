.PHONY: grpc-go

grpc-go:
	@protoc grpc/block/blockpb/block.proto --go_out=plugins=grpc:.
	@protoc grpc/chain/chainpb/chain.proto --go_out=plugins=grpc:.
	@protoc grpc/state/statepb/state.proto --go_out=plugins=grpc:.
	@protoc grpc/transaction/transactionpb/transaction.proto --go_out=plugins=grpc:.
	@protoc grpc/validator/validatorpb/validator.proto --go_out=plugins=grpc:.
