// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	. "github.com/pingcap/check"
	"github.com/pingcap/errors"
)

var _ = Suite(&testErrorsSuite{})

type testErrorsSuite struct{}

func (s *testErrorsSuite) TestOriginError(c *C) {
	c.Assert(OriginError(nil), IsNil)

	err1 := errors.New("err1")
	c.Assert(OriginError(err1), DeepEquals, err1)

	err2 := errors.Trace(err1)
	c.Assert(OriginError(err2), DeepEquals, err1)

	err3 := errors.Trace(err2)
	c.Assert(OriginError(err3), DeepEquals, err1)
}
