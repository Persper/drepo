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

// DeEmailAdresses is the list of all email addresses of a user that remove is_primary
type DeEmailAddress struct {
	Email string `xorm:"UNIQUE NOT NULL"`
}

func transferEmailAddrToDeEmailAddr(emailAddr *EmailAddress, deEmailAddr *DeEmailAddress) {
	deEmailAddr.Email = emailAddr.Email
}

func transferDeEmailAddrToEmailAddr(user *User, deEmailAddr *DeEmailAddress, emailAddr *EmailAddress) {
	//emailAddr.ID
	emailAddr.UID = user.ID
	emailAddr.Email = deEmailAddr.Email
	emailAddr.IsActivated = true
	if emailAddr.Email == user.Email {
		emailAddr.IsPrimary = true
	} else {
		emailAddr.IsPrimary = false
	}
}

// DePublicKey represents a user or deploy SSH public key.
type DePublicKey struct {
	Name        string     `xorm:"NOT NULL"`
	Fingerprint string     `xorm:"NOT NULL"`
	Content     string     `xorm:"TEXT NOT NULL"`
	Mode        AccessMode `xorm:"NOT NULL DEFAULT 2"`
	Type        KeyType    `xorm:"NOT NULL DEFAULT 1"`
	CreatedUnix int64
	UpdatedUnix int64
}

func transferPubKeyToDePubKey(pubKey *PublicKey, dePubKey *DePublicKey) {
	dePubKey.Name = pubKey.Name
	dePubKey.Fingerprint = pubKey.Fingerprint
	dePubKey.Content = pubKey.Content
	dePubKey.Mode = pubKey.Mode
	dePubKey.Type = pubKey.Type
	dePubKey.CreatedUnix = pubKey.CreatedUnix
	dePubKey.UpdatedUnix = pubKey.UpdatedUnix
}

func transferDePubKeyToPubKey(user *User, dePubKey *DePublicKey, pubKey *PublicKey) {
	//pubKey.ID
	pubKey.OwnerID = user.ID
	pubKey.Name = dePubKey.Name
	pubKey.Fingerprint = dePubKey.Fingerprint
	pubKey.Content = dePubKey.Content
	pubKey.Mode = dePubKey.Mode
	pubKey.Type = dePubKey.Type
	pubKey.CreatedUnix = dePubKey.CreatedUnix
	pubKey.UpdatedUnix = dePubKey.UpdatedUnix
}

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
	Location           string
	Website            string
	Rands              string `xorm:"VARCHAR(10)"`
	Salt               string `xorm:"VARCHAR(10)"`
	CreatedUnix        int64
	UpdatedUnix        int64
	LastRepoVisibility bool
	Avatar             string `xorm:"VARCHAR(2048) NOT NULL"`
	AvatarEmail        string `xorm:"NOT NULL"`
	UseCustomAvatar    bool
	EmailAddr          []DeEmailAddress `xorm:"-"`
	PubKey             []DePublicKey    `xorm:"-"`
	//star.repo_id[]
	//watch.repo.id[]
	//repo_blacklist
	//team_blacklist
	//org_blacklist
}

func transferUserToDeUser(user *User, deUser *DeUser) error {
	deUser.ID = user.ID
	deUser.Name = user.Name
	deUser.FullName = user.FullName
	deUser.Email = user.Email
	deUser.Passwd = user.Passwd
	deUser.LoginType = user.LoginType
	deUser.LoginSource = user.LoginSource
	deUser.LoginName = user.LoginName
	deUser.Location = user.Location
	deUser.Website = user.Website
	deUser.Rands = user.Rands
	deUser.Salt = user.Salt
	deUser.CreatedUnix = user.CreatedUnix
	deUser.UpdatedUnix = user.UpdatedUnix
	deUser.LastRepoVisibility = user.LastRepoVisibility
	deUser.Avatar = user.Avatar
	deUser.AvatarEmail = user.AvatarEmail
	deUser.UseCustomAvatar = user.UseCustomAvatar

	// ***** START: EmailAddress[] *****
	emailAddresses := make([]EmailAddress, 0)
	if err := x.Find(&emailAddresses, &EmailAddress{UID: user.ID}); err != nil {
		return fmt.Errorf("Can not get emailAddress of the user: %v", err)
	}
	for i := range emailAddresses {
		deEmailAddr := new(DeEmailAddress)
		transferEmailAddrToDeEmailAddr(&emailAddresses[i], deEmailAddr)
		deUser.EmailAddr = append(deUser.EmailAddr, *deEmailAddr)
	}
	// ***** END: EmailAddress[] *****

	// ***** START: PubKey[] *****
	publicKeys := make([]PublicKey, 0)
	if err := x.Find(&publicKeys, &PublicKey{OwnerID: user.ID}); err != nil {
		return fmt.Errorf("Can not get publicKey of the user: %v", err)
	}
	for i := range publicKeys {
		dePublicKey := new(DePublicKey)
		transferPubKeyToDePubKey(&publicKeys[i], dePublicKey)
		deUser.PubKey = append(deUser.PubKey, *dePublicKey)
	}
	// ***** END: PubKey[] *****

	return nil
}

