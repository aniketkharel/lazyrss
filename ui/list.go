package ui

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

func Construct_List(th *material.Theme, gtx layout.Context) layout.Dimensions {
	childrens := []layout.Widget{
		func(gtx layout.Context) layout.Dimensions {
			l := material.H6(th, "hello there")
			return l.Layout(gtx)
		},
		func(gtx layout.Context) layout.Dimensions {
			l := material.H6(th, "hello there")
			return l.Layout(gtx)
		},
		func(gtx layout.Context) layout.Dimensions {
			l := material.H6(th, "hello there")
			return l.Layout(gtx)
		},
	}
	list := &widget.List{
		List: layout.List{
			Axis: layout.Vertical,
		},
	}
	return material.List(th, list).Layout(gtx, len(childrens), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, childrens[i])
	})
}
