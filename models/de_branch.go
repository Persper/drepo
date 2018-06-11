// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import ()

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
