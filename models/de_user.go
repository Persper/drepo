// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/// ***** START: DeFollow *****
type DeFollow struct {
	FollowID int64 `xorm:"UNIQUE(follow)"`
}

func transferFollowToDeFollow(follow *Follow, deFollow *DeFollow) {
	deFollow.FollowID = follow.FollowID
}

func transferDeFollowToFollow(user *User, deFollow *DeFollow, follow *Follow) {
	// follow.ID can be generated at any time
	// TODO: follow.ID
	follow.UserID = user.ID
	follow.FollowID = deFollow.FollowID
}

/// ***** END: DeStar *****

/// ***** START: DeStar *****
type DeStar struct {
	RepoID int64 `xorm:"UNIQUE(s)"`
}

func transferStarToDeStar(star *Star, deStar *DeStar) {
	deStar.RepoID = star.RepoID
}

func transferDeStarToStar(user *User, deStar *DeStar, star *Star) {
	// star.ID can be generated at any time
	// TODO: star.ID
	star.UID = user.ID
	star.RepoID = deStar.RepoID
}

/// ***** END: DeStar *****

/// ***** START: DeWatch *****
type DeWatch struct {
	RepoID int64 `xorm:"UNIQUE(watch)"`
}

func transferWatchToDeWatch(watch *Watch, deWatch *DeWatch) {
	deWatch.RepoID = watch.RepoID
}

func transferDeWatchToWatch(user *User, deWatch *DeWatch, watch *Watch) {
	// watch.ID can be generated at any time
	// TODO: watch.ID
	watch.UserID = user.ID
	watch.RepoID = deWatch.RepoID
}

/// ***** END: DeWatch *****

/// ***** START: DeEmailAddress *****
type DeEmailAddress struct {
	Email string `xorm:"UNIQUE NOT NULL"`
}

func transferEmailAddrToDeEmailAddr(emailAddr *EmailAddress, deEmailAddr *DeEmailAddress) {
	deEmailAddr.Email = emailAddr.Email
}

func transferDeEmailAddrToEmailAddr(user *User, deEmailAddr *DeEmailAddress, emailAddr *EmailAddress) {
	// emailAddr.ID can be generated at any time
	// TODO: emailAddr.ID
	emailAddr.UID = user.ID
	emailAddr.Email = deEmailAddr.Email
	emailAddr.IsActivated = true
	if emailAddr.Email == user.Email {
		emailAddr.IsPrimary = true
	} else {
		emailAddr.IsPrimary = false
	}
}

/// ***** END: DeEmailAddress *****

/// ***** START: DePublicKey *****
type DePublicKey struct {
	Name        string     `xorm:"NOT NULL"`
	Fingerprint string     `xorm:"NOT NULL"`
	Content     string     `xorm:"TEXT NOT NULL"`
	Mode        AccessMode `xorm:"NOT NULL DEFAULT 2"`
	Type        KeyType    `xorm:"NOT NULL DEFAULT 1"`
	CreatedUnix int64
	UpdatedUnix int64
}

func transferPubKeyToDePubKey(pubKey *PublicKey, dePubKey *DePublicKey) {
	dePubKey.Name = pubKey.Name
	dePubKey.Fingerprint = pubKey.Fingerprint
	dePubKey.Content = pubKey.Content
	dePubKey.Mode = pubKey.Mode
	dePubKey.Type = pubKey.Type
	dePubKey.CreatedUnix = pubKey.CreatedUnix
	dePubKey.UpdatedUnix = pubKey.UpdatedUnix
}

func transferDePubKeyToPubKey(user *User, dePubKey *DePublicKey, pubKey *PublicKey) {
	// pubKey.ID can be generated at any time
	// TODO: pubKey.ID
	pubKey.OwnerID = user.ID
	pubKey.Name = dePubKey.Name
	pubKey.Fingerprint = dePubKey.Fingerprint
	pubKey.Content = dePubKey.Content
	pubKey.Mode = dePubKey.Mode
	pubKey.Type = dePubKey.Type
	pubKey.CreatedUnix = dePubKey.CreatedUnix
	pubKey.UpdatedUnix = dePubKey.UpdatedUnix
}

