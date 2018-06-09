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
	deIssue.MilestoneID = issue.Milestone
	deIssue.Priority = issue.PosterID
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
	issue.MilestoneID = deIssue.Milestone
	issue.Priority = deIssue.PosterID
	issue.AssigneeID = deIssue.AssigneeID
	issue.IsClosed = deIssue.IsClosed
	issue.IsPull = deIssue.IsPull
	issue.DeadlineUnix = deIssue.DeadlineUnix
	issue.CreatedUnix = deIssue.CreatedUnix
	issue.UpdatedUnix = deIssue.UpdatedUnix
}
