package test

import (
	"fmt"
	"github.com/0xhoang/solidity/pkg"
	"log"
)

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *CallerTestSuite) TestDeploy() {
	auth := suite.getBroker()

	address, tx, _, err := pkg.DeployCaller(auth, suite.client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sc caller = ", address.Hex())

	calleeAddr, calleeTx, _, err := pkg.DeployCallee(auth, suite.client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sc callee = ", calleeAddr.Hex())

	counterAddr, counterTx, _, err := pkg.DeployCounter(auth, suite.client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sc couter = ", counterAddr.Hex())

	suite.Waiting(tx.Hash())
	suite.Waiting(calleeTx.Hash())
	suite.Waiting(counterTx.Hash())
}
