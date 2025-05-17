package ui

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/aniketkharel/rssreader/models"
)

type (
	D = layout.Dimensions
	C = layout.Context
	W = layout.Widget
)

type FeedList struct {
	List *widget.List
}

func NewsFeedList() *FeedList {
	return &FeedList{
		List: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}

}

func (f *FeedList) Layout(th *material.Theme, gtx C,
	data *[]models.FeedXML) D {
	var childs []W
	for _, v := range *data {
		for _, item := range v.Channel.Items {
			title := item.Title
			link := item.Link

			childs = append(childs, func(gtx C) D {
				return layout.Flex{
					Axis: layout.Vertical,
				}.
					Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return layout.UniformInset(unit.Dp(8)).Layout(gtx,
								material.Body1(th, title).Layout,
							)
						}),
						layout.Rigid(func(gtx C) D {
							return layout.UniformInset(unit.Dp(4)).Layout(gtx,
								material.Body2(th, link).Layout,
							)
						}),
						layout.Rigid(func(gtx C) D {
							return layout.UniformInset(unit.Dp(4)).Layout(gtx,
								material.Subtitle2(th, item.PubDate).Layout,
							)
						}),
					)
			})
		}
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx C) D {
			return material.List(th, f.List).Layout(gtx, len(childs), func(gtx C, i int) D {
				return layout.UniformInset(unit.Dp(12)).Layout(gtx, childs[i])
			})
		}),
	)
}
