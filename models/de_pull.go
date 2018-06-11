// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import ()

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
