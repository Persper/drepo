// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	//"github.com/go-xorm/xorm"
	//"github.com/gogits/gogs/models/ipfs"
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
			var teamRepo *TeamRepo
			teamRepo = &TeamRepo{RepoID: repoId}
			//find
			_, err := x.Get(teamRepo)
			if err != nil {
				fmt.Println("TeamRepo Error!")
			}
			teamRepos = append(teamRepos, *teamRepo)
		}
		for i := 0; i < len(teamRepos); i++ {
			fmt.Println(teamRepos[i].ID)
		}

		/* OrgUser: alreay in team_user. */

		/* Add the content into the file and push the file to IPFS. */
		user_data, _ := json.Marshal(user)
		filename := uportid + ".txt"
		file, file_err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
		if file_err != nil {
			fmt.Println("Error: create user_repo file!")
		}
		defer file.Close()

		_, file_err = file.Write(user_data)
		if file_err != nil {
			fmt.Println("Write user_data error!")
		}

	} else {
		fmt.Println("Error: no this user!")
	}
}

func GetUserRepoInfo(uportid string) {
	filename := uportid + ".txt"
	user_data, _ := ioutil.ReadFile(filename)
	var newUser = new(User)
	err := json.Unmarshal(user_data, &newUser)

	if err != nil {
		fmt.Println("Error: load the uport_repo file!")
	}

	//fmt.Println(newUser)
}
