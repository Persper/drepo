// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

/*
import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = `
{
  "address": "1a9ec3b0b807464e6d3398a59d6b0a369bf422fa",
  "crypto": {
    "cipher": "aes-128-ctr",
    "ciphertext": "a471054846fb03e3e271339204420806334d1f09d6da40605a1a152e0d8e35f3",
    "cipherparams": {
      "iv": "44c5095dc698392c55a65aae46e0b5d9"
    },
    "kdf": "scrypt",
    "kdfparams": {
      "dklen": 32,
      "n": 262144,
      "p": 1,
      "r": 8,
      "salt": "e0a5fbaecaa3e75e20bccf61ee175141f3597d3b1bae6a28fe09f3507e63545e"
    },
    "mac": "cb3f62975cf6e7dfb454c2973bdd4a59f87262956d5534cdc87fb35703364043"
  },
  "id": "e08301fb-a263-4643-9c2b-d28959f66d6a",
  "version": 3
}
`

func getUserIpfsHash() (error, string) {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial("/path/to/your/.ethereum/testnet/geth.ipc")
	if err != nil {
		return fmt.Errorf("Failed to connect to the Ethereum client: %v", err), nil
	}
	// Get the smart contract via address
	userCont, err := userContract.NewMyToken(common.HexToAddress("0x5e300171d7dc10e43f959877dba98a44df5d1466"), conn)
	if err != nil {
		return fmt.Errorf("Failed to instantiate a Token contract: %v", err), nil
	}
	// Get the user ipfsHash
	ipfsHash, err := userCont.getIpfsHash(nil)
	if err != nil {
		return fmt.Errorf("Failed to get the ipfs hash of the user: %v", err), nil
	}
	return nil, ipfsHash
}

func setUserIpfsHash() (error, string) {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial("/path/to/your/.ethereum/testnet/geth.ipc")
	if err != nil {
		return fmt.Errorf("Failed to connect to the Ethereum client: %v", err), nil
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "your password")
	if err != nil {
		return fmt.Errorf("could not create auth: %v", err), nil
	}

}

*/
