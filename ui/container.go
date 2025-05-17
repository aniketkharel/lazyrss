package ui

import (
	"gioui.org/widget/material"
	"github.com/aniketkharel/rssreader/models"
)

type Wrapper struct {
	data *[]models.FeedXML
}

func UiWrapper(data *[]models.FeedXML) *Wrapper {
	return &Wrapper{
		data: data,
	}
}

func (wrapper *Wrapper) Layout(th *material.Theme, gtx C) D {
	list := NewsFeedList()
	return list.Layout(th, gtx, wrapper.data)
}
