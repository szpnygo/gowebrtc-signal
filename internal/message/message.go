// Code generated by gogs. DO NOT EDIT.
package message

import (
	"github.com/metagogs/gogs/session"
	"github.com/szpnygo/gowebrtc-signal/model"
)

func SendBindUser(s *session.Session, in *model.BindUser) error {
	return s.SendMessage(in, "BindUser")
}

func SendBindSuccess(s *session.Session, in *model.BindSuccess) error {
	return s.SendMessage(in, "BindSuccess")
}

func SendBindFail(s *session.Session, in *model.BindFail) error {
	return s.SendMessage(in, "BindFail")
}

func SendOutgoingMessage(s *session.Session, in *model.OutgoingMessage) error {
	return s.SendMessage(in, "OutgoingMessage")
}

func SendIncomingMessage(s *session.Session, in *model.IncomingMessage) error {
	return s.SendMessage(in, "IncomingMessage")
}

func SendOutgoingFail(s *session.Session, in *model.OutgoingFail) error {
	return s.SendMessage(in, "OutgoingFail")
}
