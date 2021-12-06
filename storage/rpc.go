package storage

type PeerMessage struct {
	AppID    int64
	Uid      int64
	DeviceID int64
	Cmd      int32
	Raw      []byte
}

type PeerGroupMessage struct {
	AppID    int64
	Members  []int64
	DeviceID int64
	Cmd      int32
	Raw      []byte
}

type GroupMessage struct {
	AppID    int64
	GroupID  int64
	DeviceID int64
	Cmd      int32
	Raw      []byte
}

type SyncHistory struct {
	AppID     int64
	Uid       int64
	DeviceID  int64
	GroupID   int64
	LastMsgID int64
	Timestamp int32
}

type SyncGroupHistory struct {
	AppID     int64
	Uid       int64
	DeviceID  int64
	GroupID   int64
	LastMsgID int64
	Timestamp int32
}

type HistoryRequest struct {
	AppID int64
	Uid   int64
	Limit int32
}

type HistoryMessageID struct {
	MsgID    int64
	PreMsgId int64
}


type GroupHistoryMessageID struct {
	MessageIDs []*HistoryMessageID
}

type HistoryMessage struct {
	MsgID    int64
	DeviceID int64 //消息发送者所在的设备ID
	Cmd      int32
	Raw      []byte
}

type PeerHistoryMessage struct {
	Messages  []*HistoryMessage
	LastMsgID int64
	HasMore   bool
}

type GroupHistoryMessage PeerHistoryMessage

type LatestMessage struct {
	Messages []*HistoryMessage
}

type RPCStorage interface {
	SyncMessage(syncKey *SyncHistory, result *PeerHistoryMessage) error

	SyncGroupMessage(syncKey *SyncGroupHistory, result *GroupHistoryMessage) error

	SavePeerMessage(m *PeerMessage, result *HistoryMessageID) error

	SavePeerGroupMessage(m *PeerGroupMessage, result *GroupHistoryMessageID) error

	SaveGroupMessage(m *GroupMessage, result *HistoryMessageID) error

	GetNewCount(syncKey *SyncHistory, newCount *int64) error

	GetLatestMessage(r *HistoryRequest, l *LatestMessage) error

	Ping(int, *int) error
}
