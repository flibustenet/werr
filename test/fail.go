// Copyright (c) 2023 William Dode. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import "errors"

func fail() error {
	return errors.New("fail")
}
