// Copyright (c) 2018 David Crawshaw <david@zentus.com>
// Copyright (c) 2021 Roxy Light <roxy@zombiezen.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//
// SPDX-License-Identifier: ISC

package sqlitefile

import (
	"io"
	"os"
	"testing"

	"crawshaw.io/iox/ioxtest"
	"github.com/anyproto/go-sqlite"
)

var _ interface {
	io.Reader
	io.Writer
	io.Seeker
	io.Closer
} = (*File)(nil)

func TestFileRand(t *testing.T) {
	conn, err := sqlite.OpenConn(":memory:", 0)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	f1, err := NewFile(conn)
	if err != nil {
		t.Fatal(err)
	}
	f2, err := os.CreateTemp("", "sqlitex")
	if err != nil {
		t.Fatal(err)
	}
	ft := &ioxtest.Tester{F1: f1, F2: f2, T: t}
	ft.Run()
}
