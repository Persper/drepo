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

/// ***** START: DeAccess *****
type DeAccess struct {
	UserID int64 `xorm:"UNIQUE(s)"`
	Mode   AccessMode
}

func transferAcessToDeAccess(access *Access, deAccess *DeAccess) {
	deAccess.UserID = access.UserID
	deAccess.Mode = access.Mode
}

func transferDeAcessToAccess(repo *Repository, deAccess *DeAccess, access *Access) {
	// access.ID can be generated at any time
	// TODO: access.ID
	access.RepoID = repo.ID
	access.UserID = deAccess.UserID
	access.Mode = deAccess.Mode
}

/// ***** END: DeAccess *****

/// ***** START: DeCollaboration *****
type DeCollaboration struct {
	UserID int64      `xorm:"UNIQUE(s) INDEX NOT NULL"`
	Mode   AccessMode `xorm:"DEFAULT 2 NOT NULL"`
}

func transferCollaToDeColla(colla *Collaboration, deColla *DeCollaboration) {
	deColla.UserID = colla.UserID
	deColla.Mode = colla.Mode
}

func transferDeCollaToColla(repo *Repository, deColla *DeCollaboration, colla *Collaboration) {
	// colla.ID can be generated at any time
	// TODO: colla.ID
	colla.RepoID = repo.ID
	colla.UserID = deColla.UserID
	colla.Mode = deColla.Mode
}

/// ***** END: DeCollaboration *****

/// ***** START: DeRelease *****
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
	// release.ID can be generated at any time
	// TODO: release.ID
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

/// ***** END: DeRelease *****

/// ***** START: DeMilestone *****
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

/// Prerequisite: all issue exist
func transferDeMilestoneToMilestone(repo *Repository, deMilestone *DeMilestone, milestone *Milestone) error {
	milestone.ID = deMilestone.ID
	milestone.RepoID = repo.ID
	milestone.Name = deMilestone.Name
	milestone.Content = deMilestone.Content
	milestone.IsClosed = deMilestone.IsClosed
	milestone.Completeness = deMilestone.Completeness
	milestone.DeadlineUnix = deMilestone.DeadlineUnix
	milestone.ClosedDateUnix = deMilestone.ClosedDateUnix

	// ***** START: NumIssues *****
	issue := new(Issue)
	total, err := x.Where("milestone_id = ?", milestone.ID).Count(issue)
	if err != nil {
		return fmt.Errorf("Can not get repo issues: %v\n", err)
	}
	milestone.NumIssues = int(total)
	// ***** END: NumIssues *****

	// ***** START: NumClosedIssues *****
	closedIssue := new(Issue)
	total, err = x.Where("milestone_id = ? and is_closed = ?", milestone.ID, true).Count(closedIssue)
	if err != nil {
		return fmt.Errorf("Can not get repo closedIssues: %v\n", err)
	}
	milestone.NumClosedIssues = int(total)
	// ***** END: NumClosedIssues *****

	return nil
}

/// ***** END: DeMilestone *****

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
	AllowPublicWiki       bool
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
	deRepo.AllowPublicWiki = repo.AllowPublicWiki
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
		return fmt.Errorf("Can not get accesses of the user: %v\n", err)
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
		return fmt.Errorf("Can not get collaborations of the user: %v\n", err)
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
		return fmt.Errorf("Can not get releases of the user: %v\n", err)
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
		return fmt.Errorf("Can not get milestones of the user: %v\n", err)
	}
	for i := range milestones {
		deMilestone := new(DeMilestone)
		transferMilestoneToDeMilestone(&milestones[i], deMilestone)
		deRepo.Milestones = append(deRepo.Milestones, *deMilestone)
	}
	// ***** END: Milestone[] *****

	return nil
}

