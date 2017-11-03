// Code generated by mockery v1.0.0
package netmock

import "context"
import "github.com/stretchr/testify/mock"
import "net"

// Dialer is an autogenerated mock type for the Dialer type
type Dialer struct {
	mock.Mock
}

// DialContext provides a mock function with given fields: ctx, _a1, addr
func (_m *Dialer) DialContext(ctx context.Context, _a1 string, addr string) (net.Conn, error) {
	ret := _m.Called(ctx, _a1, addr)

	var r0 net.Conn
	if rf, ok := ret.Get(0).(func(context.Context, string, string) net.Conn); ok {
		r0 = rf(ctx, _a1, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Conn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, _a1, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}