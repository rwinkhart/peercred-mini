# peercred-mini
A simple module for getting the PID and UID of processes accessing UNIX domain sockets and Windows named pipes.

# IMPORTANT - READ FOR FreeBSD+Windows SUPPORT!
There are replacements in the `go.mod` for this module. Make sure you maintain those
same replacements in your importing module, otherwise FreeBSD and Windows support will break
(support for those operating systems requires patched modules).
