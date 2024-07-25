// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcipv4

import (
	"net"

	"github.com/focela/ratcatcher/errors/rcerror"
)

// GetMac retrieves and returns the first mac address of current host.
func GetMac() (mac string, err error) {
	macs, err := GetMacArray()
	if err != nil {
		return "", err
	}
	if len(macs) > 0 {
		return macs[0], nil
	}
	return "", nil
}

// GetMacArray retrieves and returns all the mac address of current host.
func GetMacArray() (macs []string, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		err = rcerror.Wrap(err, `net.Interfaces failed`)
		return nil, err
	}
	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macs = append(macs, macAddr)
	}
	return macs, nil
}
