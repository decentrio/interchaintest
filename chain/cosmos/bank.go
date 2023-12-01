package cosmos

import (
	"context"
	"fmt"

	"cosmossdk.io/math"
	sdkmath "cosmossdk.io/math"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/strangelove-ventures/interchaintest/v8/ibc"
)

// BankSend sends tokens from one account to another.
func (tn *ChainNode) BankSend(ctx context.Context, keyName string, amount ibc.WalletAmount) error {
	_, err := tn.ExecTx(ctx,
		keyName, "bank", "send", keyName,
		amount.Address, fmt.Sprintf("%s%s", amount.Amount.String(), amount.Denom),
	)
	return err
}

// BankMultiSend sends an amount of token from one account to multiple accounts.
func (tn *ChainNode) BankMultiSend(ctx context.Context, keyName string, addresses []string, amount math.Int, denom string) error {
	cmd := append([]string{"bank", "multi-send", keyName}, addresses...)
	cmd = append(cmd, fmt.Sprintf("%s%s", amount, denom))

	_, err := tn.ExecTx(ctx, keyName, cmd...)
	return err
}

// GetBalance fetches the current balance for a specific account address and denom.
// Implements Chain interface
func (c *CosmosChain) GetBalance(ctx context.Context, address string, denom string) (sdkmath.Int, error) {
	grpcConn, err := grpc.Dial(
		c.GetNode().hostGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return sdkmath.Int{}, err
	}
	defer grpcConn.Close()

	res, err := banktypes.NewQueryClient(grpcConn).Balance(ctx, &banktypes.QueryBalanceRequest{Address: address, Denom: denom})
	return res.Balance.Amount, err
}

// BankGetBalance is an alias for GetBalance
func (c *CosmosChain) BankGetBalance(ctx context.Context, address string, denom string) (sdkmath.Int, error) {
	return c.GetBalance(ctx, address, denom)
}

// AllBalances fetches an account address's balance for all denoms it holds
func (c *CosmosChain) BankAllBalances(ctx context.Context, address string) (types.Coins, error) {
	grpcConn, err := grpc.Dial(
		c.GetNode().hostGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	res, err := banktypes.NewQueryClient(grpcConn).AllBalances(ctx, &banktypes.QueryAllBalancesRequest{Address: address})
	return res.GetBalances(), err
}

// BankDenomMetadata fetches the metadata of a specific coin denomination
func (c *CosmosChain) BankDenomMetadata(ctx context.Context, denom string) (*banktypes.Metadata, error) {
	grpcConn, err := grpc.Dial(
		c.GetNode().hostGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	res, err := banktypes.NewQueryClient(grpcConn).DenomMetadata(ctx, &banktypes.QueryDenomMetadataRequest{Denom: denom})
	return &res.Metadata, err
}

func (c *CosmosChain) BankQueryDenomMetadataByQueryString(ctx context.Context, denom string) (*banktypes.Metadata, error) {
	grpcConn, err := grpc.Dial(
		c.GetNode().hostGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	res, err := banktypes.NewQueryClient(grpcConn).DenomMetadataByQueryString(ctx, &banktypes.QueryDenomMetadataByQueryStringRequest{Denom: denom})
	return &res.Metadata, err
}

func (c *CosmosChain) BankQueryDenomOwners(ctx context.Context, denom string) ([]*banktypes.DenomOwner, error) {
	grpcConn, err := grpc.Dial(
		c.GetNode().hostGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	res, err := banktypes.NewQueryClient(grpcConn).DenomOwners(ctx, &banktypes.QueryDenomOwnersRequest{Denom: denom})
	return res.DenomOwners, err
}

func (c *CosmosChain) BankQueryDenomsMetadata(ctx context.Context) ([]banktypes.Metadata, error) {
	grpcConn, err := grpc.Dial(
		c.GetNode().hostGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	res, err := banktypes.NewQueryClient(grpcConn).DenomsMetadata(ctx, &banktypes.QueryDenomsMetadataRequest{})
	return res.Metadatas, err
}

func (c *CosmosChain) BankQueryParams(ctx context.Context) (*banktypes.Params, error) {
	grpcConn, err := grpc.Dial(
		c.GetNode().hostGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	res, err := banktypes.NewQueryClient(grpcConn).Params(ctx, &banktypes.QueryParamsRequest{})
	return &res.Params, err
}

func (c *CosmosChain) BankQuerySendEnabled(ctx context.Context, denoms []string) ([]*banktypes.SendEnabled, error) {
	grpcConn, err := grpc.Dial(
		c.GetNode().hostGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	res, err := banktypes.NewQueryClient(grpcConn).SendEnabled(ctx, &banktypes.QuerySendEnabledRequest{
		Denoms: denoms,
	})
	return res.SendEnabled, err
}

func (c *CosmosChain) BankQuerySpendableBalance(ctx context.Context, address, denom string) (*types.Coin, error) {
	grpcConn, err := grpc.Dial(
		c.GetNode().hostGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	res, err := banktypes.NewQueryClient(grpcConn).SpendableBalanceByDenom(ctx, &banktypes.QuerySpendableBalanceByDenomRequest{
		Address: address,
		Denom:   denom,
	})
	return res.Balance, err
}

func (c *CosmosChain) BankQuerySpendableBalances(ctx context.Context, address string) (*types.Coins, error) {
	grpcConn, err := grpc.Dial(
		c.GetNode().hostGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	res, err := banktypes.NewQueryClient(grpcConn).SpendableBalances(ctx, &banktypes.QuerySpendableBalancesRequest{Address: address})
	return &res.Balances, err
}

func (c *CosmosChain) BankQueryTotalSupply(ctx context.Context, address string) (*types.Coins, error) {
	grpcConn, err := grpc.Dial(
		c.GetNode().hostGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	res, err := banktypes.NewQueryClient(grpcConn).TotalSupply(ctx, &banktypes.QueryTotalSupplyRequest{})
	return &res.Supply, err
}

func (c *CosmosChain) BankQueryTotalSupplyOf(ctx context.Context, address string) (*types.Coin, error) {
	grpcConn, err := grpc.Dial(
		c.GetNode().hostGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	res, err := banktypes.NewQueryClient(grpcConn).SupplyOf(ctx, &banktypes.QuerySupplyOfRequest{Denom: address})

	return &res.Amount, err
}
