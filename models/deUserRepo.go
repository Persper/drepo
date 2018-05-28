// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	//"github.com/go-xorm/xorm"
	"github.com/gogs/gogs/models/ipfs"
)

func PushUserRepoInfo(uportid string) {
	/* User: get the corresponding user.  */
	var user *User
	user = &User{UportId: uportid}
	hasUser, err := x.Get(user)
	if err != nil {
		fmt.Println("Error: get user!")
	}

	if hasUser {
		/* Access: get the access control between user and repo. */
		accesses := make([]Access, 0)
		err = x.Where("user_id = ?", user.ID).Find(&accesses)
		if err != nil {
			fmt.Println("Access Error!")
		}

		/* TeamUser: get the relationship between user and team. */
		/* TeamUser: org, team and user. */
		teamUsers := make([]TeamUser, 0)
		err = x.Where("uid = ?", user.ID).Find(&teamUsers)
		if err != nil {
			fmt.Println("TeamUser Error!")
		}

		/* OrgUser: get the relationship between user and organization. */
		orgUsers := make([]OrgUser, 0)
		err = x.Where("uid = ?", user.ID).Find(&orgUsers)
		if err != nil {
			fmt.Println("OrgUser Error!")
		}

		/* Collaboration: get the collaboration control between user and repo. */
		collaborations := make([]Collaboration, 0)
		err = x.Where("user_id = ?", user.ID).Find(&collaborations)
		if err != nil {
			fmt.Println("Collaboration Error!")
		}

		/* Repository: get all repos related to the user. */
		repoNum := len(accesses)
		repos := make([]Repository, 0)
		for i := 0; i < repoNum; i++ {
			repoId := accesses[i].RepoID
			var repo *Repository
			repo = &Repository{ID: repoId}
			hasRepo, err := x.Get(repo)
			if err != nil || !hasRepo {
				fmt.Println("Repo Error!")
			}
			repos = append(repos, *repo)
		}

		/* Team: get all teams related to the user. */
		teamNum := len(teamUsers)
		teams := make([]Team, 0)
		for i := 0; i < teamNum; i++ {
			teamId := teamUsers[i].TeamID
			var team *Team
			team = &Team{ID: teamId}
			hasTeam, err := x.Get(team)
			if err != nil || !hasTeam {
				fmt.Println("Team Error!")
			}
			teams = append(teams, *team)
		}

		/* Organization: get all orgs related to the user. */
		orgNum := len(orgUsers)
		orgs := make([]User, 0)
		for i := 0; i < orgNum; i++ {
			orgId := orgUsers[i].OrgID
			var org *User
			org = &User{ID: orgId}
			hasOrg, err := x.Get(org)
			if err != nil || !hasOrg {
				fmt.Println("Organization Error!")
			}
			orgs = append(orgs, *org)
		}

		/* TeamRepo: get the relationship between repo and team. */
		/* TeamRepo: org, team and repo. */
		teamRepos := make([]TeamRepo, 0)
		for i := 0; i < repoNum; i++ {
			repoId := accesses[i].RepoID
			tempTeamRepos := make([]TeamRepo, 0)
			err = x.Where("repo_id = ?", repoId).Find(&tempTeamRepos)
			if err != nil {
				fmt.Println("TeamRepo Error!")
			}

			for j := 0; j < len(tempTeamRepos); j++ {
				teamRepos = append(teamRepos, tempTeamRepos[j])
			}
		}

		/* OrgUser: alreay in team_user. */

		/* Add the content into the file and push the file to IPFS. */
		filename := uportid + ".txt"
		file, file_err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
		if file_err != nil {
			fmt.Println("Error: create user_repo file!")
		}
		defer file.Close()

		/* Record user info into the local file. */
		user_data, _ := json.Marshal(user)
		_, file_err = file.Write(user_data)
		if file_err != nil {
			fmt.Println("Write user_data error!")
		}
		file.Write([]byte("\n\n"))

		/* Record access info into the local file. */
		lenNum := len(accesses)
		for i := 0; i < lenNum; i++ {
			access_data, _ := json.Marshal(accesses[i])
			_, file_err = file.Write(access_data)
			if file_err != nil {
				fmt.Println("Write access_data error!")
			}
			if i < lenNum-1 {
				file.Write([]byte("\n"))
			}
		}
		file.Write([]byte("\n\n"))

		/* Record teamUser info into the local file. */
		lenNum = len(teamUsers)
		for i := 0; i < lenNum; i++ {
			teamUser_data, _ := json.Marshal(teamUsers[i])
			_, file_err = file.Write(teamUser_data)
			if file_err != nil {
				fmt.Println("Write teamUser_data error!")
			}
			if i < lenNum-1 {
				file.Write([]byte("\n"))
			}
		}
		file.Write([]byte("\n\n"))

		/* Record orgUser info into the local file. */
		lenNum = len(orgUsers)
		for i := 0; i < lenNum; i++ {
			orgUser_data, _ := json.Marshal(orgUsers[i])
			_, file_err = file.Write(orgUser_data)
			if file_err != nil {
				fmt.Println("Write orgUser_data error!")
			}
			if i < lenNum-1 {
				file.Write([]byte("\n"))
			}
		}
		file.Write([]byte("\n\n"))

		/* Record collaboration info into the local file. */
		lenNum = len(collaborations)
		for i := 0; i < lenNum; i++ {
			collaborations_data, _ := json.Marshal(collaborations[i])
			_, file_err = file.Write(collaborations_data)
			if file_err != nil {
				fmt.Println("Write collaboration error!")
			}
			if i < lenNum-1 {
				file.Write([]byte("\n"))
			}
		}
		file.Write([]byte("\n\n"))

		/* Record repository info into the local file. */
		lenNum = len(repos)
		for i := 0; i < lenNum; i++ {
			repo_data, _ := json.Marshal(repos[i])
			_, file_err = file.Write(repo_data)
			if file_err != nil {
				fmt.Println("Write repo error!")
			}
			if i < lenNum-1 {
				file.Write([]byte("\n"))
			}
		}
		file.Write([]byte("\n\n"))

		/* Record team info into the local file. */
		lenNum = len(teams)
		for i := 0; i < lenNum; i++ {
			team_data, _ := json.Marshal(teams[i])
			_, file_err = file.Write(team_data)
			if file_err != nil {
				fmt.Println("Write team error!")
			}
			if i < lenNum-1 {
				file.Write([]byte("\n"))
			}
		}
		file.Write([]byte("\n\n"))

		/* Record organization info into the local file. */
		lenNum = len(orgs)
		for i := 0; i < lenNum; i++ {
			org_data, _ := json.Marshal(orgs[i])
			_, file_err = file.Write(org_data)
			if file_err != nil {
				fmt.Println("Write org error!")
			}
			if i < lenNum-1 {
				file.Write([]byte("\n"))
			}
		}
		file.Write([]byte("\n\n"))

		/* Record teamUser info into the local file. */
		lenNum = len(teamRepos)
		for i := 0; i < lenNum; i++ {
			teamRepo_data, _ := json.Marshal(teamRepos[i])
			_, file_err = file.Write(teamRepo_data)
			if file_err != nil {
				fmt.Println("Write teamRepo_data error!")
			}
			if i < lenNum-1 {
				file.Write([]byte("\n"))
			}
		}

		/* Push to local file to ipfs. */
		ipfs.Push_File_To_IPFS(filename)
	} else {
		fmt.Println("Error: no this user!")
	}
}

