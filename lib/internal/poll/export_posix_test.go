// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || windows
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris windows

// Export guts for testing on posix.
// Since testing imports os and os imports github.com/blahblah5555/CobaltStrikeParser-Go/lib/internal/poll,
// the github.com/blahblah5555/CobaltStrikeParser-Go/lib/internal/poll tests can not be in package poll.

package poll

func (fd *FD) EOFError(n int, err error) error {
	return fd.eofError(n, err)
}
