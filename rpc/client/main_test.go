package client_test

import (
	"os"
	"testing"

	meapp "github.com/tendermint/merkleeyes/app"
	nm "github.com/consensus/node"
	rpctest "github.com/consensus/rpc/test"
)

var node *nm.Node

func TestMain(m *testing.M) {
	// start a tendermint node (and merkleeyes) in the background to test against
	app := meapp.NewMerkleEyesApp("", 100)
	node = rpctest.StartTendermint(app)
	code := m.Run()

	// and shut down proper at the end
	node.Stop()
	node.Wait()
	os.Exit(code)
}
