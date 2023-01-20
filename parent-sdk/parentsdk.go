package parentsdk

import (
	"fmt"
	"log"
	"math/big"
	"context"
	"crypto/ecdsa"
	"os"

	"gopkg.in/AlecAivazis/survey.v1"
	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Theraine/theraine-sdk-go/abigen-binds/parent" 
)

func Theraine() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dial := "https://goerli.infura.io/v3/" + os.Getenv("INFURA_KEY")
	if dial == "https://goerli.infura.io/v3/" {
		log.Fatal("Please set INFURA_KEY in .env file")
	}

	client, err := ethclient.Dial(dial)
	if err != nil {
		log.Fatal(err)
	}
	
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chain id: ", chainID)

	address := common.HexToAddress("0xF124B6990530Fa4c1cB7f1A74c8b7D3D41E34F9F")
	fmt.Println("Contract Address: ", address.Hex())

	instance, _ := createInstance(address, client)
	signerAddress, auth := transact(client, chainID)

	choice := ""
    prompt := &survey.Select{
        Message: "Please select an operation:",
        Options: []string{"createPlatform", "getUserLatestPlatform", "getUserPlatforms", "exit"},
    }
    survey.AskOne(prompt, &choice, nil)

    switch choice {
    case "createPlatform":
		var details string
		prompt := &survey.Input{
			Message: "Please enter platform details:",
		}
		survey.AskOne(prompt, &details, nil)
		fmt.Println("Creating a platform")
		CreatePlatform(auth, instance, []byte(details), client)
		platform := GetUserLatestPlatform(signerAddress, instance)
		fmt.Println("platform address: ", platform.Platform.Hex())
    case "getUserLatestPlatform":
		fmt.Println("Getting user latest platform")
		platform := GetUserLatestPlatform(signerAddress, instance)
		fmt.Println("platform: ", platform)
    case "getUserPlatforms":
		fmt.Println("Getting user platforms")
		GetUserPlatforms(signerAddress, instance)
    case "exit":
        fmt.Println("Exiting the program")
    }
}

func createInstance(address common.Address, client *ethclient.Client) (*parent.Parent, error) {
	instance, err := parent.NewParent(address, client)
	if err != nil {
		log.Fatal(err)
	}

	return instance, nil
}

func transact(client *ethclient.Client, chainID *big.Int) (common.Address, *bind.TransactOpts) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	signerAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), signerAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice

	return signerAddress, auth
}

func CreatePlatform(auth *bind.TransactOpts, instance *parent.Parent, details []byte, client *ethclient.Client) {
	tx, err := instance.CreatePlatform(auth, details)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 1 {
		fmt.Println("Platform created successfully")
		fmt.Printf("https://goerli.etherscan.io/tx/%s\n", tx.Hash().Hex())
	} else {
		fmt.Println("Platform creation failed")
	}
}

func GetPlatform(instance *parent.Parent, platformID *big.Int) {
	platform, err := instance.GetPlatform(nil, platformID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("platform: ", platform)
}

func GetPlatforms(instance *parent.Parent) {
	platforms, err := instance.GetPlatforms(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("platforms: ", platforms)
}

func GetUserPlatforms(signerAddress common.Address, instance *parent.Parent) {
	platforms, err := instance.GetUserPlatforms(&bind.CallOpts{From: signerAddress})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("user platforms: ", platforms)
}

func GetUserLatestPlatform(signerAddress common.Address, instance *parent.Parent) parent.ParentContractPlatform {
	platform, err := instance.GetUserLatestPlatform(&bind.CallOpts{From: signerAddress})
	if err != nil {
		log.Fatal(err)
	}

	return platform
}
