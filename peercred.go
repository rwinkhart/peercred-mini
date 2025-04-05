// Copyright (c) 2021-2025 AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package peercred maps from a net.Conn to information about the
// other side of the connection, using various OS-specific facilities.
package peercred

// Creds are the peer credentials.
type Creds struct {
	pid int
	uid string
}

// PID returns the process ID associated with the other side of the connection.
func (c *Creds) PID() (pid int) {
	return c.pid
}

// UID returns the userid (or Windows SID) that owns the other side of the connection.
// The returned string is suitable to passing to os/user.LookupId.
func (c *Creds) UID() (uid string) {
	return c.uid
}
