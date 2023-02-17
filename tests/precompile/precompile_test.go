// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package precompile

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/MetalBlockchain/metalgo/api/health"
	"github.com/MetalBlockchain/subnet-evm/tests/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-cmd/cmd"
	ginkgo "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	// Import the solidity package, so that ginkgo maps out the tests declared within the package
	_ "github.com/MetalBlockchain/subnet-evm/tests/precompile/solidity"
)

var startCmd *cmd.Cmd

func TestE2E(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "subnet-evm precompile ginkgo test suite")
}

// BeforeSuite starts an MetalGo process to use for the e2e tests
var _ = ginkgo.BeforeSuite(func() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	wd, err := os.Getwd()
	gomega.Expect(err).Should(gomega.BeNil())
	log.Info("Starting MetalGo node", "wd", wd)
	startCmd, err = utils.RunCommand("./scripts/run.sh")
	gomega.Expect(err).Should(gomega.BeNil())

	// Assumes that startCmd will launch a node with HTTP Port at [utils.DefaultLocalNodeURI]
	healthClient := health.NewClient(utils.DefaultLocalNodeURI)
	healthy, err := health.AwaitReady(ctx, healthClient, 5*time.Second)
	gomega.Expect(err).Should(gomega.BeNil())
	gomega.Expect(healthy).Should(gomega.BeTrue())
	log.Info("MetalGo node is healthy")
})

var _ = ginkgo.AfterSuite(func() {
	gomega.Expect(startCmd).ShouldNot(gomega.BeNil())
	gomega.Expect(startCmd.Stop()).Should(gomega.BeNil())
	// TODO add a new node to bootstrap off of the existing node and ensure it can bootstrap all subnets
	// created during the test
})
