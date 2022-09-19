package test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
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
	//contractCaller   string
	//contractCallee   string
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *CallerTestSuite) SetupTest() {
	//suite.infura()
	suite.ganache()
}

func (suite *CallerTestSuite) infura() {
	client, err := ethclient.Dial("https://kovan.infura.io/v3/ab2f10698fd94021a4de1d5168c99a78")
	if err != nil {
		log.Fatal(err)
	}
	suite.client = client
	suite.chainId = big.NewInt(42)
	suite.privateKeyDeploy = "5d2f4e4b26bf65526cb170c6302eeb3d3c2bafbc770f7cbed76085e199034129"

	//suite.contractCaller = "0xea3a943c6d7081d28f0998ee4b7add15258ea26b"
	//suite.contractCallee = "0x9b2240c7d30d3b2d8129edc4fd94f724f5899b69"
}

func (suite *CallerTestSuite) ganache() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	suite.client = client
	suite.chainId = big.NewInt(5777)
	suite.privateKeyDeploy = "06de6d50dd2b971bdb1e79d031367526b99fce6b854f6bde1e9e5d8dffd1b765"
	//suite.contractCaller = "0x500E0d41B03c706e935925F2de7Cff6392569831"
	//suite.contractCallee = "0xdbD8A24A7a624Fe5E030e1e163DEBc353380462E"
}

func (suite *CallerTestSuite) Waiting(txHash common.Hash) {
	index1 := 0
	ctx := context.Background()

	for {
		if index1 == 10 {
			break
		}

		receipt1, err := suite.client.TransactionReceipt(ctx, txHash)
		if err != nil {
			log.Println(err)
			index1++
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Printf("tx https://kovan.etherscan.io/tx/%v is blocked, status = %v\n", txHash.String(), receipt1.Status)
		break
	}
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
