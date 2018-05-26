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
		fmt.Println("Push to IPFS: mv fails")
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
		fmt.Println("Push to IPFS: " + ipfs_hash)
	}

	// Transform .git to sample.git
	cmd = exec.Command("mv", tmp_name, paths[len(paths)-1])
	cmd.Dir = parent_path
	err = cmd.Run()
	if err != nil {
		fmt.Println("Push to IPFS: the second mv fails")
	}

	// Record the ipfs_hash
	txt_path := path + "/ipfs_hash"
	//fmt.Println("output path: " + txt_path)
	err = ioutil.WriteFile(txt_path, []byte(ipfs_hash), 0666)
	if err != nil {
		fmt.Println("Push to IPFS: record ipfs_hash fails")
	}
}

/* Push user-repo releations in database to ipfs. */
/*func Push_User_Repo_To_IPFS(c *context.Context) {
	uport_id := "2ouSwTJJwTZixoXL6QmizVWEGzspxXQQ2hA"
	u := &models.User{
		UportId: uport_id,
	}
}*/
