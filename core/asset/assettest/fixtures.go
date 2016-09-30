package assettest

import (
	"context"
	"testing"
	"time"

	"chain/core/account"
	"chain/core/asset"
	"chain/core/txbuilder"
	"chain/encoding/json"
	"chain/errors"
	"chain/protocol"
	"chain/protocol/bc"
	"chain/protocol/state"
	"chain/testutil"
)

func CreateAccountFixture(ctx context.Context, t testing.TB, keys []string, quorum int, alias string, tags map[string]interface{}) string {
	if keys == nil {
		keys = []string{testutil.TestXPub.String()}
	}
	if quorum == 0 {
		quorum = len(keys)
	}
	acc, err := account.Create(ctx, keys, quorum, alias, tags, nil)
	if err != nil {
		testutil.FatalErr(t, err)
	}
	return acc.ID
}

func CreateAssetFixture(ctx context.Context, t testing.TB, keys []string, quorum int, def map[string]interface{}, alias string, tags map[string]interface{}) bc.AssetID {
	if len(keys) == 0 {
		keys = []string{testutil.TestXPub.String()}
	}

	if quorum == 0 {
		quorum = len(keys)
	}
	var initialBlockHash bc.Hash

	asset, err := asset.Define(ctx, keys, quorum, def, initialBlockHash, alias, tags, nil)
	if err != nil {
		testutil.FatalErr(t, err)
	}

	return asset.AssetID
}

func IssueAssetsFixture(ctx context.Context, t testing.TB, c *protocol.Chain, assetID bc.AssetID, amount uint64, accountID string) state.Output {
	if accountID == "" {
		accountID = CreateAccountFixture(ctx, t, nil, 0, "", nil)
	}
	dest := NewAccountControlAction(bc.AssetAmount{AssetID: assetID, Amount: amount}, accountID, nil)

	assetAmount := bc.AssetAmount{AssetID: assetID, Amount: amount}

	src := NewIssueAction(assetAmount, nil) // does not support reference data
	tpl, err := txbuilder.Build(ctx, nil, []txbuilder.Action{dest, src}, time.Now().Add(time.Minute))
	if err != nil {
		testutil.FatalErr(t, err)
	}

	SignTxTemplate(t, ctx, tpl, &testutil.TestXPrv)

	err = txbuilder.FinalizeTx(ctx, c, bc.NewTx(*tpl.Transaction))
	if err != nil {
		testutil.FatalErr(t, err)
	}

	return state.Output{
		Outpoint: bc.Outpoint{Hash: tpl.Transaction.Hash(), Index: 0},
		TxOutput: *tpl.Transaction.Outputs[0],
	}
}

func Transfer(ctx context.Context, t testing.TB, c *protocol.Chain, actions []txbuilder.Action) *bc.Tx {
	template, err := txbuilder.Build(ctx, nil, actions, time.Now().Add(time.Minute))
	if err != nil {
		t.Log(errors.Stack(err))
		t.Fatal(err)
	}

	SignTxTemplate(t, ctx, template, &testutil.TestXPrv)

	tx := bc.NewTx(*template.Transaction)
	err = txbuilder.FinalizeTx(ctx, c, tx)
	if err != nil {
		t.Log(errors.Stack(err))
		t.Fatal(err)
	}

	return tx
}

func NewIssueAction(assetAmount bc.AssetAmount, referenceData json.Map) *asset.IssueAction {
	return &asset.IssueAction{
		TTL:           24 * time.Hour,
		AssetAmount:   assetAmount,
		ReferenceData: referenceData,
	}
}

func NewAccountSpendAction(amt bc.AssetAmount, accountID string, txHash *bc.Hash, txOut *uint32, refData json.Map) *account.SpendAction {
	return &account.SpendAction{
		AssetAmount:   amt,
		AssetAlias:    "",
		TxHash:        txHash,
		TxOut:         txOut,
		AccountID:     accountID,
		AccountAlias:  "",
		ReferenceData: refData,
	}
}

func NewAccountControlAction(amt bc.AssetAmount, accountID string, refData json.Map) *account.ControlAction {
	return &account.ControlAction{
		AssetAmount:   amt,
		AccountID:     accountID,
		ReferenceData: refData,
	}
}
