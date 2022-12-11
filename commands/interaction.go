package commands

import "github.com/andersfylling/disgord"

type InteractionResponse struct {
	*disgord.CreateInteractionResponse
}

type ResponseOptions struct {
	Context string
	Embeds  []*disgord.Embed
	Flags   int
}

func CreateInteraction(r *ResponseOptions) *InteractionResponse {
	return &InteractionResponse{
		CreateInteractionResponse: &disgord.CreateInteractionResponse{
			Type: 4,
			Data: &disgord.CreateInteractionResponseData{
				Flags: disgord.MessageFlag(r.Flags),
			},
		},
	}
}

func (ctx *InteractionResponse) SetContent(text string) *InteractionResponse {
	ctx.Data.Content = text
	return ctx
}

func (ctx *InteractionResponse) AddEmbeds(embed ...*disgord.Embed) *InteractionResponse {
	ctx.Data.Embeds = append(ctx.CreateInteractionResponse.Data.Embeds, embed...)
	return ctx
}

func (ctx *InteractionResponse) AddMessageComponent(com ...*disgord.MessageComponent) *InteractionResponse {
	ctx.Data.Components = append(ctx.CreateInteractionResponse.Data.Components, com...)
	return ctx
}

func (ctx *InteractionResponse) Build() *disgord.CreateInteractionResponse {
	return ctx.CreateInteractionResponse
}