/// ***** END: DePublicKey *****

type DeUser struct {
	ID                 int64
	Name               string `xorm:"UNIQUE NOT NULL"`
	FullName           string
	Email              string `xorm:"NOT NULL"`
	Passwd             string `xorm:"NOT NULL"`
	LoginType          LoginType
	LoginSource        int64 `xorm:"NOT NULL DEFAULT 0"`
	LoginName          string
	Location           string
	Website            string
	Rands              string `xorm:"VARCHAR(10)"`
	Salt               string `xorm:"VARCHAR(10)"`
	CreatedUnix        int64
	UpdatedUnix        int64
	LastRepoVisibility bool
	Avatar             string `xorm:"VARCHAR(2048) NOT NULL"`
	AvatarEmail        string `xorm:"NOT NULL"`
	UseCustomAvatar    bool
	EmailAddr          []DeEmailAddress `xorm:"-"`
	PubKey             []DePublicKey    `xorm:"-"`
	StarRepoID         []DeStar         `xorm:"-"`
	WatchRepoID        []DeWatch        `xorm:"-"`
	FollowUserID       []DeFollow       `xorm:"-"`
	// TODO:
	// repo_blacklist
	// team_blacklist
	// org_blacklist
}

func transferUserToDeUser(user *User, deUser *DeUser) error {
	deUser.ID = user.ID
	deUser.Name = user.Name
	deUser.FullName = user.FullName
	deUser.Email = user.Email
	deUser.Passwd = user.Passwd
	deUser.LoginType = user.LoginType
	deUser.LoginSource = user.LoginSource
	deUser.LoginName = user.LoginName
	deUser.Location = user.Location
	deUser.Website = user.Website
	deUser.Rands = user.Rands
	deUser.Salt = user.Salt
	deUser.CreatedUnix = user.CreatedUnix
	deUser.UpdatedUnix = user.UpdatedUnix
	deUser.LastRepoVisibility = user.LastRepoVisibility
	deUser.Avatar = user.Avatar
	deUser.AvatarEmail = user.AvatarEmail
	deUser.UseCustomAvatar = user.UseCustomAvatar

	// ***** START: EmailAddress[] *****
	emailAddresses := make([]EmailAddress, 0)
	if err := x.Find(&emailAddresses, &EmailAddress{UID: user.ID, IsActivated: true}); err != nil {
		return fmt.Errorf("Can not get emailAddresses of the user: %v\n", err)
	}
	for i := range emailAddresses {
		deEmailAddr := new(DeEmailAddress)
		transferEmailAddrToDeEmailAddr(&emailAddresses[i], deEmailAddr)
		deUser.EmailAddr = append(deUser.EmailAddr, *deEmailAddr)
	}
	// ***** END: EmailAddress[] *****

	// ***** START: PubKey[] *****
	publicKeys := make([]PublicKey, 0)
	if err := x.Find(&publicKeys, &PublicKey{OwnerID: user.ID}); err != nil {
		return fmt.Errorf("Can not get publicKeys of the user: %v\n", err)
	}
	for i := range publicKeys {
		dePublicKey := new(DePublicKey)
		transferPubKeyToDePubKey(&publicKeys[i], dePublicKey)
		deUser.PubKey = append(deUser.PubKey, *dePublicKey)
	}
	// ***** END: PubKey[] *****

	// ***** START: Star[] *****
	stars := make([]Star, 0)
	if err := x.Find(&stars, &Star{UID: user.ID}); err != nil {
		return fmt.Errorf("Can not get stars of the user: %v\n", err)
	}
	for i := range stars {
		deStar := new(DeStar)
		transferStarToDeStar(&stars[i], deStar)
		deUser.StarRepoID = append(deUser.StarRepoID, *deStar)
	}
	// ***** END: Star[] *****

	// ***** START: Watch[] *****
	watches := make([]Watch, 0)
	if err := x.Find(&watches, &Watch{UserID: user.ID}); err != nil {
		return fmt.Errorf("Can not get watches of the user: %v\n", err)
	}
	for i := range watches {
		deWatch := new(DeWatch)
		transferWatchToDeWatch(&watches[i], deWatch)
		deUser.WatchRepoID = append(deUser.WatchRepoID, *deWatch)
	}
	// ***** END: Watch[] *****

	// ***** START: Follow[] *****
	follows := make([]Follow, 0)
	if err := x.Find(&follows, &Follow{UserID: user.ID}); err != nil {
		return fmt.Errorf("Can not get follows of the user: %v\n", err)
	}
	for i := range follows {
		deFollow := new(DeFollow)
		transferFollowToDeFollow(&follows[i], deFollow)
		deUser.FollowUserID = append(deUser.FollowUserID, *deFollow)
	}
	// ***** END: Follow[] *****

	return nil
}

