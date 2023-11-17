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

type BindUserLogic struct {
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	session *session.Session
	*zap.Logger
}

func NewBindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext, sess *session.Session) *BindUserLogic {
	return &BindUserLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		session: sess,
		Logger:  gslog.NewLog("bind_user_logic"),
	}
}

func (l *BindUserLogic) Handler(in *model.BindUser) {
	// check uid
	if len(in.Name) == 0 {
		l.Logger.Info("uid is empty")
		_ = message.SendBindSuccess(l.session, &model.BindSuccess{})
		return
	}

	// check uid should be update
	if l.session.UID() == in.Name {
		l.Logger.Info("uid is already bind", zap.String("uid", in.Name))
		_ = message.SendBindSuccess(l.session, &model.BindSuccess{})
		return
	}

	result := gogs.GetSessionByUID(in.Name, nil)
	if len(result) > 0 {
		// uid is already online
		l.Logger.Info("uid is already online", zap.String("uid", in.Name))
		_ = message.SendBindFail(l.session, &model.BindFail{
			Reason: "uid is already online",
		})
		return
	}

	// bind uid
	l.session.SetUID(in.Name)
	_ = message.SendBindSuccess(l.session, &model.BindSuccess{})
	l.Logger.Info("bind uid success", zap.String("uid", l.session.UID()))
}