/// Prerequisite: issue / watch / star / milistone/ pulls /
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
	repo.AllowPublicWiki = deRepo.AllowPublicWiki
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

	// ***** START: Access[] *****
	for i := range deRepo.Accesses {
		access := new(Access)
		transferDeAcessToAccess(repo, &deRepo.Accesses[i], access)
		has, err := x.Get(access)
		if err != nil {
			return fmt.Errorf("Can not search the access: %v\n", err)
		}
		if !has {
			_, err = x.Insert(access)
			if err != nil {
				return fmt.Errorf("Can not add the access to the server: %v\n", err)
			}
		}
	}
	// ***** END: Access[] *****

	// ***** START: Collaboration[] *****
	for i := range deRepo.Collaborations {
		collaboration := new(Collaboration)
		transferDeCollaToColla(repo, &deRepo.Collaborations[i], collaboration)
		has, err := x.Get(collaboration)
		if err != nil {
			return fmt.Errorf("Can not search the collaboration: %v\n", err)
		}
		if !has {
			_, err = x.Insert(collaboration)
			if err != nil {
				return fmt.Errorf("Can not add the collaboration to the server: %v\n", err)
			}
		}
	}
	// ***** END: Collaboration[] *****

	// ***** START: Release[] *****
	for i := range deRepo.Releases {
		release := new(Release)
		transferDeReleaseToRelease(repo, &deRepo.Releases[i], release)
		has, err := x.Get(release)
		if err != nil {
			return fmt.Errorf("Can not search the release: %v\n", err)
		}
		if !has {
			_, err = x.Insert(release)
			if err != nil {
				return fmt.Errorf("Can not add the release to the server: %v\n", err)
			}
		}
	}
	// ***** END: Release[] *****

	// ***** START: milestone[] *****
	for i := range deRepo.Milestones {
		milestone := new(Milestone)
		transferDeMilestoneToMilestone(repo, &deRepo.Milestones[i], milestone)
		has, err := x.Get(milestone)
		if err != nil {
			return fmt.Errorf("Can not search the milestone: %v\n", err)
		}
		if !has {
			_, err = x.Insert(milestone)
			if err != nil {
				return fmt.Errorf("Can not add the milestone to the server: %v\n", err)
			}
		}

		// Update: milestoneID -> issue - > issueUser
		issues := make([]Issue, 0)
		if err := x.Find(&issues, &Issue{MilestoneID: deRepo.Milestones[i].ID}); err != nil {
			return fmt.Errorf("Can not get issues of the user: %v\n", err)
		}
		for j := range issues {
			issueUsers := make([]IssueUser, 0)
			if err := x.Find(&issueUsers, &IssueUser{IssueID: issues[j].ID}); err != nil {
				return fmt.Errorf("Can not get IssueUsers of the issue: %v\n", err)
			}
			for k := range issueUsers {
				_, err := x.Exec("UPDATE `issue_user` SET milestone_id=? WHERE id=?",
					deRepo.Milestones[i].ID, issueUsers[k].ID)
				if err != nil {
					return fmt.Errorf("Can not update milestone_id of the issueUser: %v\n", err)
				}
			}
		}
	}
	// ***** END: milestone[] *****

	// ***** START: NumIssues *****
	issue := new(Issue)
	total, err := x.Where("repo_id = ?", repo.ID).Count(issue)
	if err != nil {
		return fmt.Errorf("Can not get repo issues: %v\n", err)
	}
	repo.NumIssues = int(total)
	// ***** END: NumIssues *****

	// ***** START: NumClosedIssues *****
	closedIssue := new(Issue)
	total, err = x.Where("repo_id = ? and is_closed = ?", repo.ID, true).Count(closedIssue)
	if err != nil {
		return fmt.Errorf("Can not get repo closedIssues: %v\n", err)
	}
	repo.NumClosedIssues = int(total)
	// ***** END: NumClosedIssues *****

	// ***** START: NumWatches *****
	watch := new(Watch)
	total, err = x.Where("repo_id = ?", repo.ID).Count(watch)
	if err != nil {
		return fmt.Errorf("Can not get repo watches: %v\n", err)
	}
	repo.NumWatches = int(total)
	// ***** END: NumWatches *****

	// ***** START: NumStars *****
	star := new(Star)
	total, err = x.Where("repo_id = ?", repo.ID).Count(star)
	if err != nil {
		return fmt.Errorf("Can not get repo stars: %v\n", err)
	}
	repo.NumStars = int(total)
	// ***** END: NumStars *****

	// ***** START: NumMilestones *****
	milestone := new(Milestone)
	total, err = x.Where("repo_id = ?", repo.ID).Count(milestone)
	if err != nil {
		return fmt.Errorf("Can not get repo milestones: %v\n", err)
	}
	repo.NumMilestones = int(total)
	// ***** END: NumMilestones *****

	// ***** START: NumClosedMilestones *****
	closedMilestone := new(Milestone)
	total, err = x.Where("repo_id = ? and is_closed = ?", repo.ID, true).Count(closedMilestone)
	if err != nil {
		return fmt.Errorf("Can not get repo closedMilestones: %v\n", err)
	}
	repo.NumClosedMilestones = int(total)
	// ***** END: NumClosedMilestones *****

	// ***** START: NumPulls *****
	pullRequest := new(PullRequest)
	total, err = x.Where("repo_id = ?", repo.ID).Count(pullRequest)
	if err != nil {
		return fmt.Errorf("Can not get repo pullRequests: %v\n", err)
	}
	repo.NumPulls = int(total)
	// ***** END: NumPulls *****

	// ***** START: NumClosedPulls *****
	// Todo: not sure
	pullRequest = new(PullRequest)
	total, err = x.Where("repo_id = ? and has_merged", repo.ID, true).Count(pullRequest)
	if err != nil {
		return fmt.Errorf("Can not get repo closedPullRequests: %v\n", err)
	}
	repo.NumPulls = int(total)
	// ***** END: NumClosedPulls *****

	// ***** START: NumForks *****
	forkRepo := new(Repository)
	total, err = x.Where("fork_id = ?", repo.ID).Count(forkRepo)
	if err != nil {
		return fmt.Errorf("Can not get fork repos: %v\n", err)
	}
	repo.NumForks = int(total)

	if repo.IsFork {
		forkingRepo := &Repository{ID: repo.ForkID}
		has, err := x.Get(forkingRepo)
		if err != nil {
			return fmt.Errorf("Can not search the forkingRepo: %v\n", err)
		}
		if has {
			if _, err = x.Exec("UPDATE `repository` SET num_forks=num_forks+1 WHERE id=?", forkingRepo.ID); err != nil {
				return fmt.Errorf("update forkingRepo: %v\n", err)
			}
		}
	}
	// ***** END: NumForks *****

	return nil
}

