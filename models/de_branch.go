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

// The pull table in the IPFS
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

func transferDeBranchToBranch(deBranch *DeBranch, branch *ProtectBranch) {
	// branch.ID =
	// branch.RepoID =
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

	// Step 1: Encode branch data into JSON format
	deBranch := new(DeBranch)
	transferBranchToDeBranch(branch, deBranch)
	branch_data, err := json.Marshal(deBranch)
	if err != nil {
		return fmt.Errorf("Can not encode branch data: %v", err)
	}

	// Step 2: Put the encoded data into IPFS
	c := fmt.Sprintf("echo '%s' | ipfs add ", branch_data)
	cmd := exec.Command("sh", "-c", c)
	out, err2 := cmd.Output()
	if err2 != nil {
		return fmt.Errorf("Push branch to IPFS: fails: %v", err2)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step4: Modify the ipfsHash in the smart contract
	// TODO: setBranchInfo(ipfsHash)
	ipfsHash = ipfsHash

	return nil
}

func GetBranchInfo(user *User, branch *ProtectBranch) error {
	// TODO
	return nil
}
