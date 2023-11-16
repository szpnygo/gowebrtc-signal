package baseworld

import (
	"context"

	"github.com/metagogs/gogs/gslog"
	"github.com/metagogs/gogs/session"
	"github.com/szpnygo/gowebrtc-signal/internal/svc"
	"github.com/szpnygo/gowebrtc-signal/model"
	"go.uber.org/zap"
)

type OutgoingMessageLogic struct {
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	session *session.Session
	*zap.Logger
}

func NewOutgoingMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext, sess *session.Session) *OutgoingMessageLogic {
	return &OutgoingMessageLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		session: sess,
		Logger:  gslog.NewLog("outgoing_message_logic"),
	}
}

func (l *OutgoingMessageLogic) Handler(in *model.OutgoingMessage) {

}
