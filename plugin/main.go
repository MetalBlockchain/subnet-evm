// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"fmt"
	"os"

	"github.com/MetalBlockchain/metalgo/utils/logging"
	"github.com/MetalBlockchain/metalgo/utils/ulimit"
	"github.com/MetalBlockchain/metalgo/version"
	"github.com/MetalBlockchain/metalgo/vms/rpcchainvm"

	"github.com/MetalBlockchain/subnet-evm/plugin/evm"
)

func main() {
	printVersion, err := PrintVersion()
	if err != nil {
		fmt.Printf("couldn't get config: %s", err)
		os.Exit(1)
	}
	if printVersion {
		fmt.Printf("Subnet-EVM/%s [MetalGo=%s, rpcchainvm=%d]\n", evm.Version, version.Current, version.RPCChainVMProtocol)
		os.Exit(0)
	}
	if err := ulimit.Set(ulimit.DefaultFDLimit, logging.NoLog{}); err != nil {
		fmt.Printf("failed to set fd limit correctly due to: %s", err)
		os.Exit(1)
	}
	rpcchainvm.Serve(&evm.VM{})
}
