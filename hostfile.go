package hostfile

import (
	"fmt"

	"github.com/lextoumbourou/goodhosts"
	"github.com/mholt/caddy/caddy/setup"
	"github.com/mholt/caddy/middleware"
)

func Setup(c *setup.Controller) (middleware.Middleware, error) {
	any := false
	for c.Next() {
		if any {
			return nil, c.Err("hostfile cannot be specified more than once")
		}
		args := c.RemainingArgs()
		if len(args) != 0 {
			fmt.Println(args)
			return nil, c.ArgErr()
		}
		if c.NextBlock() {
			return nil, c.Err("hostfile accepts no arguments")
		}
		if c.Host == "localhost" {
			continue
		}
		//TODO: filter raw ips? other patterns?
		if err := addHost(c.Host); err != nil {
			return nil, err
		}
		c.Shutdown = append(c.Shutdown, func() error { return removeHost(c.Host) })
	}
	return nil, nil
}

const local = "127.0.0.1"

func addHost(host string) error {
	hosts, err := goodhosts.NewHosts()
	if err != nil {
		return err
	}
	if hosts.Has(local, host) {
		return nil
	}

	hosts.Add(local, host)
	err = hosts.Flush()
	return err
}

func removeHost(host string) error {
	hosts, err := goodhosts.NewHosts()
	if err != nil {
		return err
	}
	if !hosts.Has(local, host) {
		return nil
	}
	hosts.Remove(local, host)
	err = hosts.Flush()
	return err
}