// Push the repo info to IPFS and record the new ipfsHash in the blockchain
func PushRepoInfo(user *User, repo *Repository) (err error) {
	if !canPushToBlockchain(user) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step0: push the repo content to the IPFS
	err = PushRepoContent(user, repo.RepoPath())
	if err != nil {
		return err
	}

	// step1: push the repo table to the IPFS
	var access *Access
	access = &Access{UserID: user.ID, RepoID: repo.ID}
	_, err = x.Get(access)
	if err != nil {
		return fmt.Errorf("Can not get user access: %v\n", err)
	}

	if repo.OwnerID == user.ID || access.Mode == ACCESS_MODE_OWNER {
		// Step 1: Encode repo data into JSON format
		deRepo := new(DeRepo)
		transferRepoToDeRepo(repo, deRepo)
		repo_data, err := json.Marshal(deRepo)
		if err != nil {
			return fmt.Errorf("Can not encode repo data: %v\n", err)
		}

		// Step 2: Put the encoded data into IPFS
		c := fmt.Sprintf("echo '%s' | ipfs add ", repo_data)
		cmd := exec.Command("sh", "-c", c)
		out, err2 := cmd.Output()
		if err2 != nil {
			return fmt.Errorf("Push Repo to IPFS: fails: %v\n", err2)
		}
		ipfsHash := strings.Split(string(out), " ")[1]

		// Step3: Modify the ipfsHash in the smart contract
		// TODO: setRepoInfo(ipfsHash)
		ipfsHash = ipfsHash
		fmt.Println("Push the repo file to the IPFS: " + ipfsHash)

	}
	if access.Mode == ACCESS_MODE_ADMIN {

	}
	if access.Mode == ACCESS_MODE_WRITE {

	}

	return nil
}

// Get the new ipfsHash from the blockchain and get the repo info from IPFS
func GetRepoInfo(user *User, ipfsHash string) (*Repository, error) {
	// Step1: get the repo info hash

	// Step2: get the ipfs file and get the user data
	c := "ipfs cat " + ipfsHash
	cmd := exec.Command("sh", "-c", c)
	repo_data, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Get Repo data from IPFS: fails: %v\n", err)
	}

	// Step3: unmarshall user data
	newDeRepo := new(DeRepo)
	err = json.Unmarshal(repo_data, &newDeRepo)
	if err != nil {
		return nil, fmt.Errorf("Can not decode data: %v\n", err)
	}

	// Step4: write into the local database and mkdir the local path
	newRepo := new(Repository)
	transferDeRepoToRepo(newDeRepo, newRepo)
	has, err2 := x.Get(newRepo)
	if err2 != nil {
		return nil, fmt.Errorf("Can not search the repo: %v\n", err2)
	}

	if !has {
		sess := x.NewSession()
		defer sess.Close()
		if err = sess.Begin(); err != nil {
			return nil, err
		}

		has, err = x.Get(newRepo)
		if err != nil {
			return nil, fmt.Errorf("Can not get the repo: %v\n", err2)
		}
		if !has {
			if _, err = sess.Insert(newRepo); err != nil {
				return nil, fmt.Errorf("Can not insert the repo: %v\n", err)
			}
			user.NumRepos++
			if _, err = sess.Update(user); err != nil {
				return nil, fmt.Errorf("Can not update the user: %v\n", err)
			}

			repoPath := RepoPath(user.Name, newRepo.Name)
			fmt.Println("repopath:" + repoPath)
			if err := GetRepoContent(user, repoPath); err != nil {
				return nil, err
			}
		}

		return newRepo, sess.Commit()
	}

	// TODO: watch

	return nil, nil
}

