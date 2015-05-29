package flavorcommands

import "github.com/codegangsta/cli"

var commandPrefix = "compute flavors"

// Get returns all the commands allowed for a `compute flavors` request.
func Get() []cli.Command {
	return []cli.Command{
		list,
		get,
	}
}