/// Prerequisite: all repos exist or repo.id[] exists
func transferDeUserToUser(deUser *DeUser, user *User) error {
	user.ID = deUser.ID
	user.Name = deUser.Name
	user.FullName = deUser.FullName
	user.Email = deUser.Email
	user.Passwd = deUser.Passwd
	user.LoginType = deUser.LoginType
	user.LoginSource = deUser.LoginSource
	user.LoginName = deUser.LoginName
	user.Location = deUser.Location
	user.Website = deUser.Website
	user.Rands = deUser.Rands
	user.Salt = deUser.Salt
	user.CreatedUnix = deUser.CreatedUnix
	user.UpdatedUnix = deUser.UpdatedUnix
	user.LastRepoVisibility = deUser.LastRepoVisibility
	user.Avatar = deUser.Avatar
	user.AvatarEmail = deUser.AvatarEmail
	user.UseCustomAvatar = deUser.UseCustomAvatar

	// recovery deUser to user
	user.Type = USER_TYPE_INDIVIDUAL
	user.LowerName = strings.ToLower(user.Name)
	user.MaxRepoCreation = -1
	user.IsAdmin = false

	// org
	user.Description = ""
	user.NumTeams = 0
	user.NumMembers = 0

	// TODO: not sure
	// TODO: user.UportId
	user.IsActive = true
	user.AllowGitHook = false
	user.AllowImportLocal = false
	user.ProhibitLogin = false

	// ***** START: EmailAddress[] *****
	for i := range deUser.EmailAddr {
		emailAddr := new(EmailAddress)
		transferDeEmailAddrToEmailAddr(user, &deUser.EmailAddr[i], emailAddr)
		has, err := x.Get(emailAddr)
		if err != nil {
			return fmt.Errorf("Can not search the emailAddr: %v\n", err)
		}
		if !has {
			_, err = x.Insert(emailAddr)
			if err != nil {
				return fmt.Errorf("Can not add the email to the server: %v\n", err)
			}
		}
	}
	// ***** END: EmailAddress[] *****

	// ***** START: PubKey[] *****
	for i := range deUser.PubKey {
		pubKey := new(PublicKey)
		transferDePubKeyToPubKey(user, &deUser.PubKey[i], pubKey)
		has, err := x.Get(pubKey)
		if err != nil {
			return fmt.Errorf("Can not search the pubKey: %v\n", err)
		}
		if !has {
			_, err = x.Insert(pubKey)
			if err != nil {
				return fmt.Errorf("Can not add the pubKey to the server: %v\n", err)
			}
		}
	}
	// ***** END: PubKey[] *****

	// ***** START: Star[] *****
	for i := range deUser.StarRepoID {
		star := new(Star)
		transferDeStarToStar(user, &deUser.StarRepoID[i], star)
		has, err := x.Get(star)
		if err != nil {
			return fmt.Errorf("Can not search the star: %v\n", err)
		}
		if !has {
			_, err = x.Insert(star)
			if err != nil {
				return fmt.Errorf("Can not add the star to the server: %v\n", err)
			}
		}
	}
	user.NumStars = len(deUser.StarRepoID)
	// ***** END: Star[] *****

	// ***** START: Watch[] *****
	for i := range deUser.WatchRepoID {
		watch := new(Watch)
		transferDeWatchToWatch(user, &deUser.WatchRepoID[i], watch)
		has, err := x.Get(watch)
		if err != nil {
			return fmt.Errorf("Can not search the watch: %v\n", err)
		}
		if !has {
			_, err = x.Insert(watch)
			if err != nil {
				return fmt.Errorf("Can not add the watch to the server: %v\n", err)
			}
		}
	}
	// ***** END: Watch[] *****

	// ***** START: Follow[] *****
	for i := range deUser.FollowUserID {
		follow := new(Follow)
		transferDeFollowToFollow(user, &deUser.FollowUserID[i], follow)
		has, err := x.Get(follow)
		if err != nil {
			return fmt.Errorf("Can not search the follow: %v\n", err)
		}
		if !has {
			_, err = x.Insert(follow)
			if err != nil {
				return fmt.Errorf("Can not add the follow to the server: %v\n", err)
			}
			// Calculate the following of the followed user
			followingUser := &User{ID: deUser.FollowUserID[i].FollowID}
			has, err = x.Get(followingUser)
			if err != nil {
				return fmt.Errorf("Can not search the followingUser: %v\n", err)
			}
			if has {
				if _, err = x.Exec("UPDATE `user` SET num_following=num_following+1 WHERE id=?", followingUser.ID); err != nil {
					return fmt.Errorf("update num_following: %v\n", err)
				}
			}
		}
	}
	user.NumFollowers = len(deUser.FollowUserID)
	// ***** END: Follow[] *****

	// ***** START: NumFollowing *****
	follow := new(Follow)
	total, err := x.Where("follow_id = ?", user.ID).Count(follow)
	if err != nil {
		return fmt.Errorf("Can not get follow issues: %v\n", err)
	}
	user.NumFollowing = int(total)
	// ***** END: NumFollowing *****

	// ***** START: NumRepos *****
	// NumRepos will be updated when the corresponding repo is added
	user.NumRepos = 0
	// ***** END: NumRepos *****

	return nil
}

