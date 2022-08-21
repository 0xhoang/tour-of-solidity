package test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	caller "github.com/0xhoang/solidity/build/caller"
	"github.com/0xhoang/solidity/build/caller/deploy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/suite"
	"log"
	"math/big"
	"testing"
	"time"
)

type CallerTestSuite struct {
	suite.Suite
	client           *ethclient.Client
	chainId          *big.Int
	privateKeyDeploy string
	contractCaller   string
	contractCallee   string
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *CallerTestSuite) SetupTest() {
	client, err := ethclient.Dial("https://kovan.infura.io/v3/ab2f10698fd94021a4de1d5168c99a78")
	if err != nil {
		log.Fatal(err)
	}
	suite.client = client
	suite.chainId = big.NewInt(42)
	suite.privateKeyDeploy = "5d2f4e4b26bf65526cb170c6302eeb3d3c2bafbc770f7cbed76085e199034129"
	suite.contractCaller = "0xea3a943c6d7081d28f0998ee4b7add15258ea26b"
	suite.contractCallee = "0x9b2240c7d30d3b2d8129edc4fd94f724f5899b69"
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *CallerTestSuite) TestDeploy() {
	auth := suite.getBroker()

	address, tx, _, err := deploy.DeployCaller(auth, suite.client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sc caller = ", address.Hex())

	ctx := context.Background()
	//todo: check tx status
	index := 0
	for {
		if index == 10 {
			break
		}

		receipt, err := suite.client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Println(err)
			index++
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Printf("tx https://kovan.etherscan.io/tx/%v is blocked, status = %v\n", tx.Hash().String(), receipt.Status)

		break
	}

	address1, tx1, _, err := deploy.DeployCallee(auth, suite.client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sc callee = ", address1.Hex())
	index1 := 0
	for {
		if index1 == 10 {
			break
		}

		receipt1, err := suite.client.TransactionReceipt(ctx, tx1.Hash())
		if err != nil {
			log.Println(err)
			index1++
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Printf("tx https://kovan.etherscan.io/tx/%v is blocked, status = %v\n", tx1.Hash().String(), receipt1.Status)
		break
	}
}

func (suite *CallerTestSuite) TestSC() {
	scCallerAddr := common.HexToAddress(suite.contractCaller)
	scCalleeAddr := common.HexToAddress(suite.contractCallee)
	auth := suite.getBroker()

	callerSc, err := caller.NewCaller(scCallerAddr, suite.client)
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

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestCallerTestSuite(t *testing.T) {
	suite.Run(t, new(CallerTestSuite))
}

func (suite *CallerTestSuite) getBroker() *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(suite.privateKeyDeploy)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("created by = ", fromAddress)
	/*nonce, err := suite.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}*/

	gasPrice, err := suite.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, suite.chainId)
	if err != nil {
		log.Fatal(err)
	}

	//auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	fmt.Printf("gas price = %v\n", gasPrice.String())

	return auth
}
