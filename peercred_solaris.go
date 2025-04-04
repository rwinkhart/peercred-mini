// Copyright (c) 2021 AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package peercred

import (
	"fmt"
	"net"
	"strconv"

	"github.com/rwinkhart/sys/unix"
)

func getUnix(c *net.UnixConn) (*Creds, error) {
	raw, err := c.SyscallConn()
	if err != nil {
		return nil, fmt.Errorf("SyscallConn: %w", err)
	}

	var creds *unix.Ucred
	cerr := raw.Control(func(fd uintptr) {
		creds, err = unix.GetPeerUcred(fd)
		if err != nil {
			err = fmt.Errorf("unix.GetPeerUcred: %w", err)
			return
		}
	})
	if cerr != nil {
		return nil, fmt.Errorf("raw.Control: %w", cerr)
	}
	if err != nil {
		return nil, err
	}
	return &Creds{
		pid: creds.Getpid(),
		uid: strconv.FormatUint(uint64(creds.Geteuid()), 10),
	}, nil
}
