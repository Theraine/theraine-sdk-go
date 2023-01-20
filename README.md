# Theraine SDK Go
Theraine SDK Go is a software development kit that allows developers to interact with a smart contract on the Ethereum blockchain in a simple and easy-to-use way. This SDK provides a set of Go files that are organized into two packages: "parentsdk" and "platformsdk", and a main package that ties everything together.

## Requirements
- Go
- Go packages: "gopkg.in/AlecAivazis/survey.v1", "github.com/joho/godotenv", "github.com/ethereum/go-ethereum/accounts/abi/bind", "github.com/ethereum/go-ethereum/common", "github.com/ethereum/go-ethereum/crypto", "github.com/ethereum/go-ethereum/ethclient"

## Usage
1. Clone the repository.
```bash
git clone https://github.com/Theraine/theraine-sdk-go
```

2. Create a .env file in the root directory of the project, and set the following environment variables:

```
INFURA_KEY=<YOUR_INFURA_API_KEY>
PRIVATE_KEY=<YOUR_PRIVATE_KEY>
CONTRACT_ADDRESS=<PLATFORM_CONTRACT_ADDRESS>
```

3. Install the dependencies.
```bash
go get -d -v
```

4. Run the main file.
```bash
go run main.go
```

5. Follow the prompts to interact with the smart contract.

**Note:** Make sure you are connected to the internet and the Smart Contract is deployed on the blockchain.

## Packages

### parentsdk

This package contains the functions related to creating a new platform. The `Theraine()` function handles the interaction with the smart contract and the user. It loads the environment variables, dials the Ethereum client, and sets the contract address. Then it prompts the user to select an operation and performs the corresponding function.

### platformsdk

This package contains the functions related to interacting with an existing platform. The `Theraine()` function handles the interaction with the smart contract and the user in a similar way to the parentsdk package. It loads the environment variables, dials the Ethereum client, and sets the contract address. Then it prompts the user to select an operation and performs the corresponding function.

### main

This package ties everything together. It imports the `Theraine()` function from both parentsdk and platformsdk packages, and prompts the user to select the package they want to use. It then runs the corresponding function.

## Conclusion

The Theraine SDK Go provides a simple and easy-to-use way to interact with a smart contract on the Ethereum blockchain. It allows developers to create a new platform and interact with existing ones. The SDK is well-documented and easy to understand. It is recommended to use dotenv to store the secret keys
