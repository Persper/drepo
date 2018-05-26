// Copyright 2018 The Persper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	//"github.com/go-xorm/xorm"
	//"github.com/gogits/gogs/models/ipfs"
)

func GetUserInfo(uportid string) {
	/* Get the corresponding user.  */
	var user *User
	user = &User{UportId: uportid}
	hasUser, err := x.Get(user)

	if err != nil {
		fmt.Println("Error: get user!")
	}

	if hasUser {
		/* Get the access control between user and repo. */
		accesses := make([]Access, 0)
		err = x.Where("user_id = ?", user.ID).Find(&accesses)
		if err != nil {
			fmt.Println("Access Error!")
		}
		for i := 0; i < len(accesses); i++ {
			fmt.Println(accesses[i].ID)
		}

	} else {
		fmt.Println("Error: no this user!")
	}
}
