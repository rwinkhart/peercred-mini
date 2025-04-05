//go:build !windows

// Copyright (c) 2021-2025 AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package peercred

import "net"

// Get returns the peer credentials for c.
func Get(c net.Conn) *Creds {
	creds, _ := get(c.(*net.UnixConn))
	return creds
}
