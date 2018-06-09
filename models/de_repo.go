// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"encoding/json"
	"fmt"
	"github.com/gogs/gogs/models/ipfs"
	"os/exec"
	"strings"
)

// The repo table in the IPFS
type DeRepo struct {
	ID            int64
	OwnerID       int64  `xorm:"UNIQUE(s)"`
	LowerName     string `xorm:"UNIQUE(s) INDEX NOT NULL"`
	Name          string `xorm:"INDEX NOT NULL"`
	Description   string
	Website       string
	DefaultBranch string
	Size          int64 `xorm:"NOT NULL DEFAULT 0"`

	IsPrivate bool
	IsBare    bool
	IsMirror  bool

	EnableWiki            bool `xorm:"NOT NULL DEFAULT true"`
	EnableExternalWiki    bool
	ExternalWikiURL       string
	EnableIssues          bool `xorm:"NOT NULL DEFAULT true"`
	EnableExternalTracker bool
	ExternalTrackerURL    string
	ExternalTrackerFormat string
	ExternalTrackerStyle  string
	EnablePulls           bool `xorm:"NOT NULL DEFAULT true"`
	PullsIgnoreWhitespace bool `xorm:"NOT NULL DEFAULT false"`
	PullsAllowRebase      bool `xorm:"NOT NULL DEFAULT false"`

	IsFork      bool `xorm:"NOT NULL DEFAULT false"`
	ForkID      int64
	CreatedUnix int64
	UpdatedUnix int64
}

func transferRepoToDeRepo(deRepo *DeRepo, repo *Repository) {
	deRepo.ID = repo.ID
	deRepo.OwnerID = repo.OwnerID
	deRepo.LowerName = repo.LowerName
	deRepo.Name = repo.Name
	deRepo.Description = repo.Description
	deRepo.Website = repo.Website
	deRepo.DefaultBranch = repo.DefaultBranch
	deRepo.Size = repo.Size

	deRepo.IsPrivate = repo.IsPrivate
	deRepo.IsBare = repo.IsBare
	deRepo.IsMirror = repo.IsMirror

	deRepo.EnableWiki = repo.EnableWiki
	deRepo.EnableExternalWiki = repo.EnableExternalWiki
	deRepo.ExternalWikiURL = repo.ExternalWikiURL
	deRepo.EnableIssues = repo.EnableIssues
	deRepo.EnableExternalTracker = repo.EnableExternalTracker
	deRepo.ExternalTrackerURL = repo.ExternalTrackerURL
	deRepo.ExternalTrackerFormat = repo.ExternalTrackerFormat
	deRepo.ExternalTrackerStyle = repo.ExternalTrackerStyle

	deRepo.EnablePulls = repo.EnablePulls
	deRepo.PullsAllowRebase = repo.PullsAllowRebase
	deRepo.PullsIgnoreWhitespace = repo.PullsIgnoreWhitespace

	deRepo.IsFork = repo.IsFork
	deRepo.ForkID = repo.ForkID
	deRepo.CreatedUnix = repo.CreatedUnix
	deRepo.UpdatedUnix = repo.UpdatedUnix
}

func transferDeRepoToRepo(deRepo *DeRepo, repo *Repository) error {
	repo.ID = deRepo.ID
	repo.OwnerID = deRepo.OwnerID
	repo.LowerName = deRepo.LowerName
	repo.Name = deRepo.Name
	repo.Description = deRepo.Description
	repo.Website = deRepo.Website
	repo.DefaultBranch = deRepo.DefaultBranch
	repo.Size = deRepo.Size

	repo.IsPrivate = deRepo.IsPrivate
	repo.IsBare = deRepo.IsBare
	repo.IsMirror = deRepo.IsMirror

	repo.EnableWiki = deRepo.EnableWiki
	repo.EnableExternalWiki = deRepo.EnableExternalWiki
	repo.ExternalWikiURL = deRepo.ExternalWikiURL
	repo.EnableIssues = deRepo.EnableIssues
	repo.EnableExternalTracker = deRepo.EnableExternalTracker
	repo.ExternalTrackerURL = deRepo.ExternalTrackerURL
	repo.ExternalTrackerFormat = deRepo.ExternalTrackerFormat
	repo.ExternalTrackerStyle = deRepo.ExternalTrackerStyle

	repo.EnablePulls = deRepo.EnablePulls
	repo.PullsAllowRebase = deRepo.PullsAllowRebase
	repo.PullsIgnoreWhitespace = deRepo.PullsIgnoreWhitespace

	repo.IsFork = deRepo.IsFork
	repo.ForkID = deRepo.ForkID
	repo.CreatedUnix = deRepo.CreatedUnix
	repo.UpdatedUnix = deRepo.UpdatedUnix

	// TODO
	/*type Repository struct {
		NumWatches          int
		NumStars            int
		NumForks            int
		NumIssues           int
		NumClosedIssues     int
		NumPulls            int
		NumClosedPulls      int
		NumMilestones       int `xorm:"NOT NULL DEFAULT 0"`
		NumClosedMilestones int `xorm:"NOT NULL DEFAULT 0"`

		// Advanced settings
		AllowPublicWiki   bool
		AllowPublicIssues bool
	}*/

	return nil
}

