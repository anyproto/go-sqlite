// Copyright 2021 Roxy Light
// SPDX-License-Identifier: ISC

package main

import "github.com/anyproto/go-sqlite"

func main() {
	var db *sqlite.Conn
	var err error
	db, err = sqlite.OpenConn(":memory:", 0)
	_, _ = db, err
}
