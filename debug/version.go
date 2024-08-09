// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package debug

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/focela/aid/encoding/hash"
	"github.com/focela/aid/errors"
)

// BinVersion returns the version of current running binary.
// It uses hash.BKDRHash+BASE36 algorithm to calculate the unique version of the binary.
func BinVersion() string {
	if binaryVersion == "" {
		binaryContent, err := os.ReadFile(selfPath)
		if err != nil {
			panic(errors.Wrapf(err, "failed to read binary file: %s", selfPath))
		}
		binaryVersion = strconv.FormatInt(
			int64(hash.HashBKDR32(binaryContent)),
			36,
		)
	}
	return binaryVersion
}

// BinVersionMd5 returns the version of current running binary.
// It uses MD5 algorithm to calculate the unique version of the binary.
func BinVersionMd5() string {
	if binaryVersionMd5 == "" {
		version, err := md5File(selfPath)
		if err != nil {
			panic(err)
		}
		binaryVersionMd5 = version
	}
	return binaryVersionMd5
}

// md5File encrypts file content of `path` using MD5 algorithms.
func md5File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", errors.Wrapf(err, `failed to open file "%s"`, path)
	}
	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", errors.Wrap(err, `failed to copy file content to MD5 hash`)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
