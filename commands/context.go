package commands

import (
	"context"

	"github.com/andersfylling/disgord"
)

type Ctx struct {
	Client      *disgord.Client
	Session     *disgord.Session
	Interaction *disgord.InteractionCreate
	Command     *Command
	Router      *Router
}

type ExecutionHandler func(ctx *Ctx)

func (ctx *Ctx) Response(r *ResponseOptions) error {
	return ctx.Interaction.Reply(context.Background(), *ctx.Session, CreateInteraction(r).AddEmbeds(r.Embeds...).SetContent(r.Context).AddMessageComponent(GetComponentsFromCommand(ctx.Command)...).Build())
}

func GetComponentsFromCommand(cmd *Command) []*disgord.MessageComponent {
	var ret []*disgord.MessageComponent
	for _, c := range cmd.Components {
		ret = append(ret, &disgord.MessageComponent{
			Type:        c.Type,
			Style:       c.Style,
			Label:       c.Label,
			CustomID:    c.CustomID,
			Url:         c.Url,
			Disabled:    c.Disabled,
			Options:     c.Options,
			Placeholder: c.Placeholder,
			MinValues:   c.MinValues,
			MaxValues:   c.MaxValues,
		})
	}
	return ret
}
