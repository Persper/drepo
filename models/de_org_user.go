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

type DeOrgUser struct {
	Uid      int64 `xorm:"INDEX UNIQUE(s)"`
	IsPublic bool
}

func transferOrgUserToDeOrgUser(orgUser *OrgUser, deOrgUser *DeOrgUser) {
	deOrgUser.Uid = orgUser.Uid
	deOrgUser.IsPublic = orgUser.IsPublic
}

/// Prerequisite: team, teamUser
func transferDeOrgUserToOrgUser(user *User, org *User, deOrgUser *DeOrgUser, orgUser *OrgUser) error {
	// orgUser.ID can be generated at any time
	// TODO: orgUser.ID
	orgUser.Uid = deOrgUser.Uid
	orgUser.IsPublic = deOrgUser.IsPublic
	orgUser.OrgID = org.ID

	// TODO: restore from team_user and team?
	// orgUser.IsOwner

	// ***** START: NumTeams *****
	team := new(Team)
	total, err := x.Where("org_id = ?", org.ID).Count(team)
	if err != nil {
		return fmt.Errorf("Can not get org teams: %v", err)
	}
	orgUser.NumTeams = int(total)
	// ***** END: NumTeams *****

	return nil
}

func PushOrgUserInfo(user *User, org *User, orgUser *OrgUser) error {
	if !canPushToBlockchain(user) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step 1: Encode orgUser data into JSON format
	deOrgUser := new(DeOrgUser)
	transferOrgUserToDeOrgUser(orgUser, deOrgUser)
	orgUser_data, err := json.Marshal(deOrgUser)
	if err != nil {
		return fmt.Errorf("Can not encode orgUser data: %v\n", err)
	}

	// Step 2: Put the encoded data into IPFS
	c := fmt.Sprintf("echo '%s' | ipfs add ", orgUser_data)
	cmd := exec.Command("sh", "-c", c)
	out, err2 := cmd.Output()
	if err2 != nil {
		return fmt.Errorf("Can not push orgUser to IPFS: %v\n", err2)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step4: Modify the ipfsHash in the smart contract
	// TODO: setOrgUserInfo(ipfsHash)
	ipfsHash = ipfsHash
	fmt.Println("Push the orgUser to the IPFS: " + ipfsHash)

	return nil
}

func GetOrgUserInfo(user *User, org *User, orgUser *OrgUser) error {
	// Step1: get the issue info hash
	ipfsHash := "QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN"

	// Step2: get the ipfs file and get the pull_request data
	c := "ipfs cat " + ipfsHash
	cmd := exec.Command("sh", "-c", c)
	orgUser_data, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Get org_user data from IPFS: fails: %v\n", err)
	}

	// Step3: unmarshall pull_request data
	newDeOU := new(DeOrgUser)
	err = json.Unmarshal(orgUser_data, &newDeOU)
	if err != nil {
		return fmt.Errorf("Can not decode data: %v\n", err)
	}

	// Step4: write into the local database and mkdir the local path
	newOU := new(OrgUser)
	transferDeOrgUserToOrgUser(user, org, newDeOU, newOU)
	has, err2 := x.Get(newOU)
	if err2 != nil {
		return fmt.Errorf("Can not search the org_user: %v\n", err2)
	}
	if !has {
		_, err = x.Insert(newOU)
		if err != nil {
			return fmt.Errorf("Can not add the org_user to the server: %v\n", err)
		}
	}

	return nil
}
