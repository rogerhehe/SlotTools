package msgproto

import (
	"slotserver/common/msgid"
	"slotserver/common/pb"

	"google.golang.org/protobuf/reflect/protoreflect"
)

var srv = map[uint16]protoreflect.MessageType{
	msgid.C2S_HANDSHAKE:         (&pb.S2CHandShake).ProtoReflect().Type(),
	msgid.S2C_HANDSHAKE:         (&pb.C2SHandShake).ProtoReflect().Type(),
	msgid.C2S_ECHO:              (&pb.S2CEcho).ProtoReflect().Type(),
	msgid.C2S_SYNC_ERROR:        (&pb.C2SSyncError).ProtoReflect().Type(),
	msgid.C2S_SYNC_GAME_SERVER:  (&pb.C2SSyncGameServer).ProtoReflect().Type(),
	msgid.C2S_SYNC_ROOM_SERVER:  (&pb.C2SSyncRoomServer).ProtoReflect().Type(),
	msgid.C2S_SYNC_WORLD_SERVER: (&pb.C2SSyncWorldServer).ProtoReflect().Type(),
	msgid.C2S_SYNC_FORWARD_MSG:  (&pb.C2SForwardMsg).ProtoReflect().Type(),
}
