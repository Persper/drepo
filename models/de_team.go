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

/// ***** START: DeTeamRepo *****
type DeTeamRepo struct {
	RepoID int64 `xorm:"UNIQUE(s)"`
}

func transferTeamRepoToDeTeamRepo(teamRepo *TeamRepo, deTeamRepo *DeTeamRepo) {
	deTeamRepo.RepoID = teamRepo.RepoID
}

func transferDeTeamRepoToTeamRepo(org *User, team *Team, deTeamRepo *DeTeamRepo, teamRepo *TeamRepo) {
	// teamRepo.ID can be generated at any time
	// TODO: teamRepo.ID
	teamRepo.OrgID = org.ID
	teamRepo.TeamID = team.ID
	teamRepo.RepoID = deTeamRepo.RepoID
}

/// ***** END: DeTeamRepo *****

/// ***** START: DeTeamUser *****
type DeTeamUser struct {
	UID int64 `xorm:"UNIQUE(s)"`
}

func transferTeamUserToDeTeamUser(teamUser *TeamUser, deTeamUser *DeTeamUser) {
	deTeamUser.UID = teamUser.UID
}

func transferDeTeamUserToTeamUser(org *User, team *Team, deTeamUser *DeTeamUser, teamUser *TeamUser) {
	// teamUser.ID can be generated at any time
	// TODO: teamUser.ID
	teamUser.OrgID = org.ID
	teamUser.TeamID = team.ID
	teamUser.UID = deTeamUser.UID
}

/// ***** END: DeTeamUser *****

type DeTeam struct {
	ID          int64
	LowerName   string
	Name        string
	Description string
	Authorize   AccessMode
	DeTeamRepos []DeTeamRepo `xorm:"-"`
	DeTeamUsers []DeTeamUser `xorm:"-"`
}

func transferTeamToDeTeam(team *Team, deTeam *DeTeam) error {
	deTeam.ID = team.ID
	deTeam.LowerName = team.LowerName
	deTeam.Name = team.Name
	deTeam.Description = team.Description
	deTeam.Authorize = team.Authorize

	// ***** START: TeamRepoID[] *****
	if err := x.Table("team_repo").Cols("repo_id").Find(&deTeam.DeTeamRepos); err != nil {
		return fmt.Errorf("Can not encode team_repo data: %v\n", err)
	}
	// ***** END: TeamRepoID[] *****

	// ***** START: TeamUserID[] *****
	if err := x.Table("team_user").Cols("uid").Find(&deTeam.DeTeamUsers); err != nil {
		return fmt.Errorf("Can not encode team_user data: %v\n", err)
	}
	// ***** END: TeamUserID[] *****

	return nil
}

/// Prerequisite: team, teamUser
func transferDeTeamToTeam(org *User, deTeam *DeTeam, team *Team) error {
	team.ID = deTeam.ID
	team.LowerName = deTeam.LowerName
	team.Name = deTeam.Name
	team.Description = deTeam.Description
	team.Authorize = deTeam.Authorize

	team.OrgID = org.ID

	// ***** START: NumRepos[] *****
	for i := range deTeam.DeTeamRepos {
		teamRepo := new(TeamRepo)
		transferDeTeamRepoToTeamRepo(org, team, &deTeam.DeTeamRepos[i], teamRepo)
		has, err := x.Get(teamRepo)
		if err != nil {
			return fmt.Errorf("Can not search the teamRepo: %v\n", err)
		}
		if !has {
			_, err = x.Insert(teamRepo)
			if err != nil {
				return fmt.Errorf("Can not add the teamRepo to the server: %v\n", err)
			}
		}
	}
	team.NumRepos = len(deTeam.DeTeamRepos)
	// ***** END: NumRepos[] *****

	// ***** START: NumMembers[] *****
	for i := range deTeam.DeTeamUsers {
		teamUser := new(TeamUser)
		transferDeTeamUserToTeamUser(org, team, &deTeam.DeTeamUsers[i], teamUser)
		has, err := x.Get(teamUser)
		if err != nil {
			return fmt.Errorf("Can not search the teamUser: %v\n", err)
		}
		if !has {
			_, err = x.Insert(teamUser)
			if err != nil {
				return fmt.Errorf("Can not add the teamUser to the server: %v\n", err)
			}
		}
	}
	team.NumMembers = len(deTeam.DeTeamUsers)
	// ***** END: NumMembers[] *****

	return nil
}

func PushTeamInfo(org *User, team *Team) error {
	/*if !canPushToBlockchain(org) {
		return fmt.Errorf("The user can not push to the blockchain")
	}*/

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
	fmt.Println("Push the team to the IPFS: " + ipfsHash)

	return nil
}

func GetTeamInfo(org *User, team *Team) error {
	// Step1: get the team info hash
	ipfsHash := "QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN"

	// Step2: get the ipfs file and get the team data
	c := "ipfs cat " + ipfsHash
	cmd := exec.Command("sh", "-c", c)
	team_data, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Get team data from IPFS: fails: %v\n", err)
	}

	// Step3: unmarshall team data
	newDeTeam := new(DeTeam)
	err = json.Unmarshal(team_data, &newDeTeam)
	if err != nil {
		return fmt.Errorf("Can not decode data: %v\n", err)
	}

	// Step4: write into the local database
	newTeam := new(Team)
	transferDeTeamToTeam(org, newDeTeam, newTeam)
	has, err2 := x.Get(newTeam)
	if err2 != nil {
		return fmt.Errorf("Can not search the team: %v\n", err2)
	}
	if !has {
		_, err = x.Insert(newTeam)
		if err != nil {
			return fmt.Errorf("Can not add the newTeam to the server: %v\n", err)
		}
	}

	return nil
}
