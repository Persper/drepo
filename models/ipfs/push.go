// Copyright 2018 Persper Foundation
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ipfs

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

/* Push one repo in server to ipfs. For example, go0.git in server. */
func Push_Repo_To_IPFS(path string) {
	paths := strings.Split(path, "/")
	parent_path := strings.TrimSuffix(path, paths[len(paths)-1])
	//fmt.Println(path, parent_path)
	repo_name := strings.Split(paths[len(paths)-1], ".")
	//fmt.Println(repo_name[0], repo_name[1])
	tmp_name := "." + repo_name[1]

	// Transform sample.git to .git
	cmd := exec.Command("mv", paths[len(paths)-1], tmp_name)
	cmd.Dir = parent_path
	err := cmd.Run()
	if err != nil {
		fmt.Println("Push_Repo_To_IPFS: mv fails")
		return
	}

	// Push to IPFS
	cmd = exec.Command("git", "push", "ipfs::", "master") //the remote "ipfs_repo" only in newly generated repo, not in fork repo
	cmd.Dir = parent_path
	out, ipfs_err := cmd.CombinedOutput()
	out_str := string(out)
	var ipfs_hash string
	if ipfs_err != nil {
		fmt.Println(ipfs_err)
	} else {
		fmt.Println(out_str)
		id := strings.Index(out_str, "to IPFS as")
		ipfs_hash = out_str[id+16 : id+62]
		fmt.Println("Push_Repo_To_IPFS: " + ipfs_hash)
	}

	// Transform .git to sample.git
	cmd = exec.Command("mv", tmp_name, paths[len(paths)-1])
	cmd.Dir = parent_path
	err = cmd.Run()
	if err != nil {
		fmt.Println("Push_Repo_To_IPFS: the second mv fails")
	}

	// Record the ipfs_hash
	txt_path := path + "/ipfs_hash"
	err = ioutil.WriteFile(txt_path, []byte(ipfs_hash), 0666)
	if err != nil {
		fmt.Println("Push_Repo_To_IPFS: record ipfs_hash fails")
	}
}

/* Push one file to IPFS. */
func Push_File_To_IPFS(file_name string) {
	cmd := exec.Command("pwd")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Push_File_To_IPFS: pwd fails")
	}

	cmd = exec.Command("ipfs", "add", file_name)
	cmd.Dir = string(out[0 : len(out)-1])
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Push_File_To_IPFS: ipfs add fails")
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
