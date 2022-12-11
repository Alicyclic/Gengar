package commands

import (
	"context"

	"github.com/andersfylling/disgord"
)

type ComponentCtx struct {
	Client      *disgord.Client
	Session     *disgord.Session
	Interaction *disgord.InteractionCreate
	Router      *Router
}

type ExecutionHandlerComponent func(ctx *ComponentCtx)

func (ctx *ComponentCtx) Response(r *ResponseOptions) error {
	return ctx.Interaction.Reply(context.Background(), *ctx.Session, CreateInteraction(r).SetContent(r.Context).AddEmbeds(r.Embeds...).Build())
}
