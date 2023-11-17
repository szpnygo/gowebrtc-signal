package baseworld

import (
	"context"

	"github.com/metagogs/gogs"
	"github.com/metagogs/gogs/gslog"
	"github.com/metagogs/gogs/session"
	"github.com/szpnygo/gowebrtc-signal/internal/message"
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
	// check receiver and content is not empty
	if len(in.Recipient) == 0 || len(in.Content) == 0 {
		l.Logger.Info("receiver or content is empty")
		return
	}

	// check receiver is online
	receiver := gogs.GetSessionByUID(in.Recipient, nil)
	if len(receiver) == 0 {
		l.Logger.Info("receiver is offline", zap.String("uid", in.Recipient))
		_ = message.SendOutgoingFail(l.session, &model.OutgoingFail{
			Reason: "receiver is offline",
		})
		return
	}

	// send message
	_ = message.SendIncomingMessage(receiver[0], &model.IncomingMessage{
		Sender:  l.session.UID(),
		Content: in.Content,
	})

}
