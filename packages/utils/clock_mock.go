/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package utils

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockClock is an autogenerated mock type for the Clock type
type MockClock struct {
	mock.Mock
}

// Now provides a mock function with given fields:
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}
