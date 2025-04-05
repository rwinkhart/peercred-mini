// Copyright (c) 2021-2025 AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package peercred maps from a net.Conn to information about the
// other side of the connection, using various OS-specific facilities.
package peercred

// Creds are the peer credentials.
type Creds struct {
	PID int
	UID string
}
