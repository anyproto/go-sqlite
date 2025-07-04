// Copyright 2021 Roxy Light
// SPDX-License-Identifier: ISC

package sqlitex_test

import (
	"context"
	"fmt"

	"github.com/anyproto/go-sqlite"
	"github.com/anyproto/go-sqlite/sqlitex"
)

func ExampleExecute() {
	conn, err := sqlite.OpenConn(":memory:")
	if err != nil {
		// handle err
	}

	if err := sqlitex.Execute(conn, "CREATE TABLE t (a, b, c, d);", nil); err != nil {
		// handle err
	}

	err = sqlitex.Execute(conn, "INSERT INTO t (a, b, c, d) VALUES (?, ?, ?, ?);", &sqlitex.ExecOptions{
		Args: []any{"a1", 1, 42, 1},
	})
	if err != nil {
		// handle err
	}

	var a []string
	var b []int64
	err = sqlitex.Execute(conn, "SELECT a, b FROM t WHERE c = ? AND d = ?;", &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			a = append(a, stmt.ColumnText(0))
			b = append(b, stmt.ColumnInt64(1))
			return nil
		},
		Args: []any{42, 1},
	})
	if err != nil {
		// handle err
	}

	fmt.Println(a, b)
	// Output:
	// [a1] [1]
}

func ExampleSave() {
	doWork := func(conn *sqlite.Conn) (err error) {
		defer sqlitex.Save(conn)(&err)

		// ... do work in the transaction
		return nil
	}
	_ = doWork
}

func ExamplePool() {
	// Open a pool.
	dbpool, err := sqlitex.NewPool("foo.db", sqlitex.PoolOptions{})
	if err != nil {
		// handle err
	}
	defer func() {
		if err := dbpool.Close(); err != nil {
			// handle err
		}
	}()

	// While handling a request:
	ctx := context.TODO()
	conn, err := dbpool.Take(ctx)
	if err != nil {
		// handle err
	}
	defer dbpool.Put(conn)
}
