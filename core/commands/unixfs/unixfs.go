package unixfs

import (
	cmds "github.com/ipfs/go-ipfs/commands"
	e "github.com/ipfs/go-ipfs/core/commands/e"

	"gx/ipfs/QmPMeikDc7tQEDvaS66j1bVPQ2jBkvFwz3Qom5eA5i4xip/go-ipfs-cmds"
)

var UnixFSCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "Interact with IPFS objects representing Unix filesystems.",
		ShortDescription: `
'ipfs file' provides a familiar interface to file systems represented
by IPFS objects, which hides ipfs implementation details like layout
objects (e.g. fanout and chunking).
`,
		LongDescription: `
'ipfs file' provides a familiar interface to file systems represented
by IPFS objects, which hides ipfs implementation details like layout
objects (e.g. fanout and chunking).
`,
	},

	Subcommands: map[string]*cmds.Command{
		"ls": LsCmd,
	},
}

// copy+pasted from ../commands.go
func unwrapOutput(i interface{}) (interface{}, error) {
	var (
		ch <-chan interface{}
		ok bool
	)

	if ch, ok = i.(<-chan interface{}); !ok {
		return nil, e.TypeErr(ch, i)
	}

	return <-ch, nil
}
