// Auto-generated by avdl-compiler v1.3.9 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/gregor.avdl

package chat1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type GenericPayload struct {
	Action string `codec:"Action" json:"Action"`
}

type NewConversationPayload struct {
	Action    string         `codec:"Action" json:"Action"`
	ConvID    ConversationID `codec:"convID" json:"convID"`
	InboxVers InboxVers      `codec:"inboxVers" json:"inboxVers"`
}

type NewMessagePayload struct {
	Action    string         `codec:"Action" json:"Action"`
	ConvID    ConversationID `codec:"convID" json:"convID"`
	Message   MessageBoxed   `codec:"message" json:"message"`
	InboxVers InboxVers      `codec:"inboxVers" json:"inboxVers"`
}

type ReadMessagePayload struct {
	Action    string         `codec:"Action" json:"Action"`
	ConvID    ConversationID `codec:"convID" json:"convID"`
	MsgID     MessageID      `codec:"msgID" json:"msgID"`
	InboxVers InboxVers      `codec:"inboxVers" json:"inboxVers"`
}

type SetStatusPayload struct {
	Action    string             `codec:"Action" json:"Action"`
	ConvID    ConversationID     `codec:"convID" json:"convID"`
	Status    ConversationStatus `codec:"status" json:"status"`
	InboxVers InboxVers          `codec:"inboxVers" json:"inboxVers"`
}

type GregorInterface interface {
}

func GregorProtocol(i GregorInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "chat.1.gregor",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type GregorClient struct {
	Cli rpc.GenericClient
}