/// Push the user info to IPFS and record the new ipfsHash in the blockchain
/// pushMode: 0 - register; 1 - update; 2 - delete;
func PushUserInfo(user *User, pushMode int) (err error) {
	if user.IsOrganization() {
		return nil
	}
	if !canPushToBlockchain(user) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step1: register/deregister the user if it does not exist
	if pushMode == 0 {
		//err = registerName
	} else if pushMode == 2 {
		//err = deregisterName
	}

	// Step 2: Encode user data into JSON format
	deUser := new(DeUser)
	transferUserToDeUser(user, deUser)
	user_data, err := json.Marshal(deUser)
	if err != nil {
		return fmt.Errorf("Can not encode user data: %v\n", err)
	}

	// Step 3: Put the encoded data into IPFS
	c := fmt.Sprintf("echo '%s' | ipfs add ", user_data)
	cmd := exec.Command("sh", "-c", c)
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Ca not push user to IPFS: %v\n", err)
	}
	ipfsHash := strings.Split(string(out), " ")[1]

	// Step4: Modify the ipfsHash in the smart contract
	// TODO: setUserInfo(ipfsHash)
	ipfsHash = ipfsHash
	fmt.Println("Push the user file to the IPFS: " + ipfsHash)

	return nil
}

/// Get the new ipfsHash from the blockchain and get the user info from IPFS
func GetUserInfo(uportID string, ipfsHash string) (user *User, err error) {
	// Step1: get the ipfs file and get the user data
	c := "ipfs cat " + ipfsHash
	cmd := exec.Command("sh", "-c", c)
	user_data, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Can not get user data from IPFS: %v\n", err)
	}

	// Step2: unmarshall user data
	newDeUser := new(DeUser)
	err = json.Unmarshal(user_data, &newDeUser)
	if err != nil {
		return nil, fmt.Errorf("Can not decode user data: %v\n", err)
	}

	// Step3: write into the local database and mkdir the user path
	newUser := new(User)
	transferDeUserToUser(newDeUser, newUser)
	newUser.UportId = uportID
	has, err := x.Get(newUser)
	if err != nil {
		return nil, fmt.Errorf("Can not search the user: %v\n", err)
	}
	if !has {
		sess := x.NewSession()
		defer sess.Close()
		if err = sess.Begin(); err != nil {
			return nil, err
		}

		if _, err = sess.Insert(newUser); err != nil {
			return nil, err
		} else if err = os.MkdirAll(UserPath(newUser.Name), os.ModePerm); err != nil {
			return nil, err
		}

		return newUser, sess.Commit()
	}

	return newUser, nil
}

