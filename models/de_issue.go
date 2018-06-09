// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import ()

// The issue table in the IPFS
type DeIssue struct {
	Index        int64 `xorm:"UNIQUE(repo_index)"` // Index in one repository.
	PosterID     int64
	Title        string `xorm:"name"`
	Content      string `xorm:"TEXT"`
	MilestoneID  int64
	Priority     int
	AssigneeID   int64
	IsClosed     bool
	IsPull       bool // Indicates whether is a pull request or not.
	DeadlineUnix int64
	CreatedUnix  int64
	UpdatedUnix  int64
}

func transferIssueToDeIssue(deIssue *DeIssue, issue *Issue) {
	deIssue.Index = issue.Index
	deIssue.PosterID = issue.PosterID
	deIssue.Title = issue.Title
	deIssue.Content = issue.Content
	deIssue.MilestoneID = issue.MilestoneID
	deIssue.Priority = issue.Priority
	deIssue.AssigneeID = issue.AssigneeID
	deIssue.IsClosed = issue.IsClosed
	deIssue.IsPull = issue.IsPull
	deIssue.DeadlineUnix = issue.DeadlineUnix
	deIssue.CreatedUnix = issue.CreatedUnix
	deIssue.UpdatedUnix = issue.UpdatedUnix
}

func transferDeIssueToIssue(deIssue *DeIssue, issue *Issue) error {
	issue.Index = deIssue.Index
	issue.PosterID = deIssue.PosterID
	issue.Title = deIssue.Title
	issue.Content = deIssue.Content
	issue.MilestoneID = deIssue.MilestoneID
	issue.Priority = deIssue.Priority
	issue.AssigneeID = deIssue.AssigneeID
	issue.IsClosed = deIssue.IsClosed
	issue.IsPull = deIssue.IsPull
	issue.DeadlineUnix = deIssue.DeadlineUnix
	issue.CreatedUnix = deIssue.CreatedUnix
	issue.UpdatedUnix = deIssue.UpdatedUnix

	return nil
}

// The pull table in the IPFS
type DePull struct {
	Type   PullRequestType
	Status PullRequestStatus

	IssueID int64 `xorm:"INDEX"`
	Index   int64

	HeadRepoID   int64
	HeadUserName string
	HeadBranch   string
	BaseBranch   string
	MergeBase    string `xorm:"VARCHAR(40)"`

	HasMerged      bool
	MergedCommitID string `xorm:"VARCHAR(40)"`
	MergerID       int64
	MergedUnix     int64
}

func transferPullToDePull(dePull *DePull, pull *PullRequest) {
	dePull.Type = pull.Type
	dePull.Status = pull.Status
	dePull.IssueID = pull.IssueID
	dePull.Index = pull.Index
	dePull.HeadRepoID = pull.HeadRepoID
	dePull.HeadUserName = pull.HeadUserName
	dePull.HeadBranch = pull.HeadBranch
	dePull.BaseBranch = pull.BaseBranch
	dePull.MergeBase = pull.MergeBase
	dePull.HasMerged = pull.HasMerged
	dePull.MergedCommitID = pull.MergedCommitID
	dePull.MergerID = pull.MergerID
	dePull.MergedUnix = pull.MergedUnix
}

func transferDePullToPull(dePull *DePull, pull *PullRequest) error {
	pull.Type = dePull.Type
	pull.Status = dePull.Status
	pull.IssueID = dePull.IssueID
	pull.Index = dePull.Index
	pull.HeadRepoID = dePull.HeadRepoID
	pull.HeadUserName = dePull.HeadUserName
	pull.HeadBranch = dePull.HeadBranch
	pull.BaseBranch = dePull.BaseBranch
	pull.MergeBase = dePull.MergeBase
	pull.HasMerged = dePull.HasMerged
	pull.MergedCommitID = dePull.MergedCommitID
	pull.MergerID = dePull.MergerID
	pull.MergedUnix = dePull.MergedUnix

	return nil
}

// The pull table in the IPFS
type DeBranch struct {
	Name               string `xorm:"UNIQUE(protect_branch)"`
	Protected          bool
	RequirePullRequest bool
	EnableWhitelist    bool
	WhitelistUserIDs   string `xorm:"TEXT"`
	WhitelistTeamIDs   string `xorm:"TEXT"`
}

func transferBranchToDeBranch(deBranch *DeBranch, branch *ProtectBranch) {
	deBranch.Name = branch.Name
	deBranch.Protected = branch.Protected
	deBranch.RequirePullRequest = branch.RequirePullRequest
	deBranch.EnableWhitelist = branch.EnableWhitelist
	deBranch.WhitelistUserIDs = branch.WhitelistUserIDs
	deBranch.WhitelistTeamIDs = branch.WhitelistTeamIDs
}

func transferDeBranchToBranch(deBranch *DeBranch, branch *ProtectBranch) {
	branch.Name = deBranch.Name
	branch.Protected = deBranch.Protected
	branch.RequirePullRequest = deBranch.RequirePullRequest
	branch.EnableWhitelist = deBranch.EnableWhitelist
	branch.WhitelistUserIDs = deBranch.WhitelistUserIDs
	branch.WhitelistTeamIDs = deBranch.WhitelistTeamIDs
}
