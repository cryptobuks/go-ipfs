package commands

import (
	"fmt"

	"gx/ipfs/QmPMeikDc7tQEDvaS66j1bVPQ2jBkvFwz3Qom5eA5i4xip/go-ipfs-cmds"

	cmds "github.com/ipfs/go-ipfs/commands"
)

var daemonShutdownCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "Shut down the ipfs daemon",
	},
	Run: func(req cmds.Request, res cmds.Response) {
		nd, err := req.InvocContext().GetNode()
		if err != nil {
			res.SetError(err, cmds.ErrNormal)
			return
		}

		if nd.LocalMode() {
			res.SetError(fmt.Errorf("daemon not running"), cmds.ErrClient)
			return
		}

		if err := nd.Process().Close(); err != nil {
			log.Error("error while shutting down ipfs daemon:", err)
		}

		res.SetOutput(nil)
	},
}
