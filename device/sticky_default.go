//go:build !linux

package device

import (
	"github.com/icewst/wireguard-go/conn"
	"github.com/icewst/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
