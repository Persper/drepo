// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"

	"github.com/gogs/gogs/models/ipfs"
)

// Push the repo info to IPFS and record the new ipfsHash in the blockchain
func PushRepoInfo(owner *User, repo *Repository) (err error) {
	// admin or writer
	return nil
}

// Get the new ipfsHash from the blockchain and get the repo info from IPFS
func GetRepoInfo(owner *User, repo *Repository) (err error) {
	return nil
}

// Push the repo content to IPFS and record the new ipfsHash in the blockchain
func PushRepoContent(owner *User, repoPath string) (err error) {
	if !canPushToBlockchain(owner) {
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
func GetRepoContent(owner *User) (err error) {
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