func GetUserRepoInfo(uportid string) {
	filename := uportid + ".txt"
	user_data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: load the uport_repo file!")
	}
	json_arrs := strings.Split(string(user_data), "\n\n")

	/* Insert into the local database. */
	/* Insert user. */
	var newUser = new(User)
	err = json.Unmarshal([]byte(json_arrs[0]), &newUser)
	if err != nil {
		fmt.Println("Error: load user info!")
	}

	/* Insert access. */
	var newAccess = new(Access)
	access_arrs := strings.Split(json_arrs[1], "\n")
	for i := 0; i < len(access_arrs); i++ {
		if access_arrs[i] != "" {
			err := json.Unmarshal([]byte(access_arrs[i]), &newAccess)
			if err != nil {
				fmt.Println("Error: load access info!")
			}
		}
	}

	/* Insert teamUser. */
	var newTeamUser = new(TeamUser)
	teamUser_arrs := strings.Split(json_arrs[2], "\n")
	for i := 0; i < len(teamUser_arrs); i++ {
		if teamUser_arrs[i] != "" {
			err := json.Unmarshal([]byte(teamUser_arrs[i]), &newTeamUser)
			if err != nil {
				fmt.Println("Error: load teamUser info!")
			}
		}
	}

	/* Insert orgUser. */
	var newOrgUser = new(OrgUser)
	orgUser_arrs := strings.Split(json_arrs[3], "\n")
	for i := 0; i < len(orgUser_arrs); i++ {
		if orgUser_arrs[i] != "" {
			err := json.Unmarshal([]byte(orgUser_arrs[i]), &newOrgUser)
			if err != nil {
				fmt.Println("Error: load orgUser info!")
			}
		}
	}

	/* Insert collaboration. */
	var newCollaboration = new(Collaboration)
	collaboration_arrs := strings.Split(json_arrs[4], "\n")
	for i := 0; i < len(collaboration_arrs); i++ {
		if collaboration_arrs[i] != "" {
			err := json.Unmarshal([]byte(collaboration_arrs[i]), &newCollaboration)
			if err != nil {
				fmt.Println("Error: load collaboration info!")
			}
		}
	}

	/* Insert repository. */
	var newRepository = new(Repository)
	repository_arrs := strings.Split(json_arrs[5], "\n")
	for i := 0; i < len(repository_arrs); i++ {
		if repository_arrs[i] != "" {
			err := json.Unmarshal([]byte(repository_arrs[i]), &newRepository)
			if err != nil {
				fmt.Println("Error: load repository info!")
			}
		}
	}

	/* Insert team. */
	var newTeam = new(Team)
	team_arrs := strings.Split(json_arrs[6], "\n")
	for i := 0; i < len(team_arrs); i++ {
		if team_arrs[i] != "" {
			err := json.Unmarshal([]byte(team_arrs[i]), &newTeam)
			if err != nil {
				fmt.Println("Error: load team info!")
			}
		}
	}

	/* Insert organization. */
	var newOraganization = new(User)
	organization_arrs := strings.Split(json_arrs[7], "\n")
	for i := 0; i < len(organization_arrs); i++ {
		if organization_arrs[i] != "" {
			err := json.Unmarshal([]byte(organization_arrs[i]), &newOraganization)
			if err != nil {
				fmt.Println("Error: load organization info!")
			}
		}
	}

	/* Insert teamRepo. */
	var newTeamRepo = new(TeamRepo)
	teamRepo_arrs := strings.Split(json_arrs[8], "\n")
	for i := 0; i < len(teamRepo_arrs); i++ {
		if teamRepo_arrs[i] != "" {
			err := json.Unmarshal([]byte(teamRepo_arrs[i]), &newTeamRepo)
			if err != nil {
				fmt.Println("Error: load teamRepo info!")
			}
		}
	}
}
