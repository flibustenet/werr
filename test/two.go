// Copyright (c) 2023 William Dode. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import (
	"go.flibuste.net/werr"
)

func three() error {
	return werr.Wrapf(two(), "three")
}
func two() error {
	return werr.Wrapf(one(), "two")
}
func one() error {
	return werr.Wrapf(fail(), "one")
}
