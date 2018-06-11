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

type DeAccess struct {
	UserID int64 `xorm:"UNIQUE(s)"`
	Mode   AccessMode
}

func transferAcessToDeAccess(access *Access, deAccess *DeAccess) {
	deAccess.UserID = access.UserID
	deAccess.Mode = access.Mode
}

func transferDeAcessToAccess(repo *Repository, deAccess *DeAccess, access Access) {
	// access.ID
	access.UserID = deAccess.UserID
	access.Mode = deAccess.Mode
	access.RepoID = repo.ID
}

type DeCollaboration struct {
	UserID int64      `xorm:"UNIQUE(s) INDEX NOT NULL"`
	Mode   AccessMode `xorm:"DEFAULT 2 NOT NULL"`
}

func transferCollaToDeColla(colla *Collaboration, deColla *DeCollaboration) {
	deColla.UserID = colla.UserID
	deColla.Mode = colla.Mode
}

func transferDeCollaToColla(repo *Repository, deColla *DeCollaboration, colla Collaboration) {
	//colla.ID
	colla.RepoID = repo.ID
	colla.UserID = deColla.UserID
	colla.Mode = deColla.Mode
}

type DeRelease struct {
	PublisherID  int64
	TagName      string
	LowerTagName string
	Target       string
	Title        string
	Sha1         string `xorm:"VARCHAR(40)"`
	NumCommits   int64
	Note         string `xorm:"TEXT"`
	IsDraft      bool   `xorm:"NOT NULL DEFAULT false"`
	IsPrerelease bool
	CreatedUnix  int64
}

func transferReleaseToDeRelease(release *Release, deRelease *DeRelease) {
	deRelease.PublisherID = release.PublisherID
	deRelease.TagName = release.TagName
	deRelease.LowerTagName = release.LowerTagName
	deRelease.Target = release.Target
	deRelease.Title = release.Title
	deRelease.Sha1 = release.Sha1
	deRelease.NumCommits = release.NumCommits
	deRelease.Note = release.Note
	deRelease.IsDraft = release.IsDraft
	deRelease.IsPrerelease = release.IsPrerelease
	deRelease.CreatedUnix = release.CreatedUnix
}

func transferDeReleaseToRelease(repo *Repository, deRelease *DeRelease, release *Release) {
	//release.ID =
	release.RepoID = repo.ID
	release.PublisherID = deRelease.PublisherID
	release.TagName = deRelease.TagName
	release.LowerTagName = deRelease.LowerTagName
	release.Target = deRelease.Target
	release.Title = deRelease.Title
	release.Sha1 = deRelease.Sha1
	release.NumCommits = deRelease.NumCommits
	release.Note = deRelease.Note
	release.IsDraft = deRelease.IsDraft
	release.IsPrerelease = deRelease.IsPrerelease
	release.CreatedUnix = deRelease.CreatedUnix
}

type DeMilestone struct {
	ID             int64
	Name           string
	Content        string `xorm:"TEXT"`
	IsClosed       bool
	Completeness   int // Percentage(1-100).
	DeadlineUnix   int64
	ClosedDateUnix int64
}

func transferMilestoneToDeMilestone(milestone *Milestone, deMilestone *DeMilestone) {
	deMilestone.ID = milestone.ID
	deMilestone.Name = milestone.Name
	deMilestone.Content = milestone.Content
	deMilestone.IsClosed = milestone.IsClosed
	deMilestone.Completeness = milestone.Completeness
	deMilestone.DeadlineUnix = milestone.DeadlineUnix
	deMilestone.ClosedDateUnix = milestone.ClosedDateUnix
}

func transferDeMilestoneToMilestone(repo *Repository, deMilestone *DeMilestone, milestone *Milestone) {
	milestone.ID = deMilestone.ID
	milestone.RepoID = repo.ID
	milestone.Name = deMilestone.Name
	milestone.Content = deMilestone.Content
	milestone.IsClosed = deMilestone.IsClosed
	milestone.Completeness = deMilestone.Completeness
	milestone.DeadlineUnix = deMilestone.DeadlineUnix
	milestone.ClosedDateUnix = deMilestone.ClosedDateUnix

	// TODO:
	// NumIssues
	// NumClosedIssues
}

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
	AllowPublicIssues     bool
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

	Accesses       []DeAccess        `xorm:"-"`
	Collaborations []DeCollaboration `xorm:"-"`
	Releases       []DeRelease       `xorm:"-"`
	Milestones     []DeMilestone     `xorm:"-"`

	/*
		release.{publisher_id, tag_name, lower_tag_name, target, title, sha1, num_commits, note, is_draft, is_prerelease, created_unix}[]
		milestone.{id, name, content, is_closed, completeness, deadline_unix, closed_date_unix}[]: The milestone ID begins with the address of the owner of this repo.
	*/
}

