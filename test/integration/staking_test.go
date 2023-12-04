//go:build integration

package integration

import (
	"context"
	"fmt"
	"os"
	"time"

	"testing"

	"github.com/jjangg96/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/oasisprotocol/oasis-core/go/common/quantity"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestGetStaking(t *testing.T) {
	height := 6200000

	var start time.Time

	ctx := context.Background()

	host := os.Getenv("OASIS_RPC_PROXY_HOST")

	conn, err := grpc.Dial(host, grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1073741824)),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(1073741824)),
	)
	if err != nil {
		t.Error(err)
	}

	stateServer := statepb.NewStateServiceClient(conn)

	start = time.Now()
	data, err := stateServer.GetStakingByHeight(ctx, &statepb.GetStakingByHeightRequest{Height: int64(height), OmitAccountsAndDelegations: true})
	fmt.Printf("GetStakingByHeight with omit duration: %f\n", time.Since(start).Seconds())
	if err != nil {
		t.Error(err)
	}

	totalSupply := &quantity.Quantity{}
	totalSupply.UnmarshalBinary(data.Staking.TotalSupply)

	commonPool := &quantity.Quantity{}
	commonPool.UnmarshalBinary(data.Staking.CommonPool)

	minDelegationAmount := &quantity.Quantity{}
	minDelegationAmount.UnmarshalBinary(data.Staking.Parameters.MinDelegationAmount)

	debondingInterval := data.Staking.Parameters.DebondingInterval

	fmt.Println(totalSupply, commonPool, data.Staking.Parameters.DebondingInterval, minDelegationAmount)

	start = time.Now()
	stakingState, err := stateServer.GetStakingByHeight(ctx, &statepb.GetStakingByHeightRequest{Height: int64(height)})
	fmt.Printf("GetStakingByHeight duration: %f\n", time.Since(start).Seconds())
	if err != nil {
		t.Error(err)
	}

	totalSupply2 := &quantity.Quantity{}
	totalSupply2.UnmarshalBinary(stakingState.Staking.TotalSupply)
	require.Equal(t, totalSupply, totalSupply2)

	commonPool2 := &quantity.Quantity{}
	commonPool2.UnmarshalBinary(stakingState.Staking.CommonPool)
	require.Equal(t, commonPool, commonPool2)

	minDelegationAmount2 := &quantity.Quantity{}
	minDelegationAmount2.UnmarshalBinary(stakingState.Staking.Parameters.MinDelegationAmount)
	require.Equal(t, minDelegationAmount, minDelegationAmount2)

	debondingInterval2 := stakingState.Staking.Parameters.DebondingInterval

	require.Equal(t, debondingInterval, debondingInterval2)

	fmt.Println(totalSupply2, commonPool2, debondingInterval2, minDelegationAmount2)
}
