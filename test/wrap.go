// Copyright (c) 2023 William Dode. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import (
	"go.flibuste.net/werr"
)

func wrap() error {
	return werr.Wrap(fail())
}
