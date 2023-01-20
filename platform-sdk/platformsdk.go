package platformsdk

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

	"github.com/Theraine/theraine-sdk-go/abigen-binds/platform" 
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

	address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	fmt.Println("Contract Address: ", address.Hex())

	instance, _ := createInstance(address, client)
	signerAddress, auth := transact(client, chainID)

	choice := ""
    prompt := &survey.Select{
        Message: "Please select an operation:",
        Options: []string{
			"addPlan", 
			"removePlan", 
			"getPlan", 
			"getPlans",
			"checkUserStatus",
			"exit",
		},
    }
    survey.AskOne(prompt, &choice, nil)

    switch choice {
    case "addPlan":
		var price int
		var duration int
		promptPrice := &survey.Input{
			Message: "Please enter plan price:",
		}
		promptDuration := &survey.Input{
			Message: "Please enter plan duration:",
		}
		survey.AskOne(promptPrice, &price, nil)
		survey.AskOne(promptDuration, &duration, nil)
		fmt.Println("Adding a plan")
		addPlan(auth, instance, client, price, duration)
    case "removePlan":
		var id uint8
		prompt := &survey.Input{
			Message: "Please enter plan id:",
		}
		survey.AskOne(prompt, &id, nil)
		fmt.Println("Removing plan Id")
		removePlan(auth, instance, client, id)
    case "getPlan":
		var id uint8
		prompt := &survey.Input{
			Message: "Please enter plan id:",
		}
		survey.AskOne(prompt, &id, nil)
		fmt.Println("Getting plan")
		getPlan(instance, id)
	case "getPlans":
		fmt.Println("Getting plans")
		getPlans(instance)
	case "checkUserStatus":
		var address string
		prompt := &survey.Input{
			Message: "Please enter user address:",
		}
		survey.AskOne(prompt, &address, nil)
		fmt.Println("Checking user status")
		checkUserStatus(signerAddress, instance, address)
    case "exit":
        fmt.Println("Exiting the program")
    }
}

func createInstance(address common.Address, client *ethclient.Client) (*platform.Platform, error) {
	instance, err := platform.NewPlatform(address, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance, err
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

func addPlan(auth *bind.TransactOpts, instance *platform.Platform, client *ethclient.Client, price int, duration int) {
	tx, err := instance.AddPlan(auth, big.NewInt(int64(price)), big.NewInt(int64(duration)))
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 1 {
		fmt.Println("Plan added successfully")
		fmt.Printf("https://goerli.etherscan.io/tx/%s\n", tx.Hash().Hex())
	} else {
		fmt.Println("Transaction failed")
	}
}

func removePlan(auth *bind.TransactOpts, instance *platform.Platform, client *ethclient.Client, id uint8) {
	tx, err := instance.RemovePlan(auth, id)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 1 {
		fmt.Println("Plan removed successfully")
		fmt.Printf("https://goerli.etherscan.io/tx/%s\n", tx.Hash().Hex())
	} else {
		fmt.Println("Transaction failed")
	}
}

func getPlan(instance *platform.Platform, id uint8) {
	plan, err := instance.GetPlan(nil, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Plan: ", plan)
}

func getPlans(instance *platform.Platform) {
	plans, err := instance.GetPlans(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Plans: ", plans)
}

func checkUserStatus(signerAddress common.Address, instance *platform.Platform, address string) {
	status, err := instance.UserStatus(&bind.CallOpts{From: signerAddress}, common.HexToAddress(address))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status: ", status)
}