// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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
	UpdatedUnix        int64
	LastRepoVisibility bool
	Description        string
	UseCustomAvatar    bool
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
	deOrg.UpdatedUnix = org.UpdatedUnix
	deOrg.LastRepoVisibility = org.LastRepoVisibility
	deOrg.Description = org.Description
	deOrg.UseCustomAvatar = org.UseCustomAvatar
}

/// Prerequisite: teamUser, team, repo
func transferDeOrgToOrg(deOrg *DeOrg, org *User) error {
	org.ID = deOrg.ID
	org.Name = deOrg.Name
	org.FullName = deOrg.FullName
	org.Location = deOrg.Location
	org.Website = deOrg.Website
	org.Rands = deOrg.Rands
	org.Salt = deOrg.Salt
	org.CreatedUnix = deOrg.CreatedUnix
	org.UpdatedUnix = deOrg.UpdatedUnix
	org.LastRepoVisibility = deOrg.LastRepoVisibility
	org.Description = deOrg.Description
	org.UseCustomAvatar = deOrg.UseCustomAvatar

	// recovery deUser to user
	org.Type = USER_TYPE_ORGANIZATION
	org.LowerName = strings.ToLower(org.Name)
	org.UpdatedUnix = org.CreatedUnix
	org.MaxRepoCreation = -1
	org.IsAdmin = false

	// Not need for the organization:
	// org.Email
	// org.Passwd
	// org.LoginType
	// org.LoginSource
	// org.LoginName
	// org.Avatar
	// org.AvatarEmail

	// TODO: not sure
	org.UportId = ""
	org.IsActive = true
	org.AllowGitHook = false
	org.AllowImportLocal = false
	org.ProhibitLogin = false
	org.NumFollowers = 0
	org.NumFollowing = 0
	org.NumStars = 0

	// ***** START: NumTeams *****
	team := new(Team)
	total, err := x.Where("org_id = ?", org.ID).Count(team)
	if err != nil {
		return fmt.Errorf("Can not get org teams: %v\n", err)
	}
	org.NumTeams = int(total)
	// ***** END: NumTeams *****

	// ***** START: NumMembers *****
	teamUser := new(TeamUser)
	total, err = x.Where("org_id = ?", org.ID).Count(teamUser)
	if err != nil {
		return fmt.Errorf("Can not get org team_user: %v\n", err)
	}
	org.NumMembers = int(total)
	// ***** END: NumMembers *****

	// ***** START: NumRepos *****
	// Note: the numRepos is only used for the user rather than the organization ?
	org.NumRepos = 0
	// ***** END: NumRepos *****

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
	// TODO: setOrgInfo(ipfsHash)
	ipfsHash = ipfsHash
	fmt.Println("Push the org info to the IPFS: " + ipfsHash)

	return nil
}

func GetOrgInfo(user *User, org *User) error {
	// Step1: get the user info hash via addrToUserInfo
	ipfsHash := "QmTMU8bqRX1YcQvbe7AwjUca3U2KAFsfn3i9YwfTp1gY3C"

	// Step2: get the ipfs file and get the org data
	c := "ipfs cat " + ipfsHash
	cmd := exec.Command("sh", "-c", c)
	org_data, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Can not get org data from IPFS: %v\n", err)
	}

	// Step3: unmarshall org data
	newDeOrg := new(DeOrg)
	err = json.Unmarshal(org_data, &newDeOrg)
	if err != nil {
		return fmt.Errorf("Can not decode org data: %v\n", err)
	}

	// Step4: write into the local database and mkdir the org path
	newOrg := new(User)
	transferDeOrgToOrg(newDeOrg, newOrg)
	has, err := x.Get(newOrg)
	if err != nil {
		return fmt.Errorf("Can not search the org: %v\n", err)
	}
	if !has {
		sess := x.NewSession()
		defer sess.Close()
		if err = sess.Begin(); err != nil {
			return err
		}

		if _, err = sess.Insert(newOrg); err != nil {
			return err
		} else if err = os.MkdirAll(UserPath(newOrg.Name), os.ModePerm); err != nil {
			return err
		}

		return sess.Commit()
	}

	return nil
}

/// Push the org info and all related tables to IPFS
func PushOrgAndRelatedTables(user *User, contextOrg *User) (err error) {
	if !canPushToBlockchain(contextOrg) {
		return fmt.Errorf("The org can not push to the blockchain")
	}

	// Step0: get the corresponding user.
	var org *User
	org = &User{ID: contextOrg.ID}
	_, err = x.Get(org)
	if err != nil {
		return fmt.Errorf("Can not get org data: %v", err)
	}

	// Step1: push org: check update or create
	if err := PushUserInfo(org, 1); err != nil {
		return fmt.Errorf("Can not push orgInfo: %v", err)
	}

	// Step2: push orgUser
	orgUsers := make([]OrgUser, 0)
	if err = x.Find(&orgUsers, &OrgUser{Uid: org.ID}); err != nil {
		return fmt.Errorf("Can not get orgUsers of the user: %v", err)
	}
	for i := range orgUsers {
		if err = PushOrgUserInfo(user, org, &orgUsers[i]); err != nil {
			return fmt.Errorf("Can not push orgUser data: %v", err)
		}
	}

	// Step3: push team
	teams := make([]Team, 0)
	if err = x.Find(&teams, &Team{OrgID: org.ID}); err != nil {
		return fmt.Errorf("Can not get teams of the user: %v", err)
	}
	for j := range teams {
		if err = PushTeamInfo(org, &teams[j]); err != nil {
			return fmt.Errorf("Can not push team data: %v", err)
		}
	}

	// Step4: push repo
	repos := make([]Repository, 0)
	if err = x.Find(&repos, &Repository{OwnerID: org.ID}); err != nil {
		return fmt.Errorf("Can not get owned repos of the org: %v", err)
	}

	for i := range repos {
		if err = PushRepoInfo(org, &repos[i]); err != nil {
			return fmt.Errorf("Can not push repo data: %v", err)
		}

		issues := make([]Issue, 0)
		if err = x.Find(&issues, &Issue{RepoID: repos[i].ID}); err != nil {
			return fmt.Errorf("Can not get issues of the repo: %v", err)
		}
		for j := range issues {
			if err = PushIssueInfo(org, &issues[j]); err != nil {
				return fmt.Errorf("Can not push issue data: %v", err)
			}

			pulls := make([]PullRequest, 0)
			if err = x.Find(&pulls, &PullRequest{IssueID: issues[j].ID}); err != nil {
				return fmt.Errorf("Can not get pulls of the repo: %v", err)
			}

			for k := range pulls {
				if err = PushPullInfo(org, &pulls[k]); err != nil {
					return fmt.Errorf("Can not push pull_request data: %v", err)
				}
			}
		}

		branches := make([]ProtectBranch, 0)
		if err = x.Find(&branches, &ProtectBranch{RepoID: repos[i].ID}); err != nil {
			return fmt.Errorf("Can not get branches of the repo: %v", err)
		}
		for j := range branches {
			if err = PushBranchInfo(org, &branches[j]); err != nil {
				return fmt.Errorf("Can not push branch data: %v", err)
			}
		}
	}

	return nil
}
