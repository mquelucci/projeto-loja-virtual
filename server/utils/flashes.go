package utils

import "github.com/gin-contrib/sessions"

type SessionFlashes struct {
	MsgSucesso string
	MsgFalha   string
	MsgInfo    string
}

func (flashes *SessionFlashes) GetFlashes(session sessions.Session) {
	flashes.MsgSucesso = session.Flashes("MsgSucesso")[0].(string)
	flashes.MsgFalha = session.Flashes("MsgFalha")[0].(string)
	flashes.MsgInfo = session.Flashes("MsgInfo")[0].(string)
	session.Flashes()
	session.Save()
}
