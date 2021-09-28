// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package main

import (
	"errors"
	"os"

	"github.com/docker/docker/quota"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var (
	basePath = pflag.String("base", ".", "base path to establish XFS control on")
	dir      = pflag.String("dir", "", "directory to establish the quota on")
	size     = pflag.Uint64("size", 0, "size of the quota to allow for a directory")
)

func main() {
	pflag.Parse()

	if *dir == "" {
		logrus.Fatal("-dir is mandatory")
	}
	if *size == 0 {
		logrus.Fatal("-size is mandatory and must be greater than 0")
	}

	ctrl, err := quota.NewControl(*basePath)
	if errors.Is(err, quota.ErrQuotaNotSupported) {
		logrus.WithError(err).Warnf("no XFS available on %s", *basePath)
		os.Exit(67)
	}
	if err != nil {
		logrus.Fatal(err)
	}
	err = ctrl.SetQuota(*dir, quota.Quota{Size: *size})
	if err != nil {
		logrus.Fatal(err)
	}

}
