package components

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/aniketkharel/rssreader/models"
)

type (
	D  = layout.Dimensions
	C  = layout.Context
	M  = models.FeedXML
	th = *material.Theme
)

func Feed_Item(gtx C, th th, data M) D {
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			fmt.Printf("Expand: %v\n", gtx.Constraints)

			return layoutWidget(gtx, th, data.Channel.Title)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layoutWidget(gtx, th, data.Channel.Description)
		}),
	)
}

func layoutWidget(gtx C, th th, s string) layout.Dimensions {
	title := material.H6(th, s)
	return title.Layout(gtx)
}
