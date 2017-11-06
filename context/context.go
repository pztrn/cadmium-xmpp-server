// Cadmium XMPP server.
// Copyright (c) 2017, Cadmium developers and contributors.
//
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package context

import (
	// stdlib
	"os"

	// local
	"github.com/cadmium-im/cadmium-xmpp-server/xmpp/interface"

	// other
	"lab.pztrn.name/golibs/flagger"
	"lab.pztrn.name/golibs/mogrus"
)

type Context struct {
	Flagger *flagger.Flagger
	Log     *mogrus.LoggerHandler
	XMPP    xmppinterface.XMPPInterface
}

// Initialize application-wide Context.
func (c *Context) Initialize() {
	l := mogrus.New()
	l.Initialize()
	c.Log = l.CreateLogger("opensaps")
	c.Log.CreateOutput("stdout", os.Stdout, true)

	c.Flagger = flagger.New(c.Log)
	c.Flagger.Initialize()
}

// RegisterXMPPInterface registers XMPP interface.
func (c *Context) RegisterXMPPInterface(xi xmppinterface.XMPPInterface) {
	c.XMPP = xi
	c.XMPP.Initialize()
}
