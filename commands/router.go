package commands

import (
	"log"

	"github.com/andersfylling/disgord"
)

type Router struct {
	CommandMap           map[string]*Command
	Commands             []*Command
	CommandComponents    []*CommandMessageComponent
	CommandComponentsMap map[string]*Command
	Client               *disgord.Client
}

func Create(client *disgord.Client) *Router {
	return &Router{
		Commands:             []*Command{},
		Client:               client,
		CommandComponentsMap: make(map[string]*Command),
	}
}

func (r *Router) Init() {
	r.Client.Gateway().Ready(r.InitializeCommands())
	r.Client.Gateway().InteractionCreate(r.Handler())
}

func (r *Router) Handler() disgord.HandlerInteractionCreate {
	return func(s disgord.Session, h *disgord.InteractionCreate) {
		ctx := &Ctx{
			Client:      r.Client,
			Session:     &s,
			Interaction: h,
			Router:      r,
		}
		for _, cmd := range r.Commands {
			if h.Data.Name != "" && h.Data.Name == cmd.Name {

				ctx.Command = cmd

				cmd.Handler(ctx)
				return
			}
		}

		for _, component := range r.CommandComponents {
			if h.Data.CustomID != "" && h.Data.CustomID == component.CustomID {
				component.Handler(&ComponentCtx{
					Client:      r.Client,
					Session:     &s,
					Interaction: h,
					Router:      r,
				})
				return
			}
		}
	}
}

func (r *Router) RegisterCMDList(commands ...*Command) {
	r.Commands = append(r.Commands, commands...)
	for _, c := range commands {
		r.CommandComponents = append(r.CommandComponents, c.Components...)
	}
}

func (r *Router) GetCommand(name string) *Command {
	return r.CommandMap[name]
}

func (r *Router) RegisterCommand(command *Command) {
	r.Commands = append(r.Commands, command)
	r.CommandComponents = append(r.CommandComponents, command.Components...)
}

func (r *Router) InitializeCommands() disgord.HandlerReady {
	return func(s disgord.Session, h *disgord.Ready) {
		user, _ := r.Client.Cache().GetCurrentUser()
		for i := range r.Commands {
			if err := r.Client.ApplicationCommand(user.ID).Global().Create(&disgord.CreateApplicationCommand{
				Type:        disgord.ApplicationCommandType(r.Commands[i].Type),
				Name:        r.Commands[i].Name,
				Description: r.Commands[i].Description,
				Options:     r.Commands[i].Options,
			}); err != nil {
				log.Fatal(err)
			}
		}
	}
}
