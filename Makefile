.PHONY: grpc-go

grpc-go:
	@protoc -I ./ grpc/account/accountpb/account.proto --go_out=plugins=grpc:.
	@protoc -I ./ grpc/block/blockpb/block.proto --go_out=plugins=grpc:.
	@protoc -I ./ grpc/chain/chainpb/chain.proto --go_out=plugins=grpc:.
	@protoc -I ./ grpc/state/statepb/state.proto --go_out=plugins=grpc:.
	@protoc -I ./ grpc/transaction/transactionpb/transaction.proto --go_out=plugins=grpc:.
	@protoc -I ./ grpc/validator/validatorpb/validator.proto --go_out=plugins=grpc:.
