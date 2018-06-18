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

type DePullRequest struct {
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

func transferPullToDePull(pr *PullRequest, dePr *DePullRequest) {
	dePr.Type = pr.Type
	dePr.Status = pr.Status
	dePr.IssueID = pr.IssueID
	dePr.Index = pr.Index
	dePr.HeadRepoID = pr.HeadRepoID
	dePr.HeadUserName = pr.HeadUserName
	dePr.HeadBranch = pr.HeadBranch
	dePr.BaseBranch = pr.BaseBranch
	dePr.MergeBase = pr.MergeBase
	dePr.HasMerged = pr.HasMerged
	dePr.MergedCommitID = pr.MergedCommitID
	dePr.MergerID = pr.MergerID
	dePr.MergedUnix = pr.MergedUnix
}

func transferDePullToPull(repo *Repository, dePr *DePullRequest, pr *PullRequest) error {
	// pr.ID can be generated at any time
	// TODO: pr.ID
	pr.BaseRepoID = repo.ID
	pr.Type = dePr.Type
	pr.Status = dePr.Status
	pr.IssueID = dePr.IssueID
	pr.Index = dePr.Index
	pr.HeadRepoID = dePr.HeadRepoID
	pr.HeadUserName = dePr.HeadUserName
	pr.HeadBranch = dePr.HeadBranch
	pr.BaseBranch = dePr.BaseBranch
	pr.MergeBase = dePr.MergeBase
	pr.HasMerged = dePr.HasMerged
	pr.MergedCommitID = dePr.MergedCommitID
	pr.MergerID = dePr.MergerID
	pr.MergedUnix = dePr.MergedUnix

	return nil
}

func PushPullInfo(user *User, pr *PullRequest) error {
	if !canPushToBlockchain(user) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step1: Encode org data into JSON format
	dePr := new(DePullRequest)
	transferPullToDePull(pr, dePr)
	pr_data, err := json.Marshal(dePr)
	if err != nil {
		return fmt.Errorf("Can not encode pull_request data: %v\n", err)
	}

	// Step2: Put the encoded data into IPFS
	c := fmt.Sprintf("echo '%s' | ipfs add ", pr_data)
	cmd := exec.Command("sh", "-c", c)
	out, err2 := cmd.Output()
	if err2 != nil {
		return fmt.Errorf("Push pull_request to IPFS: fails: %v\n", err2)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step3: Modify the ipfsHash in the smart contract
	// TODO: setPullInfo(ipfsHash)
	ipfsHash = ipfsHash
	fmt.Println("Push the pull_request file to the IPFS: " + ipfsHash)

	return nil
}

func GetPullInfo(user *User, repo *Repository, ipfsHash string) error {
	// Step1: get the issue info hash
	//ipfsHash := "QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN"

	// Step2: get the ipfs file and get the pull_request data
	c := "ipfs cat " + ipfsHash
	cmd := exec.Command("sh", "-c", c)
	pr_data, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Get pull_request data from IPFS: fails: %v\n", err)
	}

	// Step3: unmarshall pull_request data
	newDePr := new(DePullRequest)
	err = json.Unmarshal(pr_data, &newDePr)
	if err != nil {
		return fmt.Errorf("Can not decode data: %v", err)
	}

	// Step4: write into the local database
	newPr := new(PullRequest)
	transferDePullToPull(repo, newDePr, newPr)
	has, err2 := x.Get(newPr)
	if err2 != nil {
		return fmt.Errorf("Can not search the pull request: %v\n", err2)
	}
	if !has {
		_, err = x.Insert(newPr)
		if err != nil {
			return fmt.Errorf("Can not add the pull request to the server: %v\n", err)
		}
	}

	return nil
}