/// The user button: push the user info and all related tables to IPFS
func PushUserAndOwnedRepos(contextUser *User) (err error) {
	if !canPushToBlockchain(contextUser) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step1: get the corresponding user.
	var user *User
	user = &User{ID: contextUser.ID}
	_, err = x.Get(user)
	if err != nil {
		return fmt.Errorf("Can not get user data: %v\n", err)
	}

	// Step2: push user to the blockchain and the IPFS
	// TODO: check update or create
	if err := PushUserInfo(user, 1); err != nil {
		return err
	}

	// Step3: push the owned repo
	repos := make([]Repository, 0)
	if err = x.Find(&repos, &Repository{OwnerID: user.ID}); err != nil {
		return fmt.Errorf("Can not get owned repos of the user: %v\n", err)
	}

	for i := range repos {
		if err = PushRepoInfo(user, &repos[i]); err != nil {
			return err
		}

		issues := make([]Issue, 0)
		if err = x.Find(&issues, &Issue{RepoID: repos[i].ID}); err != nil {
			return fmt.Errorf("Can not get issues of the repo: %v\n", err)
		}
		for j := range issues {
			if err = PushIssueInfo(user, &issues[j]); err != nil {
				return err
			}

			prs := make([]PullRequest, 0)
			if err = x.Find(&prs, &PullRequest{IssueID: issues[j].ID}); err != nil {
				return fmt.Errorf("Can not get pull_request of the repo: %v\n", err)
			}
			for k := range prs {
				if err = PushPullInfo(user, &prs[k]); err != nil {
					return err
				}
			}
		}

		branches := make([]ProtectBranch, 0)
		if err = x.Find(&branches, &ProtectBranch{RepoID: repos[i].ID}); err != nil {
			return fmt.Errorf("Can not get branches of the repo: %v\n", err)
		}
		for j := range branches {
			if err = PushBranchInfo(user, &branches[j]); err != nil {
				return err
			}
		}
	}

	// Step4: push the owned org

	return nil
}

/// The user button: get the user info and all related tables to IPFS
func GetUserAndOwnedRepos(uportID string) (err error) {
	// Just for test
	/*
		var testUser *User
		testIpfsHash := "QmTMU8bqRX1YcQvbe7AwjUca3U2KAFsfn3i9YwfTp1gY3C"
		if testUser, err = GetUserInfo(uportID, testIpfsHash); err != nil {
			return err
		}

		testIpfsHash = "Qmcodn79uJF7GE9vqQFhD9u4cRMFe5vAMHV7zieHvepBLp"
		if err := GetRepoAndRelatedTables(testUser, testIpfsHash); err != nil {
			return err
		}

		testIpfsHash = "QmXWQDNWkN4j72vM2iriPPxEt6Kz1b6fn42x2rTWFhiZBy"
		if err := GetOrgAndRelatedTables(testUser, testIpfsHash); err != nil {
			return err
		}
		return nil
	*/

	// Step1: get the user table
	// TODO: get userIpfsHash from the blockchain
	userIpfsHash := ""
	var user *User
	if user, err = GetUserInfo(uportID, userIpfsHash); err != nil {
		return err
	}

	// Step2: get the owned repo
	// TODO: get repoIpfsHashes from the blockchain
	repoIpfsHashes := make([]string, 0)
	for i := range repoIpfsHashes {
		if err := GetRepoAndRelatedTables(user, repoIpfsHashes[i]); err != nil {
			return err
		}
	}

	// Step3: get the owned org
	// TODO: get orgIpfsHashes from the blockchain
	orgIpfsHashes := make([]string, 0)
	for i := range orgIpfsHashes {
		if err := GetOrgAndRelatedTables(user, orgIpfsHashes[i]); err != nil {
			return err
		}
	}
	return nil
}

