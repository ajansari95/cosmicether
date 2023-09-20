package keeper_test

import (
	"io"
	"testing"
	"time"

	"github.com/ajansari95/cosmicether/app"
	"github.com/ajansari95/cosmicether/x/ethstate/types"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
	"github.com/stretchr/testify/require"
)

func newApp(t *testing.T, options ...func(*app.App)) *app.App {
	t.Helper()

	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = app.DefaultNodeHome
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue

	app := app.New(
		log.NewNopLogger(),
		dbm.NewMemDB(),
		io.Discard,
		true,
		map[int64]bool{},
		t.TempDir(),
		5,
		app.MakeEncodingConfig(),
		appOptions,
	)

	return app
}

func TestGetSetSlotData(t *testing.T) {
	app := newApp(t)
	ctx := app.BaseApp.NewContext(true, tmproto.Header{Height: 1, ChainID: "mercury-1", Time: time.Now().UTC()})

	kpr := app.EthstateKeeper
	contractAddress := "0xC18360217D8F7Ab5e7c516566761Ea12Ce7F9D72"
	slot := "0x02"

	// Get slot data before setting
	slotData, ok := kpr.GetSlotData(ctx, contractAddress, slot)
	require.False(t, ok)
	require.Equal(t, types.SlotData{}, slotData)

	// Set slot data
	slotData = types.SlotData{
		ContractAddress: contractAddress,
		Slot:            slot,
		Data:            []byte("someData"),
		Proof:           []byte("someProof"),
		Height:          1223231,
		Verified:        false,
	}

	kpr.SetSlotData(ctx, &slotData)
	gotslotData, ok := kpr.GetSlotData(ctx, contractAddress, slot)
	require.True(t, ok)
	require.Equal(t, slotData, gotslotData)

	slotData2 := types.SlotData{
		ContractAddress: contractAddress,
		Slot:            "0x03",
		Data:            []byte("someData2"),
		Proof:           []byte("someProof2"),
		Height:          1223231,
		Verified:        false,
	}

	kpr.SetSlotData(ctx, &slotData2)
	gotslotData2, ok := kpr.GetSlotData(ctx, contractAddress, "0x03")
	require.True(t, ok)
	require.Equal(t, slotData2, gotslotData2)
}
