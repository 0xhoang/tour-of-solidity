package test

import (
	"fmt"
	"github.com/0xhoang/solidity/pkg"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"time"
)

func (suite *CallerTestSuite) TestCaller() {
	scCallerAddr := common.HexToAddress("0x500E0d41B03c706e935925F2de7Cff6392569831")
	scCalleeAddr := common.HexToAddress("0xdbD8A24A7a624Fe5E030e1e163DEBc353380462E")
	auth := suite.getBroker()

	callerSc, err := pkg.NewCaller(scCallerAddr, suite.client)
	if err != nil {
		log.Fatal(err)
	}

	value, err := callerSc.SomeAction(nil, scCalleeAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SomeAction = ", value.String())

	time.Sleep(10 * time.Second)
	tx, err := callerSc.StoreAction(auth, scCalleeAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("StoreAction tx = ", tx.Hash().String())

	time.Sleep(10 * time.Second)

	tx, err = callerSc.SomeUnsafeAction(auth, scCalleeAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SomeUnsafeAction tx = ", tx.Hash().String())
	time.Sleep(10 * time.Second)

	tx, err = callerSc.StoreAction(auth, scCalleeAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("StoreAction tx = ", tx.Hash().String())
}

func (suite *CallerTestSuite) TestCounter() {
	scAddr := common.HexToAddress("0x32740184fDB6cBF990dd948668C9B2C2081C08E5")
	auth := suite.getBroker()

	sc, err := pkg.NewCounter(scAddr, suite.client)
	if err != nil {
		log.Fatal(err)
	}

	tx1, err := sc.Inc(auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx1 = ", tx1.Hash().String())
	time.Sleep(10 * time.Second)

	tx2, err := sc.Get(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx2 = ", tx2.String())
	
	tx3, err := sc.Desc(auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx3 = ", tx3.Hash().String())
	time.Sleep(10 * time.Second)
}
