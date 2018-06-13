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

/// ***** START: DeIssueUser *****
type DeIssueUser struct {
	UID         int64 `xorm:"INDEX"` // User ID.
	IsRead      bool
	IsAssigned  bool
	IsMentioned bool
	IsPoster    bool
	IsClosed    bool
}

func transferIssueUserToDeIssueUser(issueUser *IssueUser, deIssueUser *DeIssueUser) {
	deIssueUser.UID = issueUser.UID
	deIssueUser.IsRead = issueUser.IsRead
	deIssueUser.IsAssigned = issueUser.IsAssigned
	deIssueUser.IsMentioned = issueUser.IsMentioned
	deIssueUser.IsPoster = issueUser.IsPoster
	deIssueUser.IsClosed = issueUser.IsClosed
}

func transferDeIssueUserToIssueUser(issue *Issue, repo *Repository, milestone *Milestone,
	issueUser *IssueUser, deIssueUser *DeIssueUser) {
	// issueUser.ID can be generated at any time
	// TODO: issueUser.ID
	issueUser.UID = deIssueUser.UID
	issueUser.IssueID = issue.ID
	issueUser.RepoID = repo.ID
	issueUser.MilestoneID = milestone.ID
	issueUser.IsRead = deIssueUser.IsRead
	issueUser.IsAssigned = deIssueUser.IsAssigned
	issueUser.IsMentioned = deIssueUser.IsMentioned
	issueUser.IsPoster = deIssueUser.IsPoster
	issueUser.IsClosed = deIssueUser.IsClosed
}

/// ***** END: DeIssueUser *****

/// ***** START: DeComment *****
type DeComment struct {
	Type        CommentType
	PosterID    int64
	CommitID    int64
	Line        int64
	Content     string `xorm:"TEXT"`
	CreatedUnix int64
	UpdatedUnix int64
	CommitSHA   string `xorm:"VARCHAR(40)"`
}

func transferCommentToDeComment(comment *Comment, deComment *DeComment) {
	deComment.Type = comment.Type
	deComment.PosterID = comment.PosterID
	deComment.CommitID = comment.CommitID
	deComment.Line = comment.Line
	deComment.Content = comment.Content
	deComment.CreatedUnix = comment.CreatedUnix
	deComment.UpdatedUnix = comment.UpdatedUnix
	deComment.CommitSHA = comment.CommitSHA
}

func transferDeCommentToComment(issue *Issue, deComment *DeComment, comment *Comment) {
	// comment.ID can be generated at any time
	// TODO: comment.ID
	comment.Type = deComment.Type
	comment.PosterID = deComment.PosterID
	comment.IssueID = issue.ID
	comment.CommitID = deComment.CommitID
	comment.Line = deComment.Line
	comment.Content = deComment.Content
	comment.CreatedUnix = deComment.CreatedUnix
	comment.UpdatedUnix = deComment.UpdatedUnix
	comment.CommitSHA = deComment.CommitSHA
}

/// ***** END: DeComment *****

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
	Comments     []DeComment   `xorm:"-"`
	IssueUsers   []DeIssueUser `xorm:"-"`
}

func transferIssueToDeIssue(issue *Issue, deIssue *DeIssue) error {
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

	// ***** START: Comment[] *****
	comments := make([]Comment, 0)
	if err := x.Find(&comments, &Comment{IssueID: issue.ID}); err != nil {
		return fmt.Errorf("Can not get comments of the user: %v\n", err)
	}
	for i := range comments {
		deComment := new(DeComment)
		transferCommentToDeComment(&comments[i], deComment)
		deIssue.Comments = append(deIssue.Comments, *deComment)
	}
	// ***** END: Comment[] *****

	// ***** START: IssueUser[] *****
	issueUsers := make([]IssueUser, 0)
	if err := x.Find(&issueUsers, &IssueUser{IssueID: issue.ID}); err != nil {
		return fmt.Errorf("Can not get issueUsers of the user: %v\n", err)
	}
	for i := range issueUsers {
		deIssueUser := new(DeIssueUser)
		transferIssueUserToDeIssueUser(&issueUsers[i], deIssueUser)
		deIssue.IssueUsers = append(deIssue.IssueUsers, *deIssueUser)
	}
	// ***** END: IssueUser[] *****

	return nil
}

