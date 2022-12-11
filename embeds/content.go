package embeds

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andersfylling/disgord"
)

type EmbedBuilder struct {
	*disgord.Embed
}

func CreateEmbed() *EmbedBuilder {
	return &EmbedBuilder{
		Embed: &disgord.Embed{},
	}
}

func (e *EmbedBuilder) TextLimiter(str string, limiter int) string {
	if len(str) > limiter {
		str = str[:limiter]
	}
	return str
}

func (e *EmbedBuilder) SetAuthor(name, icon string) *EmbedBuilder {
	e.Author = &disgord.EmbedAuthor{
		Name:    name,
		IconURL: icon,
	}
	return e
}

func (e *EmbedBuilder) AddField(name, value string, inline bool) *EmbedBuilder {
	e.Fields = append(e.Fields, &disgord.EmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})
	return e
}

func (e *EmbedBuilder) SetTitle(title string) *EmbedBuilder {
	e.Title = e.TextLimiter(title, 256)
	return e
}

func (e *EmbedBuilder) SetDescription(description string) *EmbedBuilder {
	e.Description = e.TextLimiter(description, 2048)
	return e
}

func (e *EmbedBuilder) SetThumbnail(url string) *EmbedBuilder {
	e.Thumbnail = &disgord.EmbedThumbnail{
		URL: url,
	}
	return e
}

func (e *EmbedBuilder) SetColor(color string) *EmbedBuilder {
	color = strings.Replace(color, "0x", "", -1)
	color = strings.Replace(color, "0X", "", -1)
	color = strings.Replace(color, "#", "", -1)
	colorInt, err := strconv.ParseInt(color, 16, 64)
	if err != nil {
		panic(err)
	}
	e.Color = int(colorInt)
	return e
}

func (e *EmbedBuilder) SetImage(url string) *EmbedBuilder {
	e.Image = &disgord.EmbedImage{
		URL: url,
	}
	return e
}

func (e *EmbedBuilder) ConvertFieldsToDescription() *EmbedBuilder {
	if e.Description == "" || len(e.Fields) > 0 {
		for _, field := range e.Fields {
			if e.Description != "" {
				e.Description += "\n"
			}
			e.Description += fmt.Sprintf("**%s**: %s", field.Name, field.Value)
		}
	}
	return e
}

func (e *EmbedBuilder) Build() *disgord.Embed {
	return e.Embed
}
