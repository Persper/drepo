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

type DeBranch struct {
	Name               string `xorm:"UNIQUE(protect_branch)"`
	Protected          bool
	RequirePullRequest bool
	EnableWhitelist    bool
	WhitelistUserIDs   string `xorm:"TEXT"`
	WhitelistTeamIDs   string `xorm:"TEXT"`
}

func transferBranchToDeBranch(branch *ProtectBranch, deBranch *DeBranch) {
	deBranch.Name = branch.Name
	deBranch.Protected = branch.Protected
	deBranch.RequirePullRequest = branch.RequirePullRequest
	deBranch.EnableWhitelist = branch.EnableWhitelist
	deBranch.WhitelistUserIDs = branch.WhitelistUserIDs
	deBranch.WhitelistTeamIDs = branch.WhitelistTeamIDs
}

func transferDeBranchToBranch(repo *Repository, deBranch *DeBranch, branch *ProtectBranch) {
	// branch.ID can be generated at any time
	// TODO: branch.ID
	branch.RepoID = repo.ID
	branch.Name = deBranch.Name
	branch.Protected = deBranch.Protected
	branch.RequirePullRequest = deBranch.RequirePullRequest
	branch.EnableWhitelist = deBranch.EnableWhitelist
	branch.WhitelistUserIDs = deBranch.WhitelistUserIDs
	branch.WhitelistTeamIDs = deBranch.WhitelistTeamIDs
}

func PushBranchInfo(user *User, branch *ProtectBranch) error {
	if !canPushToBlockchain(user) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step1: encode branch data into JSON format
	deBranch := new(DeBranch)
	transferBranchToDeBranch(branch, deBranch)
	branch_data, err := json.Marshal(deBranch)
	if err != nil {
		return fmt.Errorf("Can not encode branch data: %v\n", err)
	}

	// Step2: push the encoded data into IPFS
	c := fmt.Sprintf("echo '%s' | ipfs add ", branch_data)
	cmd := exec.Command("sh", "-c", c)
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Push branch to IPFS: %v\n", err)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step3: modify the ipfsHash in the smart contract
	// TODO: setBranchInfo(ipfsHash)
	ipfsHash = ipfsHash
	fmt.Println("Push the protect_branch file to the IPFS: " + ipfsHash)

	return nil
}

func GetBranchInfo(user *User, repo *Repository, ipfsHash string) error {
	// Step1: get the ipfs file and get the branch data
	c := "ipfs cat " + ipfsHash
	cmd := exec.Command("sh", "-c", c)
	branch_data, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Get branch data from IPFS: %v\n", err)
	}

	// Step2: unmarshall pull data
	newDeBranch := new(DeBranch)
	err = json.Unmarshal(branch_data, &newDeBranch)
	if err != nil {
		return fmt.Errorf("Can not decode data: %v\n", err)
	}

	// Step3: write into the local database
	newBranch := new(ProtectBranch)
	transferDeBranchToBranch(repo, newDeBranch, newBranch)
	has, err := x.Get(newBranch)
	if err != nil {
		return fmt.Errorf("Can not search the branch: %v\n", err)
	}
	if !has {
		_, err = x.Insert(newBranch)
		if err != nil {
			return fmt.Errorf("Can not add the branch request to the server: %v\n", err)
		}
	}

	// TODO: patch?
	// TODO: watch

	return nil
}
