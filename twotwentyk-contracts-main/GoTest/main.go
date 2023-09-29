package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionResult struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int64  `json:"id"`
	Hash    string `json:"result"`
}

var (
	rpc       = "https://polygon-mumbai.g.alchemy.com/v2/PTi4_upQduxYasRED-Hy1EBldW7pBstt"
	walletKey = "fbc533ef71a592206fb846e1ce4f2ca6ee3cb08388d13b2ff555745c28eb0c74"

	value    = big.NewInt(0)
	gasLimit = uint64(20_000_000)
	chainId  = big.NewInt(80001)

	ownerAddr = common.HexToAddress("0xf22679BBaf587B9c776D0A25a05e64B214f19CAd")

	collectionAddr         = common.HexToAddress("0x2AC89A2eAF323558bac2358b2B508DaD2891D330")
	packAddress            = common.HexToAddress("0x8BAf45B2D89ff1A15AE854C612b6EB9F56941929")
	categoryAddress        common.Address
	yearAddress            common.Address
	dayMonthAddress        common.Address
	craftingCardAddress    common.Address
	triggerAddress         common.Address
	collectionCreatedTopic = "0xfef410e8debc1749c63e587a214c2c77bedf014092686ff45b910c64465a0925"

	collectionId    = big.NewInt(2)
	standardPackAmt = big.NewInt(300)
	premiumPackAmt  = big.NewInt(200)
	ellitePackAmt   = big.NewInt(100)
	totalPackAmt    = big.NewInt(300 + 200 + 100)

	packOpenerAddr = common.HexToAddress("0x5D6A8216590cEEeFd149f3CB26ad096B8516bB00")
	craftingAddr   = common.HexToAddress("0xC64da131E851E9623Afc5c616e1c218eEf008bB4")
	predictionAddr = common.HexToAddress("0xd36294eFADa12EC5091257aF6C33CF9dbd8e6997")

	collectionInstance   *Collection
	cardpackInstance     *Cardpack
	categoryInstance     *Category
	yearInstance         *Year
	daymonthInstance     *Daymonth
	craftingcardInstance *CraftingCard
	triggerInstance      *Trigger
	PackopenerInstance   *Packopener
	IdentifyInstance     *Crafting
	predictionInstance   *Prediction

	err error
)

func main() {
	deployCollections()
	getAddresses(collectionId)
	setMinterAndCrafter()
	createCard()
	openPack()
	// craftingCollection()

}
func deployCollections() {

	client, _ := ethclient.Dial(rpc)
	auth := getTransactor()

	fmt.Println(standardPackAmt)
	fmt.Println(premiumPackAmt)
	fmt.Println(ellitePackAmt)
	fmt.Println(totalPackAmt)
	collectionInstance, err = NewCollection(collectionAddr, client)
	tx, _ := collectionInstance.CreateCollection(auth, "https://example.com/collection/1", standardPackAmt, premiumPackAmt, ellitePackAmt, totalPackAmt)

	fmt.Printf("create collection tx hash: %s", tx.Hash())

}

func getAddresses(collectionId *big.Int) (common.Address, common.Address, common.Address, common.Address, common.Address, common.Address, error) {
	client, _ := ethclient.Dial(rpc)

	var (
		logs []types.Log
	)

	collectionABI, err := CollectionMetaData.GetAbi()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{collectionAddr},
		Topics:    [][]common.Hash{{common.HexToHash(collectionCreatedTopic)}},
	}

	logs, err = client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
		return common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, err
	}

	for _, vlog := range logs {
		topic := vlog.Topics[0].Hex()
		fmt.Println(topic)
		if topic == collectionCreatedTopic {

			data, _ := collectionABI.Unpack("CollectionCreated", vlog.Data)
			// spew.Dump("topic", data)

			tokenId := data[0].(*big.Int)
			fmt.Println(tokenId, collectionId, tokenId.Cmp(collectionId))
			if tokenId.Cmp(collectionId) == 0 {
				packAddress = data[2].(common.Address)
				categoryAddress = data[3].(common.Address)
				yearAddress = data[4].(common.Address)
				dayMonthAddress = data[5].(common.Address)
				craftingCardAddress = data[6].(common.Address)
				triggerAddress = data[7].(common.Address)

				fmt.Println("card pack address", packAddress)
				fmt.Println("category address", categoryAddress)
				fmt.Println("year address", yearAddress)
				fmt.Println("daymonth address", dayMonthAddress)
				fmt.Println("trigger address", triggerAddress)
				fmt.Println("craftingcard address", craftingCardAddress)

				return packAddress, categoryAddress, yearAddress, dayMonthAddress, craftingCardAddress, triggerAddress, err
			}
		}
	}
	return common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, err
}

