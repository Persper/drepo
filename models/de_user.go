// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

/// Push the user info to IPFS and record the new ipfsHash in the blockchain
func PushUserInfo(contextUser *User) (err error) {
	// Do some checks
	if contextUser.IsOrganization() {
		return nil
	}
	if !canPushToBlockchain(contextUser) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Get the corresponding user.
	var user *User
	user = &User{ID: contextUser.ID}
	hasUser, err := x.Get(user)
	if err != nil {
		return fmt.Errorf("Can not get user data: %v", err)
	}

	if hasUser {
		// Step 1: Encode user data into JSON format
		user_data, err2 := json.Marshal(user)
		if err2 != nil {
			return fmt.Errorf("Can not encode user data: %v", err2)
		}

		// Step 2: Put the encoded data into IPFS
		c := fmt.Sprintf("echo '%s' | ipfs add ", user_data)
		cmd := exec.Command("sh", "-c", c)
		out, err3 := cmd.Output()
		if err3 != nil {
			return fmt.Errorf("Push User to IPFS: fails: %v", err3)
		}
		ipfsHash := strings.Split(string(out), " ")[1]

		// Step3: Modify the ipfsHash in the smart contract
		// TODO: setUserTableIpfsHash(ipfsHash)
		ipfsHash = ipfsHash
	}

	return nil
}

// Get the new ipfsHash from the blockchain and get the user info from IPFS
func GetUserInfo(contextUser *User) (err error) {
	// Step1: getUserTableIPFSHash() from the smart contract
	ipfsHash := "QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN"

	// Step2: get the ipfs file and get the user data
	c := fmt.Sprintf("ipfs cat ", ipfsHash)
	cmd := exec.Command("sh", "-c", c)
	user_data, err2 := cmd.Output()
	if err2 != nil {
		return fmt.Errorf("Get User data from IPFS: fails: %v", err)
	}

	// Step3: unmarshall user data
	var newUser = new(User)
	err2 = json.Unmarshal(user_data, &newUser)
	if err2 != nil {
		return fmt.Errorf("Can not decode data: %v", err)
	}

	// Step4: write into the local database
	// TODO:
	CreateUser(newUser)

	return nil
}
