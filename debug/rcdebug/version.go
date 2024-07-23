// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcdebug

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/focela/ratcatcher/encoding/rchash"
	"github.com/focela/ratcatcher/errors/rcerror"
)

// BinVersion returns the version of current running binary.
// It uses rchash.BKDRHash+BASE36 algorithm to calculate the unique version of the binary.
func BinVersion() string {
	if binaryVersion == "" {
		binaryContent, _ := os.ReadFile(selfPath)
		binaryVersion = strconv.FormatInt(
			int64(rchash.BKDR(binaryContent)),
			36,
		)
	}
	return binaryVersion
}

// BinVersionMd5 returns the version of current running binary.
// It uses MD5 algorithm to calculate the unique version of the binary.
func BinVersionMd5() string {
	if binaryVersionMd5 == "" {
		binaryVersionMd5, _ = md5File(selfPath)
	}
	return binaryVersionMd5
}

// md5File encrypts file content of `path` using MD5 algorithms.
func md5File(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		err = rcerror.Wrapf(err, `os.Open failed for name "%s"`, path)
		return "", err
	}
	defer f.Close()
	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		err = rcerror.Wrap(err, `io.Copy failed`)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
