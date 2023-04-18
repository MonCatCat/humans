package main_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

	"github.com/0x4139/humans/app"
	humansd "github.com/0x4139/humans/cmd/humansd"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := humansd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",       // Test the init cmd
		"humanstest", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, "humans_9000-1"),
	})

	err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome)
	require.NoError(t, err)
}