.PHONY: grpc-go

grpc-go:
	@protoc -I ./ grpc/account/accountpb/account.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/account/accountpb/account.pb.go grpc/account/accountpb/account.pb.go
	@protoc -I ./ grpc/block/blockpb/block.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockpb/block.pb.go grpc/block/blockpb/block.pb.go
	@protoc -I ./ grpc/chain/chainpb/chain.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/chain/chainpb/chain.pb.go grpc/chain/chainpb/chain.pb.go
	@protoc -I ./ grpc/state/statepb/state.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb/state.pb.go grpc/state/statepb/state.pb.go
	@protoc -I ./ grpc/transaction/transactionpb/transaction.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionpb/transaction.pb.go grpc/transaction/transactionpb/transaction.pb.go
	@protoc -I ./ grpc/validator/validatorpb/validator.proto --go_out=plugins=grpc:.
	@mv github.com/figment-networks/oasis-rpc-proxy/grpc/validator/validatorpb/validator.pb.go grpc/validator/validatorpb/validator.pb.go
	@rm -rvf github.com
