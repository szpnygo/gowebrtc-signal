// Code generated by gogs. DO NOT EDIT.
package model

import (
	"context"
	"reflect"

	"github.com/metagogs/gogs"
	"github.com/metagogs/gogs/component"
	"github.com/metagogs/gogs/packet"
	"github.com/metagogs/gogs/session"
)

func RegisterAllComponents(s *gogs.App, srv Component) {
	registerBaseWorldComponent(s, srv)

}

func registerBaseWorldComponent(s *gogs.App, srv Component) {
	s.RegisterComponent(_BaseWorldComponentDesc, srv)
}

type Component interface {
	BindUser(ctx context.Context, s *session.Session, in *BindUser)

	OutgoingMessage(ctx context.Context, s *session.Session, in *OutgoingMessage)
}

func _BaseWorldComponent_BindUser_Handler(srv interface{}, ctx context.Context, sess *session.Session, in interface{}) {
	srv.(Component).BindUser(ctx, sess, in.(*BindUser))
}

func _BaseWorldComponent_OutgoingMessage_Handler(srv interface{}, ctx context.Context, sess *session.Session, in interface{}) {
	srv.(Component).OutgoingMessage(ctx, sess, in.(*OutgoingMessage))
}

var _BaseWorldComponentDesc = component.ComponentDesc{
	ComponentName:  "BaseWorldComponent",
	ComponentIndex: 1, // equal to module index
	ComponentType:  (*Component)(nil),
	Methods: []component.ComponentMethodDesc{
		{
			MethodIndex: packet.CreateAction(packet.ServicePacket, 1, 1), // 0x810001 8454145
			FieldType:   reflect.TypeOf(BindUser{}),
			Handler:     _BaseWorldComponent_BindUser_Handler,
			FieldHandler: func() interface{} {
				return new(BindUser)
			},
		},
		{
			MethodIndex: packet.CreateAction(packet.ServicePacket, 1, 2), // 0x810002 8454146
			FieldType:   reflect.TypeOf(BindSuccess{}),
			Handler:     nil,
			FieldHandler: func() interface{} {
				return new(BindSuccess)
			},
		},
		{
			MethodIndex: packet.CreateAction(packet.ServicePacket, 1, 3), // 0x810003 8454147
			FieldType:   reflect.TypeOf(BindFail{}),
			Handler:     nil,
			FieldHandler: func() interface{} {
				return new(BindFail)
			},
		},
		{
			MethodIndex: packet.CreateAction(packet.ServicePacket, 1, 4), // 0x810004 8454148
			FieldType:   reflect.TypeOf(OutgoingMessage{}),
			Handler:     _BaseWorldComponent_OutgoingMessage_Handler,
			FieldHandler: func() interface{} {
				return new(OutgoingMessage)
			},
		},
		{
			MethodIndex: packet.CreateAction(packet.ServicePacket, 1, 5), // 0x810005 8454149
			FieldType:   reflect.TypeOf(IncomingMessage{}),
			Handler:     nil,
			FieldHandler: func() interface{} {
				return new(IncomingMessage)
			},
		},
		{
			MethodIndex: packet.CreateAction(packet.ServicePacket, 1, 6), // 0x810006 8454150
			FieldType:   reflect.TypeOf(OutgoingFail{}),
			Handler:     nil,
			FieldHandler: func() interface{} {
				return new(OutgoingFail)
			},
		},
	},
}
