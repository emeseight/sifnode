package keeper

import (
	"github.com/Sifchain/sifnode/x/clp/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"io/ioutil"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier { //nolint
	querier := Querier{keeper}

	tracer.Start(
		tracer.WithService("sifnode"),
		tracer.WithEnv("test"),
	)
	defer tracer.Stop()

	// Start a root span.
	span := tracer.StartSpan("clp.query")
	defer span.Finish()

	// // Create a child of it, computing the time needed to read a file.
	// child := tracer.StartSpan("read.file", tracer.ChildOf(span.Context()))
	// child.SetTag(ext.ResourceName, "test.json")

	// // Perform an operation.
	// _, err := ioutil.ReadFile("~/test.json")

	// // We may finish the child span using the returned error. If it's
	// // nil, it will be disregarded.
	// child.Finish(tracer.WithError(err))
	// tracer.Stop()

	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		child := tracer.StartSpan("read.file", tracer.ChildOf(span.Context()))
		child.SetTag(ext.ResourceName, "test.json")

		switch path[0] {
		case types.QueryPool:
			return queryPool(ctx, path[1:], req, legacyQuerierCdc, querier)
		case types.QueryPools:
			return queryPools(ctx, path[1:], legacyQuerierCdc, querier, child)
		case types.QueryLiquidityProvider:
			return queryLiquidityProvider(ctx, path[1:], req, legacyQuerierCdc, querier)
		case types.QueryLiquidityProviderData:
			return queryLiquidityProviderData(ctx, path[1:], req, legacyQuerierCdc, querier)
		case types.QueryAssetList:
			return queryAssetList(ctx, path[1:], req, legacyQuerierCdc, querier)
		case types.QueryLPList:
			return queryLPList(ctx, path[1:], req, legacyQuerierCdc, querier)
		case types.QueryAllLP:
			return queryAllLP(ctx, path[1:], req, legacyQuerierCdc, querier)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown clp query endpoint")
		}
	}
}

func queryPool(ctx sdk.Context, path []string, req abci.RequestQuery, legacyQuerierCdc *codec.LegacyAmino, querier Querier) ([]byte, error) { //nolint
	var params types.PoolReq
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.GetPool(sdk.WrapSDKContext(ctx), &params)
	if err != nil {
		return nil, err
	}
	bz, err := legacyQuerierCdc.MarshalJSON(&res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryPools(ctx sdk.Context, path []string, legacyQuerierCdc *codec.LegacyAmino, querier Querier, child tracer.Span) ([]byte, error) {
	// Perform an operation.
	_, err := ioutil.ReadFile("~/test.json")
	// We may finish the child span using the returned error. If it's
	// nil, it will be disregarded.
	child.Finish(tracer.WithError(err))

	res, err := querier.GetPools(sdk.WrapSDKContext(ctx), &types.PoolsReq{})
	if err != nil {
		return nil, err
	}
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryLiquidityProvider(ctx sdk.Context, path []string, req abci.RequestQuery, legacyQuerierCdc *codec.LegacyAmino, querier Querier) ([]byte, error) { //nolint
	var params types.LiquidityProviderReq
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.GetLiquidityProvider(sdk.WrapSDKContext(ctx), &params)
	if err != nil {
		return nil, err
	}
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryLiquidityProviderData(ctx sdk.Context, path []string, req abci.RequestQuery, legacyQuerierCdc *codec.LegacyAmino, querier Querier) ([]byte, error) { //nolint
	var params types.LiquidityProviderDataReq
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.GetLiquidityProviderData(sdk.WrapSDKContext(ctx), &params)
	if err != nil {
		return nil, err
	}
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryAssetList(ctx sdk.Context, path []string, req abci.RequestQuery, legacyQuerierCdc *codec.LegacyAmino, querier Querier) ([]byte, error) { //nolint
	var params types.AssetListReq
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.GetAssetList(sdk.WrapSDKContext(ctx), &params)
	if err != nil {
		return nil, err
	}
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res.Assets)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryLPList(ctx sdk.Context, path []string, req abci.RequestQuery, legacyQuerierCdc *codec.LegacyAmino, querier Querier) ([]byte, error) { //nolint
	var params types.LiquidityProviderListReq
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.GetLiquidityProviderList(sdk.WrapSDKContext(ctx), &params)
	if err != nil {
		return nil, err
	}
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res.LiquidityProviders)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryAllLP(ctx sdk.Context, path []string, req abci.RequestQuery, legacyQuerierCdc *codec.LegacyAmino, querier Querier) ([]byte, error) { //nolint
	res, err := querier.GetLiquidityProviders(sdk.WrapSDKContext(ctx), &types.LiquidityProvidersReq{})
	if err != nil {
		return nil, err
	}
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res.LiquidityProviders)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
