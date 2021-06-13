/*
 * SPDX-License-Identifier: AGPL-3.0-or-later
 * Copyright 2021 Kopano and its licensors
 */

package gen

import (
	"github.com/libregraph/idm/cmd/idmd/gen"
	"github.com/libregraph/idm/cmd/idmd/gen/newusers"
	"github.com/spf13/cobra"
)

// Override default idm template to extend with Kopano Groupware specific
// attributes.
var DefaultLDIFUserTemplate = `<<- /* */ ->>
dn: uid=<<.entry.Name>>,<<.BaseDN>>
objectClass: posixAccount
objectClass: top
objectClass: inetOrgPerson
objectClass: kopano-user
uid: <<.entry.Name>>
uidNumber: <<with .detail.uidNumber>><<.>><<else>><<AutoIncrement>><<end>>
<<- with .detail.gidNumber>>
gidNumber: <<.>>
<<- end>>
<<- with .detail.userPassword>>
userPassword: <<.>>
<<- end>>
mail: <<.entry.Name>>@{{.MailDomain}}
<<- range .detail.mail>>
kopanoAliases: <<.>>
<<- end>>
cn: <<.detail.cn>>
<<- with .detail.givenName>>
givenName: <<.>>
<<- end>>
<<- with .detail.sn>>
sn: <<.>>
<<- end>>
kopanoAccount: 1
kopanoAdmin: 0
`

func setDefaults() {
	newusers.DefaultLDIFUserTemplate = DefaultLDIFUserTemplate
}

func CommandServe() *cobra.Command {
	setDefaults()

	genCmd := gen.CommandGen()

	return genCmd
}
