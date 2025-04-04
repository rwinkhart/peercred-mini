// Copyright (c) 2021 AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package peercred maps from a net.Conn to information about the
// other side of the connection, using various OS-specific facilities.
package peercred

import (
	"net"
)

// Creds are the peer credentials.
type Creds struct {
	pid int
	uid string
}

func (c *Creds) PID() (pid int) {
	return c.pid
}

// UserID returns the userid (or Windows SID) that owns the other side
// of the connection.
// The returned string is suitable to passing to os/user.LookupId.
func (c *Creds) UserID() (uid string) {
	return c.uid
}

func osGet(c net.Conn) (*Creds, error) {
	return getUnix(c.(*net.UnixConn))
}

// Get returns the peer credentials for c.
func Get(c net.Conn) *Creds {
	creds, _ := osGet(c)
	return creds
}
