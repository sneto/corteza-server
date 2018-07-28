package websocket

import (
	"context"
	"github.com/crusttech/crust/auth"
	"github.com/crusttech/crust/sam/service"
	"github.com/crusttech/crust/sam/types"
	"github.com/crusttech/crust/sam/websocket/incoming"
	"github.com/crusttech/crust/sam/websocket/outgoing"
)

func (s *Session) channelJoin(ctx context.Context, p incoming.ChannelJoin) error {
	// @todo: check access / can we join this channel (should be done by service layer)

	s.subs.Add(p.ChannelID)

	// Telling all subscribers of the channel we're joining that we are joining.
	var chJoin = &outgoing.ChannelJoin{
		ID:     p.ChannelID,
		UserID: uint64toa(auth.GetIdentityFromContext(ctx).GetID()),
	}

	// Send to all channel subscribers
	s.broadcast(chJoin, &p.ChannelID)

	return nil
}

func (s *Session) channelPart(ctx context.Context, p incoming.ChannelPart) error {
	// @todo: check access / can we part this channel? (should be done by service layer)

	// First, let's unsubscribe, so we don't hear echos
	if p.ChannelID != nil {
		s.subs.Delete(*p.ChannelID)
	} else {
		s.subs.DeleteAll()
	}

	// This payload will tell everyone that we're parting from ALL channels
	var chPart = &outgoing.ChannelPart{
		ID:     p.ChannelID,
		UserID: uint64toa(auth.GetIdentityFromContext(ctx).GetID()),
	}

	s.broadcast(chPart, p.ChannelID)

	return nil
}

func (s *Session) channelList(ctx context.Context, p incoming.ChannelList) error {
	channels, err := service.Channel().Find(ctx, nil)
	if err != nil {
		return err
	}

	return s.respond(payloadFromChannels(channels))
}

func (s *Session) channelOpen(ctx context.Context, p incoming.ChannelOpen) error {
	var (
		filter = &types.MessageFilter{
			ChannelID:     parseUInt64(p.ChannelID),
			FromMessageID: parseUInt64(p.Since),
		}
	)

	messages, err := service.Message().Find(ctx, filter)
	if err != nil {
		return err
	}

	return s.respond(payloadFromMessages(messages))
}
