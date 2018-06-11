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

// The pull request table in the IPFS
type DePull struct {
	Type   PullRequestType
	Status PullRequestStatus

	IssueID int64 `xorm:"INDEX"`
	Index   int64

	HeadRepoID   int64
	HeadUserName string
	HeadBranch   string
	BaseBranch   string
	MergeBase    string `xorm:"VARCHAR(40)"`

	HasMerged      bool
	MergedCommitID string `xorm:"VARCHAR(40)"`
	MergerID       int64
	MergedUnix     int64
}

func transferPullToDePull(pull *PullRequest, dePull *DePull) {
	dePull.Type = pull.Type
	dePull.Status = pull.Status
	dePull.IssueID = pull.IssueID
	dePull.Index = pull.Index
	dePull.HeadRepoID = pull.HeadRepoID
	dePull.HeadUserName = pull.HeadUserName
	dePull.HeadBranch = pull.HeadBranch
	dePull.BaseBranch = pull.BaseBranch
	dePull.MergeBase = pull.MergeBase
	dePull.HasMerged = pull.HasMerged
	dePull.MergedCommitID = pull.MergedCommitID
	dePull.MergerID = pull.MergerID
	dePull.MergedUnix = pull.MergedUnix
}

func transferDePullToPull(dePull *DePull, pull *PullRequest) error {
	// pull.ID =
	// pull.BaseRepoID =
	pull.Type = dePull.Type
	pull.Status = dePull.Status
	pull.IssueID = dePull.IssueID
	pull.Index = dePull.Index
	pull.HeadRepoID = dePull.HeadRepoID
	pull.HeadUserName = dePull.HeadUserName
	pull.HeadBranch = dePull.HeadBranch
	pull.BaseBranch = dePull.BaseBranch
	pull.MergeBase = dePull.MergeBase
	pull.HasMerged = dePull.HasMerged
	pull.MergedCommitID = dePull.MergedCommitID
	pull.MergerID = dePull.MergerID
	pull.MergedUnix = dePull.MergedUnix

	return nil
}

func PushPullInfo(user *User, pull *PullRequest) error {
	if !canPushToBlockchain(user) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step 1: Encode org data into JSON format
	dePull := new(DePull)
	transferPullToDePull(pull, dePull)
	pull_data, err := json.Marshal(dePull)
	if err != nil {
		return fmt.Errorf("Can not encode pull_request data: %v", err)
	}

	// Step 2: Put the encoded data into IPFS
	c := fmt.Sprintf("echo '%s' | ipfs add ", pull_data)
	cmd := exec.Command("sh", "-c", c)
	out, err2 := cmd.Output()
	if err2 != nil {
		return fmt.Errorf("Push pull_request to IPFS: fails: %v", err2)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step4: Modify the ipfsHash in the smart contract
	// TODO: setPullInfo(ipfsHash)
	ipfsHash = ipfsHash

	return nil
}

func GetPullInfo(user *User, pull *PullRequest) error {
	// TODO
	return nil
}
