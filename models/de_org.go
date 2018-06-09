// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import ()

// The org table in the IPFS
type DeOrg struct {
	Description string
}

func transferUserToDeUser(deOrg *DeOrg, org *User) {
	DeOrg.Description = org.Description
}

func deTransferUserToDeUser(deOrg *DeOrg, org *User) {
	org.Description = DeOrg.Description

	//TODO:
	org.NumTeams = 0
	org.NumMembers = 0
}