// Push the repo content to IPFS and record the new ipfsHash in the blockchain
func PushRepoContent(user *User, repoPath string) (err error) {
	if !canPushToBlockchain(user) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step1: Push the repo to IPFS
	c := "ipfs add -r  " + repoPath
	cmd := exec.Command("sh", "-c", c)
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Push repoContent to IPFS: fails: %v\n", err)
	}
	ipfsHashs := strings.Split(string(out), " ")
	ipfsHash := ipfsHashs[len(ipfsHashs)-2]

	// Step2: Modify the RepoContentIpfsHash in the smart contract
	// TODO: setRepoContentIpfsHash(ipfsHash)
	ipfsHash = ipfsHash
	fmt.Println("Push the repo content to the IPFS: " + ipfsHash)

	return nil
}

// Get the new ipfsHash from the blockchain and get the repo content from IPFS
func GetRepoContent(modifier *User, repoPath string) (err error) {
	// Step1: get the repo content hash
	ipfsHash := "QmYkMofbGtqBozUrG5LjFpMpg8Fhxw7ffJa8WwxtAvooRe"
	// QmS63hLK2uridjdJyKPNyk8enAqNH74YJDX1t6H4rjdY31

	// Step2: get the ipfs file and get the user data
	c := "ipfs get " + ipfsHash
	cmd := exec.Command("sh", "-c", c)
	_, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("Get repo content from IPFS: fails: %v\n", err)
	}

	// Step3: move the .git dir to the repoPath
	c = "mv " + ipfsHash + " " + repoPath
	cmd = exec.Command("sh", "-c", c)
	_, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("Move repo content to the targeted dir: fails: %v\n", err)
	}

	return nil
}

/// TODO: Push the whole .git dir rather than only a branch
/// Push the repo content to IPFS and record the new ipfsHash in the blockchain
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

/// Get the new ipfsHash from the blockchain and get the repo content from IPFS
func GetRepoContentByDegit(modifier *User) (err error) {
	// Step1: getRepoContentIpfsHash() from the smart contract
	ipfsHash := "QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN"

	// Step2: get the ipfs file and get the user data
	// TODO: ipfsHash, err := ipfs.Get_Repo_From_IPFS(ipfsHash)
	ipfsHash = ipfsHash

	return nil
}

/// The repo button: push the repo info and all related tables to IPFS
func PushRepoAndRelatedTables(contextUser *User, repo *Repository) (err error) {
	if !canPushToBlockchain(contextUser) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step1: get the corresponding user.
	var user *User
	user = &User{ID: contextUser.ID}
	_, err = x.Get(user)
	if err != nil {
		return fmt.Errorf("Can not get user data: %v", err)
	}

	// Step2: push the repo
	if err = PushRepoInfo(user, repo); err != nil {
		return fmt.Errorf("Can not push repo data: %v", err)
	}

	issues := make([]Issue, 0)
	if err = x.Find(&issues, &Issue{RepoID: repo.ID}); err != nil {
		return fmt.Errorf("Can not get issues of the repo: %v", err)
	}
	for j := range issues {
		if err = PushIssueInfo(user, &issues[j]); err != nil {
			return fmt.Errorf("Can not push issue data: %v", err)
		}

		prs := make([]PullRequest, 0)
		if err = x.Find(&prs, &PullRequest{IssueID: issues[j].ID}); err != nil {
			return fmt.Errorf("Can not get pull_request of the repo: %v", err)
		}

		for k := range prs {
			if err = PushPullInfo(user, &prs[k]); err != nil {
				return fmt.Errorf("Can not push pull_request data: %v", err)
			}
		}
	}

	branches := make([]ProtectBranch, 0)
	if err = x.Find(&branches, &ProtectBranch{RepoID: repo.ID}); err != nil {
		return fmt.Errorf("Can not get branches of the repo: %v", err)
	}
	for j := range branches {
		if err = PushBranchInfo(user, &branches[j]); err != nil {
			return fmt.Errorf("Can not push branch data: %v", err)
		}
	}

	return nil
}

/// The repo button: get the repo info and all related tables to IPFS
func GetRepoAndRelatedTables(user *User, ipfsHash string) (err error) {
	// Just for test
	var repo *Repository
	if repo, err = GetRepoInfo(user, ipfsHash); err != nil {
		return err
	}
	return nil

	// TODO: from the blockchain
	branchHashes := make([]string, 0)
	for i := range branchHashes {
		if err := GetBranchInfo(user, repo, branchHashes[i]); err != nil {
			return err
		}
	}
	// TODO: from the blockchain
	prHashes := make([]string, 0)
	for i := range prHashes {
		if err := GetPullInfo(user, repo, prHashes[i]); err != nil {
			return err
		}
	}
	// TODO: from the blockchain
	issues := make([]Issue, 0)
	for i := range issues {
		if err := GetIssueInfo(user, repo, &issues[i]); err != nil {
			return err
		}
	}
	if _, err := GetRepoInfo(user, ipfsHash); err != nil {
		return err
	}

	return nil
}
