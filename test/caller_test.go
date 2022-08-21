package test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	caller "github.com/0xhoang/solidity/build/caller/deploy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *CallerTestSuite) TestDeploy() {
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
	nonce, err := suite.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := suite.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, suite.chainId)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	fmt.Printf("gas price = %v\n", gasPrice.String())

	address, tx, _, err := caller.DeployCaller(auth, suite.client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

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

		fmt.Printf("tx %v is blocked, status = %v\n", tx.Hash().String(), receipt.Status)
		break
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestCallerTestSuite(t *testing.T) {
	suite.Run(t, new(CallerTestSuite))
}
