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

// The user table in the IPFS
type DeUser struct {
	ID                 int64
	Name               string `xorm:"UNIQUE NOT NULL"`
	FullName           string
	Email              string `xorm:"NOT NULL"`
	Passwd             string `xorm:"NOT NULL"`
	LoginType          LoginType
	LoginSource        int64 `xorm:"NOT NULL DEFAULT 0"`
	LoginName          string
	Type               UserType
	Location           string
	Website            string
	Rands              string `xorm:"VARCHAR(10)"`
	Salt               string `xorm:"VARCHAR(10)"`
	CreatedUnix        int64
	LastRepoVisibility bool
	Avatar             string `xorm:"VARCHAR(2048) NOT NULL"`
	AvatarEmail        string `xorm:"NOT NULL"`
	UseCustomAvatar    bool
	//user.email_address.email[]
	//star.repo_id[]
	//watch.repo.id[]
}

func transferUserToDeUser(deUser *DeUser, user *User) {
	deUser.ID = user.ID
	deUser.Name = user.Name
	deUser.FullName = user.FullName
	deUser.Email = user.Email
	deUser.Passwd = user.Passwd
	deUser.LoginType = user.LoginType
	deUser.LoginSource = user.LoginSource
	deUser.LoginName = user.LoginName
	deUser.Type = user.Type
	deUser.Location = user.Location
	deUser.Website = user.Website
	deUser.Rands = user.Rands
	deUser.Salt = user.Salt
	deUser.CreatedUnix = user.CreatedUnix
	deUser.LastRepoVisibility = user.LastRepoVisibility
	deUser.Avatar = user.Avatar
	deUser.AvatarEmail = user.AvatarEmail
	deUser.UseCustomAvatar = user.UseCustomAvatar
}

func deTransferUserToDeUser(deUser *DeUser, user *User) {
	user.ID = deUser.ID
	user.Name = deUser.Name
	user.FullName = deUser.FullName
	user.Email = deUser.Email
	user.Passwd = deUser.Passwd
	user.LoginType = deUser.LoginType
	user.LoginSource = deUser.LoginSource
	user.LoginName = deUser.LoginName
	user.Type = deUser.Type
	user.Location = deUser.Location
	user.Website = deUser.Website
	user.Rands = deUser.Rands
	user.Salt = deUser.Salt
	user.CreatedUnix = deUser.CreatedUnix
	user.LastRepoVisibility = deUser.LastRepoVisibility
	user.Avatar = deUser.Avatar
	user.AvatarEmail = deUser.AvatarEmail
	user.UseCustomAvatar = deUser.UseCustomAvatar

	// recovery deUser to user
	user.LowerName = strings.ToLower(user.Name)
	user.UpdatedUnix = user.CreatedUnix
	user.MaxRepoCreation = -1
	user.IsAdmin = false
	user.Description = ""
	user.NumTeams = 0
	user.NumMembers = 0

	//TODO:
	user.NumFollowers = 0
	user.NumFollowing = 0
	user.NumStars = 0
	user.NumRepos = 0

	// TODO: not sure
	//user.UportId
	user.IsActive = true
	user.AllowGitHook = false
	user.AllowImportLocal = false
	user.ProhibitLogin = false
}

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
		deUser := new(DeUser)
		transferUserToDeUser(deUser, user)
		user_data, err := json.Marshal(deUser)
		if err != nil {
			return fmt.Errorf("Can not encode user data: %v", err)
		}

		// Step 2: Put the encoded data into IPFS
		c := fmt.Sprintf("echo '%s' | ipfs add ", user_data)
		cmd := exec.Command("sh", "-c", c)
		out, err2 := cmd.Output()
		if err2 != nil {
			return fmt.Errorf("Push User to IPFS: fails: %v", err2)
		}
		ipfsHash := strings.Split(string(out), " ")[1]

		// Step3: Modify the ipfsHash in the smart contract
		// TODO: setUserInfo(ipfsHash)
		ipfsHash = ipfsHash
	}

	return nil
}

// Get the new ipfsHash from the blockchain and get the user info from IPFS
func GetUserInfo(contextUser *User) (err error) {
	// Step1: get the user info hash via addrToUserInfo
	ipfsHash := "QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN"

	// Step2: get the ipfs file and get the user data
	c := fmt.Sprintf("ipfs cat ", ipfsHash)
	cmd := exec.Command("sh", "-c", c)
	user_data, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Get User data from IPFS: fails: %v", err)
	}

	// Step3: unmarshall user data
	newDeUser := new(DeUser)
	err = json.Unmarshal(user_data, &newDeUser)
	if err != nil {
		return fmt.Errorf("Can not decode data: %v", err)
	}

	newUser := new(User)
	deTransferUserToDeUser(newDeUser, newUser)

	// Step4: remove the isAdmin column
	newUser.IsAdmin = false

	// Step5: write into the local database
	// TODO:
	CreateUser(newUser)

	return nil
}
