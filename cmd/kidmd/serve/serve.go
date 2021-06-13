/*
 * SPDX-License-Identifier: AGPL-3.0-or-later
 * Copyright 2021 Kopano and its licensors
 */

package serve

import (
	"github.com/libregraph/idm/cmd/idmd/serve"
	"github.com/libregraph/idm/server/handler/ldif"
	"github.com/spf13/cobra"
)

var (
	DefaultEnvBase = "KIDMD_"
)

func setDefaults() {
	serve.DefaultEnvBase = DefaultEnvBase
}

func setIndexes() {
	ldif.AddIndexAttribute("kopanoAccount", "eq")
	ldif.AddIndexAttribute("kopanoAliases", "eq")
	ldif.AddIndexAttribute("kopanoViewPrivilege", "eq")
}

func CommandServe() *cobra.Command {
	setDefaults()
	setIndexes()

	serveCmd := serve.CommandServe()

	return serveCmd
}
