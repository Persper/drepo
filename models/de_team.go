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

type DeTeamRepo struct {
	RepoID int64 `xorm:"UNIQUE(s)"`
}

func transferTeamRepoToDeTeamRepo(teamRepo *TeamRepo, deTeamRepo *DeTeamRepo) {
	deTeamRepo.RepoID = teamRepo.RepoID
}

func transferDeTeamRepoToTeamRepo(org *User, team *Team, repo *Repository, deTeamRepo *DeTeamRepo, teamRepo *TeamRepo) {
	// TODO: teamRepo.ID
	teamRepo.OrgID = org.ID
	teamRepo.TeamID = team.ID
	teamRepo.RepoID = repo.ID
}

type DeTeamUser struct {
	UID int64 `xorm:"UNIQUE(s)"`
}

func transferTeamUserToDeTeamUser(teamUser *TeamUser, deTeamUser *DeTeamUser) {
	deTeamUser.UID = teamUser.UID
}

func transferDeTeamUserToTeamUser(org *User, team *Team, user *User, teamUser *TeamUser, deTeamUser *DeTeamUser) {
	// TODO: teamUser.ID
	teamUser.OrgID = org.ID
	teamUser.TeamID = team.ID
	teamUser.UID = user.ID
}

// The Team table in the IPFS
type DeTeam struct {
	ID          int64
	LowerName   string
	Name        string
	Description string
	Authorize   AccessMode
	TeamRepoID  []int64 `xorm:"-"`
	TeamUserID  []int64 `xorm:"-"`
}

func transferTeamToDeTeam(team *Team, deTeam *DeTeam) error {
	deTeam.ID = team.ID
	deTeam.LowerName = team.LowerName
	deTeam.Name = team.Name
	deTeam.Description = team.Description
	deTeam.Authorize = team.Authorize

	if err := x.Table("team_repo").Cols("repo_id").Find(&deTeam.TeamRepoID); err != nil {
		return fmt.Errorf("Can not encode team_repo data: %v", err)
	}

	if err := x.Table("team_user").Cols("uid").Find(&deTeam.TeamUserID); err != nil {
		return fmt.Errorf("Can not encode team_user data: %v", err)
	}

	return nil
}

func transferDeTeamToTeam(org *User, deTeam *DeTeam, team *Team) error {
	team.OrgID = org.ID
	team.ID = deTeam.ID
	team.LowerName = deTeam.LowerName
	team.Name = deTeam.Name
	team.Description = deTeam.Description
	team.Authorize = deTeam.Authorize

	// TODO:
	// team.NumRepos
	// team.NumMembers

	// TODO:
	// Restore teamUser/teamRepo

	return nil
}

func PushTeamInfo(user *User, team *Team) error {
	if !canPushToBlockchain(user) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step 1: Encode team data into JSON format
	deTeam := new(DeTeam)
	transferTeamToDeTeam(team, deTeam)
	team_data, err := json.Marshal(deTeam)
	if err != nil {
		return fmt.Errorf("Can not encode team data: %v", err)
	}

	// Step 2: Put the encoded data into IPFS
	c := fmt.Sprintf("echo '%s' | ipfs add ", team_data)
	cmd := exec.Command("sh", "-c", c)
	out, err2 := cmd.Output()
	if err2 != nil {
		return fmt.Errorf("Push org to IPFS: fails: %v", err2)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step4: Modify the ipfsHash in the smart contract
	// TODO: setTeamInfo(ipfsHash)
	ipfsHash = ipfsHash

	return nil
}

func GetTeamInfo() error {
	// TODO
	return nil
}