func setMinterAndCrafter() {
	client, _ := ethclient.Dial(rpc)
	var (
		auth *bind.TransactOpts
	)

	cardpackInstance, err = NewCardpack(packAddress, client)
	auth = getTransactor()
	cardpackInstance.ChangeAdmin(auth, ownerAddr)
	auth = getTransactor()
	cardpackInstance.ChangeMinter(auth, packOpenerAddr)

	categoryInstance, err = NewCategory(categoryAddress, client)
	auth = getTransactor()
	categoryInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	categoryInstance.ChangeCrafter(auth, craftingAddr)

	yearInstance, err = NewYear(yearAddress, client)
	auth = getTransactor()
	yearInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	yearInstance.ChangeCrafter(auth, craftingAddr)

	daymonthInstance, err = NewDaymonth(dayMonthAddress, client)
	auth = getTransactor()
	daymonthInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	daymonthInstance.ChangeCrafter(auth, craftingAddr)

	craftingcardInstance, err = NewCraftingCard(craftingCardAddress, client)
	auth = getTransactor()
	craftingcardInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	craftingcardInstance.ChangeCrafter(auth, craftingAddr)

	triggerInstance, err = NewTrigger(triggerAddress, client)
	auth = getTransactor()
	triggerInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	triggerInstance.ChangeCrafter(auth, predictionAddr)

}

func createCard() {
	var (
		tokenURI = "https://example.com/cardpack/1"
	)
	client, _ := ethclient.Dial(rpc)
	cardpackInstance, err = NewCardpack(packAddress, client)
	auth := getTransactor()
	tx, _ := cardpackInstance.CreateStandardCard(auth, tokenURI)
	fmt.Printf("card pack tx hash: %s", tx.Hash())
}

func openPack() {
	var (
		cardPackTokenId = big.NewInt(0)

		categoryTokenURIs     = []string{"https://example.com/category/1"}
		yearTOkenURIs         = []string{"https://example.com/year/1"}
		daymonthTokenURIs     = []string{"https://example.com/daymonth/1"}
		craftingCardTokenURIs = []string{"https://example.com/craftingcard/1", "https://example.com/craftingcard/2"}
		triggerTokenURIs      = []string{"https://example.com/trigger/1"}
	)

	client, _ := ethclient.Dial(rpc)
	PackopenerInstance, err = NewPackopener(packOpenerAddr, client)
	auth := getTransactor()

	tx, _ := PackopenerInstance.OpenPack(
		auth,
		cardPackTokenId,
		packAddress,
		categoryAddress,
		yearAddress,
		dayMonthAddress,
		craftingCardAddress,
		triggerAddress,
		categoryTokenURIs,
		yearTOkenURIs,
		daymonthTokenURIs,
		craftingCardTokenURIs,
		triggerTokenURIs,
	)

	fmt.Printf("open pack tx hash: %s", tx.Hash())
}

func craftingCollection() {
	// var (
	// 	categoryTokenId = big.NewInt(12)
	// 	dayMonthTokenId = big.NewInt(6)
	// 	yearTokenId     = big.NewInt(9)
	// 	craftingTokenId = big.NewInt(6)

	// 	collectionTokenUrl = "https://example.com/identity/1"
	// )
	// client, _ := ethclient.Dial(rpc)

	// auth := getTransactor()

	// identify, _ := NewCrafting(common.HexToAddress(craftingAddr), client)
	// tx, _ := identify.CraftCollection(auth, dayMonthTokenId, yearTokenId, categoryTokenId, craftingTokenId, collectionTokenUrl)

	// fmt.Printf("crafting identify tx hash: %s", tx.Hash())
}

func craftingPrediction() {
	// var (
	// 	identityTokenId           = big.NewInt(1)
	// 	triggerTokenIds           = []*big.Int{big.NewInt(1)}
	// 	predictionCraftingTokenId = big.NewInt(1)
	// 	predictionTokenUrl        = "https://example.com/prediction/1"
	// )
	// client, _ := ethclient.Dial(rpc)

	// auth := getTransactor()

	// prediction, _ := NewPrediction(common.HexToAddress(predictionAddr), client)
	// tx, _ := prediction.CraftCollection(auth, identityTokenId, triggerTokenIds, predictionCraftingTokenId, predictionTokenUrl)

	// fmt.Printf("tx sent: %s", tx.Hash())
}

func getTransactor() *bind.TransactOpts {
	client, _ := ethclient.Dial(rpc)

	privatekey, _ := crypto.HexToECDSA(walletKey)

	gasPrice, _ := client.SuggestGasPrice(context.Background())
	auth := bind.NewKeyedTransactor(privatekey)
	auth.Nonce = big.NewInt(int64(getNonce(privatekey)))
	auth.Value = value
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	signfunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewEIP155Signer(chainId)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privatekey)
		if err != nil {
			fmt.Printf("Error signing: %v\n", err)
			os.Exit(1)
			return nil, nil
		}
		return tx.WithSignature(signer, signature)
	}
	auth.Signer = signfunc

	return auth
}

func getNonce(privatekey *ecdsa.PrivateKey) uint64 {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privatekey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Println("from", fromAddress, nonce)

	return nonce
}