// Push the repo info to IPFS and record the new ipfsHash in the blockchain
func PushRepoInfo(modifier *User, repo *Repository) (err error) {
	if !canPushToBlockchain(modifier) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// admin or writer
	var access *Access
	access = &Access{UserID: modifier.ID, RepoID: repo.ID}
	hasAccess, err := x.Get(access)
	if err != nil {
		return fmt.Errorf("Can not get user access: %v", err)
	}

	if hasAccess {
		if access.Mode == ACCESS_MODE_OWNER {
			// Step 1: Encode repo data into JSON format
			repo_data, err := json.Marshal(repo)
			if err != nil {
				return fmt.Errorf("Can not encode repo data: %v", err)
			}

			// Step 2: Put the encoded data into IPFS
			c := fmt.Sprintf("echo '%s' | ipfs add ", repo_data)
			cmd := exec.Command("sh", "-c", c)
			out, err2 := cmd.Output()
			if err2 != nil {
				return fmt.Errorf("Push Repo to IPFS: fails: %v", err2)
			}
			ipfsHash := strings.Split(string(out), " ")[1]

			// Step3: Modify the ipfsHash in the smart contract
			// TODO: setRepoInfo(ipfsHash)
			ipfsHash = ipfsHash

		}
		if access.Mode == ACCESS_MODE_ADMIN {

		}
		if access.Mode == ACCESS_MODE_WRITE {

		}
	}

	return nil
}

// Get the new ipfsHash from the blockchain and get the repo info from IPFS
func GetRepoInfo(owner *User, repo *Repository) (err error) {
	// Step1: get the repo info hash via
	ipfsHash := "QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN"

	// Step2: get the ipfs file and get the user data
	c := fmt.Sprintf("ipfs cat ", ipfsHash)
	cmd := exec.Command("sh", "-c", c)
	repo_data, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Get Repo data from IPFS: fails: %v", err)
	}

	// Step3: unmarshall user data
	var newRepo = new(Repository)
	err = json.Unmarshal(repo_data, &newRepo)
	if err != nil {
		return fmt.Errorf("Can not decode data: %v", err)
	}

	// Step4: write into the local database
	// TODO:
	//CreateRepo(newRepo)

	return nil
}

// Push the repo content to IPFS and record the new ipfsHash in the blockchain
func PushRepoContent(modifier *User, repoPath string) (err error) {
	if !canPushToBlockchain(modifier) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step1: Push the repo to IPFS
	ipfsHash, err := ipfs.Push_Repo_To_IPFS(repoPath)
	if err != nil {
		return err
	}

	// Step2: Modify the RepoContentIpfsHash in the smart contract
	// TODO: setRepoContentIpfsHash(ipfsHash)
	ipfsHash = ipfsHash

	return nil
}

// Get the new ipfsHash from the blockchain and get the repo content from IPFS
func GetRepoContent(modifier *User) (err error) {
	// Step1: getRepoContentIpfsHash() from the smart contract
	ipfsHash := "QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN"

	// Step2: get the ipfs file and get the user data
	// TODO: ipfsHash, err := ipfs.Get_Repo_From_IPFS(ipfsHash)
	ipfsHash = ipfsHash

	return nil
}

// Push the repo collaborations to IPFS and record the new ipfsHash in the blockchain
func PushRepoCollaboration(owner *User, repo *Repository) (err error) {
	return nil
}

// Get the new ipfsHash from the blockchain and get the repo collaborations from IPFS
func GetRepoCollaboration(oowner *User, repo *Repository) (err error) {
	return nil
}

// Push the repo accesses to IPFS and record the new ipfsHash in the blockchain
func PushRepoAccess(owner *User, repo *Repository) (err error) {
	return nil
}

// Get the new ipfsHash from the blockchain and get the repo accesses from IPFS
func GetRepoAccess(owner *User, repo *Repository) (err error) {
	return nil
}
