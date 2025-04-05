// Copyright (c) 2021-2025 AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package peercred

import (
	"fmt"

	"github.com/rwinkhart/sys-freebsd-13-xucred/windows"
)

// Get returns the peer credentials for c.
func Get(h windows.Handle) *Creds {
	creds, _ := get(&h)
	return creds
}

func get(h *windows.Handle) (*Creds, error) {
	// get PID
	var clientPID uint32
	windows.GetNamedPipeClientProcessId(*h, &clientPID)

	// get SID (UID)
	var access uint32 = windows.PROCESS_QUERY_INFORMATION
	hProcess, herr := windows.OpenProcess(access, false, clientPID)
	defer windows.CloseHandle(hProcess)
	var hToken windows.Token
	windows.OpenProcessToken(hProcess, windows.TOKEN_QUERY, &hToken)
	defer hToken.Close()
	tokenUser, terr := hToken.GetTokenUser()
	if terr != nil {
		return nil, fmt.Errorf("windows.OpenProcessToken: %w", terr)
	}
	if herr != nil {
		return nil, fmt.Errorf("windows.OpenProcess: %w", herr)
	}

	return &Creds{
		PID: int(clientPID),
		UID: tokenUser.User.Sid.String(),
	}, nil
}
