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

type DeLabel struct {
	ID    int64
	Name  string
	Color string `xorm:"VARCHAR(7)"`
}

func transferLabelToDeLabel(label *Label, deLabel *DeLabel) {
	deLabel.ID = label.ID
	deLabel.Name = label.Name
	deLabel.Color = label.Color
}

// Prerequisite: issue / issue_label
func transferDeLabelToLabel(repo *Repository, deLabel *DeLabel, label *Label) error {
	label.ID = deLabel.ID
	label.Name = deLabel.Name
	label.Color = deLabel.Color
	label.RepoID = repo.ID

	// ***** START: NumIssues and NumClosedIssues *****
	issueLabels := make([]IssueLabel, 0)
	if err := x.Find(&issueLabels, &IssueLabel{LabelID: label.ID}); err != nil {
		return fmt.Errorf("Can not get issue_label of the label: %v\n", err)
	}
	label.NumIssues = len(issueLabels)
	label.NumClosedIssues = 0
	for i := range issueLabels {
		issue := &Issue{ID: issueLabels[i].IssueID, IsClosed: true}
		has, err := x.Get(issue)
		if err != nil {
			return fmt.Errorf("Can not search the issue: %v\n", err)
		}
		if has {
			label.NumClosedIssues = label.NumClosedIssues + 1
		}
	}
	// ***** END: NumIssues and NumClosedIssues *****

	return nil
}

func PushLabelInfo(user *User, label *Label) error {
	if !canPushToBlockchain(user) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step1: encode label data into JSON format
	deLabel := new(DeLabel)
	transferLabelToDeLabel(label, deLabel)
	label_data, err := json.Marshal(deLabel)
	if err != nil {
		return fmt.Errorf("Can not encode label data: %v\n", err)
	}

	// Step2: push the encoded data into IPFS
	c := fmt.Sprintf("echo '%s' | ipfs add ", label_data)
	cmd := exec.Command("sh", "-c", c)
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Push label to IPFS: %v\n", err)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step3: modify the ipfsHash in the smart contract
	// TODO: setLabelInfo(ipfsHash)
	ipfsHash = ipfsHash
	fmt.Println("Push the label file to the IPFS: " + ipfsHash)

	return nil
}

func GetLabelInfo(user *User, repo *Repository, ipfsHash string) error {
	// Step1: get the ipfs file and get the label data
	c := "ipfs cat " + ipfsHash
	cmd := exec.Command("sh", "-c", c)
	label_data, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Get label data from IPFS: %v\n", err)
	}

	// Step2: unmarshall pull data
	newDeLabel := new(DeLabel)
	err = json.Unmarshal(label_data, &newDeLabel)
	if err != nil {
		return fmt.Errorf("Can not decode data: %v\n", err)
	}

	// Step3: write into the local database
	newLabel := new(Label)
	transferDeLabelToLabel(repo, newDeLabel, newLabel)
	has, err := x.Get(newLabel)
	if err != nil {
		return fmt.Errorf("Can not search the label: %v\n", err)
	}
	if !has {
		_, err = x.Insert(newLabel)
		if err != nil {
			return fmt.Errorf("Can not add the label request to the server: %v\n", err)
		}
	}

	return nil
}