/// Prerequisite: all milestone exist
func transferDeIssueToIssue(repo *Repository, deIssue *DeIssue, issue *Issue) error {
	// issue.ID can be generated at any time
	// TODO: issue.ID
	issue.RepoID = repo.ID
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

	// ***** START: Comment[] *****
	for i := range deIssue.Comments {
		comment := new(Comment)
		transferDeCommentToComment(issue, &deIssue.Comments[i], comment)
		has, err := x.Get(comment)
		if err != nil {
			return fmt.Errorf("Can not search the comment: %v\n", err)
		}
		if !has {
			_, err = x.Insert(comment)
			if err != nil {
				return fmt.Errorf("Can not add the comment to the server: %v\n", err)
			}
		}
	}
	issue.NumComments = len(deIssue.Comments)
	// ***** END: Comment[] *****

	// ***** START: IssueUser[] *****
	// TODO: need milestone
	/*for i := range deIssue.IssueUsers {
		issueUser := new(IssueUser)
		transferDeIssueUserToIssueUser(issue, repo, milestone, &deIssue.IssueUsers[i], issueUser)
		has, err := x.Get(issueUser)
		if err != nil {
			return fmt.Errorf("Can not search the issueUser: %v\n", err)
		}
		if !has {
			_, err = x.Insert(issueUser)
			if err != nil {
				return fmt.Errorf("Can not add the issueUser to the server: %v\n", err)
			}
		}
	}*/
	// ***** END: IssueUser[] *****
	return nil
}

func PushIssueInfo(user *User, issue *Issue) error {
	if !canPushToBlockchain(user) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step 1: Encode org data into JSON format
	deIssue := new(DeIssue)
	transferIssueToDeIssue(issue, deIssue)
	issue_data, err := json.Marshal(deIssue)
	if err != nil {
		return fmt.Errorf("Can not encode issue data: %v", err)
	}

	// Step 2: Put the encoded data into IPFS
	c := fmt.Sprintf("echo '%s' | ipfs add ", issue_data)
	cmd := exec.Command("sh", "-c", c)
	out, err2 := cmd.Output()
	if err2 != nil {
		return fmt.Errorf("Push issue to IPFS: fails: %v", err2)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step4: Modify the ipfsHash in the smart contract
	// TODO: setIssueInfo(ipfsHash)
	ipfsHash = ipfsHash
	fmt.Println("Push the issue file to the IPFS: " + ipfsHash)

	return nil
}

func GetIssueInfo(user *User, repo *Repository, issue *Issue) error {
	// Step1: get the issue info hash
	ipfsHash := "QmZULkCELmmk5XNfCgTnCyFgAVxBRBXyDHGGMVoLFLiXEN"

	// Step2: get the ipfs file and get the issue data
	c := "ipfs cat " + ipfsHash
	cmd := exec.Command("sh", "-c", c)
	issue_data, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Get issue data from IPFS: fails: %v\n", err)
	}

	// Step3: unmarshall issue data
	newDeIssue := new(DeIssue)
	err = json.Unmarshal(issue_data, &newDeIssue)
	if err != nil {
		return fmt.Errorf("Can not decode data: %v", err)
	}

	// Step4: write into the local database and mkdir the local path
	newIssue := new(Issue)
	transferDeIssueToIssue(repo, newDeIssue, newIssue)
	has, err2 := x.Get(newIssue)
	if err2 != nil {
		return fmt.Errorf("Can not search the issue: %v\n", err)
	}

	if !has {
		_, err = x.Insert(newIssue)
		if err != nil {
			return fmt.Errorf("Can not add the issue to the server: %v\n", err)
		}
	}

	return nil
}
