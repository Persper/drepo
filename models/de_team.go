// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import ()

// The Team table in the IPFS
type DeTeam struct {
	ID          int64
	LowerName   string
	Name        string
	Description string
	Authorize   AccessMode
	// team_repo.repo_id[]
	// team_user.uid[]
}

func transferTeamToDeTeam(deTeam *DeTeam, team *Team) {
	deTeam.ID = team.ID
	deTeam.LowerName = team.LowerName
	deTeam.Name = team.Name
	deTeam.Description = team.Description
	deTeam.Authorize = team.Authorize
}

func transferDeTeamToTeam(deTeam *DeTeam, team *Team, org *User) error {
	team.ID = deTeam.ID
	team.LowerName = deTeam.LowerName
	team.Name = deTeam.Name
	team.Description = deTeam.Description
	team.Authorize = deTeam.Authorize

	// Restore
	team.OrgID = org.ID

	// TODO:
	// team.NumRepos
	// team.NumMembers

	return nil
}