/// Uncompleted
/// TODO: Push the user info and all related tables to IPFS
/*func PushUserAllInfos(contextUser *User) (err error) {
	if !canPushToBlockchain(contextUser) {
		return fmt.Errorf("The user can not push to the blockchain")
	}

	// Step0: get the corresponding user.
	var user *User
	user = &User{ID: contextUser.ID}
	_, err = x.Get(user)
	if err != nil {
		return fmt.Errorf("Can not get user data: %v", err)
	}

	// Step1: push user: check update or create
	if err := PushUserInfo(user, 1); err != nil {
		return err
	}

	// Step2: push the related orgs
	orgUsers := make([]OrgUser, 0)
	if err = x.Find(&orgUsers, &OrgUser{Uid: user.ID}); err != nil {
		return fmt.Errorf("Can not get orgUsers of the user: %v", err)
	}
	for i := range orgUsers {
		var org *User
		org = &User{ID: orgUsers[i].ID}
		hasOrg, err := x.Get(org)
		if err != nil {
			return fmt.Errorf("Can not get org data: %v", err)
		}
		if hasOrg {
			if err = PushOrgInfo(user, org); err != nil {
				return err
			}
			// TODO: only owner?
			if err = PushOrgUserInfo(user, org, &orgUsers[i]); err != nil {
				return err
			}

			teams := make([]Team, 0)
			if err = x.Find(&teams, &Team{OrgID: org.ID}); err != nil {
				return fmt.Errorf("Can not get teams of the user: %v", err)
			}
			for j := range teams {
				if err = PushTeamInfo(user, &teams[j]); err != nil {
					return err
				}
			}
		}
	}

	// Step3: push the related repo
	accesses := make([]Access, 0)
	if err = x.Find(&accesses, &Access{UserID: user.ID}); err != nil {
		return fmt.Errorf("Can not get accesses of the user: %v", err)
	}
	for i := range accesses {
		var repo *Repository
		repo = &Repository{ID: accesses[i].RepoID}
		hasRepo, err := x.Get(repo)
		if err != nil {
			return fmt.Errorf("Can not get repo data: %v", err)
		}
		if hasRepo {
			if err = PushRepoInfo(user, repo); err != nil {
				return err
			}
		}

		issues := make([]Issue, 0)
		if err = x.Find(&issues, &Issue{RepoID: accesses[i].RepoID}); err != nil {
			return fmt.Errorf("Can not get issues of the repo: %v", err)
		}

		for j := range issues {
			if err = PushIssueInfo(user, &issues[j]); err != nil {
				return err
			}

			pulls := make([]PullRequest, 0)
			if err = x.Find(&pulls, &PullRequest{IssueID: issues[j].ID}); err != nil {
				return fmt.Errorf("Can not get pulls of the repo: %v", err)
			}

			for k := range pulls {
				if err = PushPullInfo(user, &pulls[k]); err != nil {
					return err
				}
			}
		}

		branches := make([]ProtectBranch, 0)
		if err = x.Find(&branches, &ProtectBranch{RepoID: accesses[i].RepoID}); err != nil {
			return fmt.Errorf("Can not get branches of the repo: %v", err)
		}

		for j := range branches {
			if err = PushBranchInfo(user, &branches[j]); err != nil {
				return err
			}
		}
	}

	return nil
}*/