func transferRepoToDeRepo(repo *Repository, deRepo *DeRepo) error {
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
	deRepo.AllowPublicIssues = repo.AllowPublicIssues
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

	// ***** START: Access[] *****
	accesses := make([]Access, 0)
	if err := x.Find(&accesses, &Access{RepoID: repo.ID}); err != nil {
		return fmt.Errorf("Can not get accesses of the user: %v", err)
	}
	for i := range accesses {
		deAccess := new(DeAccess)
		transferAcessToDeAccess(&accesses[i], deAccess)
		deRepo.Accesses = append(deRepo.Accesses, *deAccess)
	}
	// ***** END: Access[] *****

	// ***** START: Collaboration[] *****
	collaborations := make([]Collaboration, 0)
	if err := x.Find(&collaborations, &Collaboration{RepoID: repo.ID}); err != nil {
		return fmt.Errorf("Can not get collaborations of the user: %v", err)
	}
	for i := range collaborations {
		deCollaboration := new(DeCollaboration)
		transferCollaToDeColla(&collaborations[i], deCollaboration)
		deRepo.Collaborations = append(deRepo.Collaborations, *deCollaboration)
	}
	// ***** END: Collaboration[] *****

	// ***** START: Release[] *****
	releases := make([]Release, 0)
	if err := x.Find(&releases, &Release{RepoID: repo.ID}); err != nil {
		return fmt.Errorf("Can not get releases of the user: %v", err)
	}
	for i := range releases {
		deRelease := new(DeRelease)
		transferReleaseToDeRelease(&releases[i], deRelease)
		deRepo.Releases = append(deRepo.Releases, *deRelease)
	}
	// ***** END: Release[] *****

	// ***** START: Milestone[] *****
	milestones := make([]Milestone, 0)
	if err := x.Find(&milestones, &Milestone{RepoID: repo.ID}); err != nil {
		return fmt.Errorf("Can not get milestones of the user: %v", err)
	}
	for i := range milestones {
		deMilestone := new(DeMilestone)
		transferMilestoneToDeMilestone(&milestones[i], deMilestone)
		deRepo.Milestones = append(deRepo.Milestones, *deMilestone)
	}
	// ***** END: Milestone[] *****

	return nil
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
	repo.AllowPublicIssues = deRepo.AllowPublicIssues
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

	// TODO: Access
	// TODO: Collaboration
	// TODO: Release

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

	if hasAccess || repo.OwnerID == modifier.ID {
		if repo.OwnerID == modifier.ID || access.Mode == ACCESS_MODE_OWNER {
			// Step 0: Push the repo content to the IPFS
			err0 := PushRepoContent(modifier, repo.RepoPath())
			if err0 != nil {
				return fmt.Errorf("Can not push repo content: %v", err0)
			}

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
	c := fmt.Sprintf("echo '%s' | ipfs add -r ", repoPath)
	cmd := exec.Command("sh", "-c", c)
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Push RepoContent to IPFS: fails: %v", err)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step2: Modify the RepoContentIpfsHash in the smart contract
	// TODO: setRepoContentIpfsHash(ipfsHash)
	ipfsHash = ipfsHash

	return nil
}

// Get the new ipfsHash from the blockchain and get the repo content from IPFS
func GetRepoContent(modifier *User) (err error) {
	// TODO
	return nil
}

// Push the repo content to IPFS and record the new ipfsHash in the blockchain
// Push only a branch
func PushRepoContentByDegit(modifier *User, repoPath string) (err error) {
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
// Get only a branch
func GetRepoContentByDegit(modifier *User) (err error) {
	// Step1: getRepoContentIpfsHash() from the smart contract
	ipfsHash := "QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN"

	// Step2: get the ipfs file and get the user data
	// TODO: ipfsHash, err := ipfs.Get_Repo_From_IPFS(ipfsHash)
	ipfsHash = ipfsHash

	return nil
}
