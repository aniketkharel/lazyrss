package components

import (
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

func Feed_Item(gtx C, th, data M) D {
	return layout
}
