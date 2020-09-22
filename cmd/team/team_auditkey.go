package team

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
	"samhofi.us/x/infobot/pkg/infobot"
	"samhofi.us/x/infobot/pkg/utils"
	"samhofi.us/x/keybase/v2"
)

var teamAuditKeyCmd = cli.Command{
	Name:      "audit-key",
	Usage:     "Audit a key",
	ArgsUsage: "<team name> <key>",
	Action:    teamAuditKey,
}

func teamAuditKey(c *cli.Context) (err error) {
	var args = c.Args()
	if args.Len() == 0 {
		return fmt.Errorf("missing team name and key")
	}
	if args.Len() == 1 {
		return fmt.Errorf("missing key")
	}

	kb := keybase.New(keybase.SetHomePath(c.Path("home")))

	var (
		teamName = strings.ToLower(args.Get(0))
		key      = args.Get(1)
	)

	// fetch key info from the team's kvstore
	info, err := infobot.FetchKey(kb, teamName, key)
	if err != nil {
		return fmt.Errorf("Error: Unable to fetch key -- %v", err)
	}

	fmt.Fprintf(c.App.Writer, "%s\n", utils.ToJsonPretty(info))
	return nil
}
