require "Assets/_LuaScripts/Game/Net/Proto/Importer"

local pb_scheme = {
	[MsgID.C2S_HANDSHAKE] = "pb.S2CHandShake",
	[MsgID.S2C_HANDSHAKE] = "pb.S2CHandShake",
	[MsgID.C2S_ECHO] = "pb.S2CEcho",
	[MsgID.C2S_SYNC_ERROR] = "pb.C2SSyncError",
	[MsgID.C2S_SYNC_GAME_SERVER] = "pb.C2SSyncGameServer",
	[MsgID.C2S_SYNC_ROOM_SERVER] = "pb.C2SSyncRoomServer",
	[MsgID.C2S_SYNC_WORLD_SERVER] = "pb.C2SSyncWorldServer",
	[MsgID.C2S_SYNC_FORWARD_MSG] = "pb.C2SForwardMsg",
}

return pb_scheme