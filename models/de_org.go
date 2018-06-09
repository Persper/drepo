// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import ()

// The org table in the IPFS
type DeOrg struct {
	Description string
}

func transferOrgToDeOrg(deOrg *DeOrg, org *User) {
	deOrg.Description = org.Description
}

func deTransferOrgToDeOrg(deOrg *DeOrg, org *User) {
	org.Description = deOrg.Description

	//TODO:
	org.NumTeams = 0
	org.NumMembers = 0
}
