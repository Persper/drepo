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

// The orgUser table in the IPFS
type DeOrgUser struct {
	Uid      int64 `xorm:"INDEX UNIQUE(s)"`
	IsPublic bool
}

func transferOrgUserToDeOrgUser(orgUser *OrgUser, deOrgUser *DeOrgUser) {
	deOrgUser.Uid = orgUser.Uid
	deOrgUser.IsPublic = orgUser.IsPublic
}

func transferDeOrgUserToOrgUser(user *User, org *User, deOrgUser *DeOrgUser, orgUser *OrgUser) error {
	orgUser.Uid = deOrgUser.Uid
	orgUser.IsPublic = deOrgUser.IsPublic
	orgUser.OrgID = org.ID
	// TODO:
	// orgUser.IsOwner =
	// orgUser.ID

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
		return fmt.Errorf("Can not encode orgUser data: %v", err)
	}

	// Step 2: Put the encoded data into IPFS
	c := fmt.Sprintf("echo '%s' | ipfs add ", orgUser_data)
	cmd := exec.Command("sh", "-c", c)
	out, err2 := cmd.Output()
	if err2 != nil {
		return fmt.Errorf("Push OrgUser to IPFS: fails: %v", err2)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step4: Modify the ipfsHash in the smart contract
	// TODO: setOrgUserInfo(ipfsHash)
	ipfsHash = ipfsHash

	return nil
}

func GetOrgUserInfo() error {
	// TODO
	return nil
}

// The org table in the IPFS
type DeOrg struct {
	ID                 int64
	Name               string `xorm:"UNIQUE NOT NULL"`
	FullName           string
	Location           string
	Website            string
	Rands              string `xorm:"VARCHAR(10)"`
	Salt               string `xorm:"VARCHAR(10)"`
	CreatedUnix        int64
	LastRepoVisibility bool
	UseCustomAvatar    bool
	Description        string
}

func transferOrgToDeOrg(org *User, deOrg *DeOrg) {
	deOrg.ID = org.ID
	deOrg.Name = org.Name
	deOrg.FullName = org.FullName
	deOrg.Location = org.Location
	deOrg.Website = org.Website
	deOrg.Rands = org.Rands
	deOrg.Salt = org.Salt
	deOrg.CreatedUnix = org.CreatedUnix
	deOrg.LastRepoVisibility = org.LastRepoVisibility
	deOrg.Description = org.Description
	deOrg.UseCustomAvatar = org.UseCustomAvatar
}

func transferDeOrgToOrg(deOrg *DeOrg, org *User) error {
	org.ID = deOrg.ID
	org.Name = deOrg.Name
	org.FullName = deOrg.FullName
	org.Location = deOrg.Location
	org.Website = deOrg.Website
	org.Rands = deOrg.Rands
	org.Salt = deOrg.Salt
	org.CreatedUnix = deOrg.CreatedUnix
	org.LastRepoVisibility = deOrg.LastRepoVisibility
	org.Description = deOrg.Description
	org.UseCustomAvatar = deOrg.UseCustomAvatar

	// recovery deUser to user
	org.Type = USER_TYPE_ORGANIZATION
	org.LowerName = strings.ToLower(org.Name)
	org.UpdatedUnix = org.CreatedUnix
	org.MaxRepoCreation = -1
	org.IsAdmin = false

	// TODO: the follow and star table is lost
	// ***** START: NumFollowers *****
	follow := new(Follow)
	total, err := x.Where("follow_id = ?", org.ID).Count(follow)
	if err != nil {
		return fmt.Errorf("Can not get org numfollowers: %v", err)
	}
	org.NumFollowers = int(total)
	// ***** END: NumFollowers *****

	// ***** START: NumFollowing *****
	total, err = x.Where("user_id = ?", org.ID).Count(follow)
	if err != nil {
		return fmt.Errorf("Can not get org numfollowing: %v", err)
	}
	org.NumFollowing = int(total)
	// ***** END: NumFollowing *****

	// star is useless
	// ***** START: NumStars *****
	org.NumStars = 0
	// ***** END: NumStars *****

	// ***** START: NumRepos *****
	repo := new(Repository)
	total, err = x.Where("owner_id = ?", org.ID).Count(repo)
	if err != nil {
		return fmt.Errorf("Can not get org numRepos: %v", err)
	}
	org.NumRepos = int(total)
	// ***** END: NumRepos *****

	// ***** START: NumTeams *****
	team := new(Team)
	total, err = x.Where("org_id = ?", org.ID).Count(team)
	if err != nil {
		return fmt.Errorf("Can not get org teams: %v", err)
	}
	org.NumTeams = int(total)
	// ***** END: NumTeams *****

	// ***** START: NumMembers *****
	teamUser := new(TeamUser)
	total, err = x.Where("org_id = ?", org.ID).Count(teamUser)
	if err != nil {
		return fmt.Errorf("Can not get org team_user: %v", err)
	}
	org.NumMembers = int(total)
	// ***** END: NumMembers *****

	// TODO: not sure
	// user.UportId
	org.IsActive = true
	org.AllowGitHook = false
	org.AllowImportLocal = false
	org.ProhibitLogin = false

	return nil
}

func PushOrgInfo(user *User, org *User) error {
	if !canPushToBlockchain(user) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step 1: Encode org data into JSON format
	deOrg := new(DeOrg)
	transferOrgToDeOrg(org, deOrg)
	org_data, err := json.Marshal(deOrg)
	if err != nil {
		return fmt.Errorf("Can not encode org data: %v", err)
	}

	// Step 2: Put the encoded data into IPFS
	c := fmt.Sprintf("echo '%s' | ipfs add ", org_data)
	cmd := exec.Command("sh", "-c", c)
	out, err2 := cmd.Output()
	if err2 != nil {
		return fmt.Errorf("Push org to IPFS: fails: %v", err2)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step4: Modify the ipfsHash in the smart contract
	// TODO: setOrgUserInfo(ipfsHash)
	ipfsHash = ipfsHash

	return nil
}

func GetOrgInfo(user *User, org *User) error {
	// TODO
	return nil
}
