// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"strings"
)

// The orgUser table in the IPFS
type DeOrgUser struct {
	Uid      int64 `xorm:"INDEX UNIQUE(s)"`
	IsPublic bool
}

func transferOrgUserToDeOrgUser(deOrgUser *DeOrgUser, orgUser *OrgUser) {
	deOrgUser.Uid = orgUser.Uid
	deOrgUser.IsPublic = orgUser.IsPublic
}

func transferDeOrgUserToOrgUser(deOrgUser *DeOrgUser, orgUser *OrgUser, org *User) error {
	orgUser.Uid = deOrgUser.Uid
	orgUser.IsPublic = deOrgUser.IsPublic

	// TODO:
	// orgUser.IsOwner =
	// orgUser.OrgID =
	// orgUser.ID

	team := new(Team)
	total, err := x.Where("org_id = ?", org.ID).Count(team)
	if err != nil {
		return fmt.Errorf("Can not get org teams: %v", err)
	}
	orgUser.NumTeams = int(total)

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

func transferOrgToDeOrg(deOrg *DeOrg, org *User) {
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
	var err error
	var total int64

	follow := new(Follow)
	total, err = x.Where("follow_id = ?", org.ID).Count(follow)
	if err != nil {
		return fmt.Errorf("Can not get org numfollowers: %v", err)
	}
	org.NumFollowers = int(total)

	total, err = x.Where("user_id = ?", org.ID).Count(follow)
	if err != nil {
		return fmt.Errorf("Can not get org numfollowing: %v", err)
	}
	org.NumFollowing = int(total)

	// star is useless
	org.NumStars = 0

	repo := new(Repository)
	total, err = x.Where("owner_id = ?", org.ID).Count(repo)
	if err != nil {
		return fmt.Errorf("Can not get org numRepos: %v", err)
	}
	org.NumRepos = int(total)

	team := new(Team)
	total, err = x.Where("org_id = ?", org.ID).Count(team)
	if err != nil {
		return fmt.Errorf("Can not get org teams: %v", err)
	}
	org.NumTeams = int(total)

	teamUser := new(TeamUser)
	total, err = x.Where("org_id = ?", org.ID).Count(teamUser)
	if err != nil {
		return fmt.Errorf("Can not get org team_user: %v", err)
	}
	org.NumMembers = int(total)

	// TODO: not sure
	// user.UportId
	org.IsActive = true
	org.AllowGitHook = false
	org.AllowImportLocal = false
	org.ProhibitLogin = false

	return nil
}

func PushOrgInfo() error {

	return nil
}

func GetOrgInfo() error {

	return nil
}