func transferDeUserToUser(deUser *DeUser, user *User) error {
	user.ID = deUser.ID
	user.Name = deUser.Name
	user.FullName = deUser.FullName
	user.Email = deUser.Email
	user.Passwd = deUser.Passwd
	user.LoginType = deUser.LoginType
	user.LoginSource = deUser.LoginSource
	user.LoginName = deUser.LoginName
	user.Location = deUser.Location
	user.Website = deUser.Website
	user.Rands = deUser.Rands
	user.Salt = deUser.Salt
	user.CreatedUnix = deUser.CreatedUnix
	user.UpdatedUnix = deUser.UpdatedUnix
	user.LastRepoVisibility = deUser.LastRepoVisibility
	user.Avatar = deUser.Avatar
	user.AvatarEmail = deUser.AvatarEmail
	user.UseCustomAvatar = deUser.UseCustomAvatar

	// recovery deUser to user
	user.Type = USER_TYPE_INDIVIDUAL
	user.LowerName = strings.ToLower(user.Name)
	user.MaxRepoCreation = -1
	user.IsAdmin = false

	// org
	user.Description = ""
	user.NumTeams = 0
	user.NumMembers = 0

	// TODO: not sure
	// user.UportId
	user.IsActive = true
	user.AllowGitHook = false
	user.AllowImportLocal = false
	user.ProhibitLogin = false

	// TODO: the watch and star table is lost
	// ***** START: NumFollowers *****
	follow := new(Follow)
	total, err := x.Where("follow_id = ?", user.ID).Count(follow)
	if err != nil {
		return fmt.Errorf("Can not get user numfollowers: %v", err)
	}
	user.NumFollowers = int(total)
	// ***** END: NumFollowers *****

	// ***** START: NumFollowing *****
	total, err = x.Where("user_id = ?", user.ID).Count(follow)
	if err != nil {
		return fmt.Errorf("Can not get user numfollowing: %v", err)
	}
	user.NumFollowing = int(total)
	// ***** END: NumFollowing *****

	// ***** START: NumStars *****
	user.NumStars = 0
	// ***** END: NumStars *****

	// ***** START: NumRepos *****
	repo := new(Repository)
	total, err = x.Where("owner_id = ?", user.ID).Count(repo)
	if err != nil {
		return fmt.Errorf("Can not get user numRepos: %v", err)
	}
	user.NumRepos = int(total)
	// ***** END: NumRepos *****

	return nil
}

/// Push the user info to IPFS and record the new ipfsHash in the blockchain
/// pushMode: 0 - register; 1 - update; 2 - delete;
func PushUserInfo(user *User, pushMode int) (err error) {
	// Do some checks
	if user.IsOrganization() {
		return nil
	}
	if !canPushToBlockchain(user) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step1: register/deregister the user if it does not exist
	if pushMode == 1 {
		//err = registerName
	}
	if pushMode == 2 {
		//err = deregisterName
	}

	// Step 2: Encode user data into JSON format
	deUser := new(DeUser)
	transferUserToDeUser(user, deUser)
	user_data, err := json.Marshal(deUser)
	if err != nil {
		return fmt.Errorf("Can not encode user data: %v", err)
	}

	// Step 3: Put the encoded data into IPFS
	c := fmt.Sprintf("echo '%s' | ipfs add ", user_data)
	cmd := exec.Command("sh", "-c", c)
	out, err2 := cmd.Output()
	if err2 != nil {
		return fmt.Errorf("Push User to IPFS: fails: %v", err2)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step4: Modify the ipfsHash in the smart contract
	// TODO: setUserInfo(ipfsHash)
	ipfsHash = ipfsHash

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
	transferDeUserToUser(newDeUser, newUser)

	// Step4: write into the local database
	// TODO:
	CreateUser(newUser)

	return nil
}

/// Push the user info and all related tables to IPFS
func PushUserAllInfos(contextUser *User) (err error) {
	if !canPushToBlockchain(contextUser) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step0: get the corresponding user.
	var user *User
	user = &User{ID: contextUser.ID}
	_, err = x.Get(user)
	if err != nil {
		return fmt.Errorf("Can not get user data: %v", err)
	}

	// Step1: push user: check update or create
	if err := PushUserInfo(user, 1); err != nil {
		return fmt.Errorf("Can not push userInfo: %v", err)
	}

	// Step2: push the related orgs
	orgUsers := make([]OrgUser, 0)
	if err = x.Find(&orgUsers, &OrgUser{Uid: user.ID}); err != nil {
		return fmt.Errorf("Can not get orgUsers of the user: %v", err)
	}
	for i := range orgUsers {
		var org *User
		org = &User{ID: orgUsers[i].ID}
		hasOrg, err := x.Get(org)
		if hasOrg {
			if err != nil {
				return fmt.Errorf("Can not get org data: %v", err)
			}
			if err = PushOrgInfo(user, org); err != nil {
				return fmt.Errorf("Can not push org data: %v", err)
			}
			// TODO: only owner?
			if err = PushOrgUserInfo(user, org, &orgUsers[i]); err != nil {
				return fmt.Errorf("Can not push orgUser data: %v", err)
			}

			teams := make([]Team, 0)
			if err = x.Find(&teams, &Team{OrgID: org.ID}); err != nil {
				return fmt.Errorf("Can not get teams of the user: %v", err)
			}

			for j := range teams {
				if err = PushTeamInfo(user, &teams[j]); err != nil {
					return fmt.Errorf("Can not push team data: %v", err)
				}
			}
		}
	}

	return nil
}
