package screen

import (
	"eklase/state"
	"image/color"
	"image/png"
	"log"
	"os"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// MainMenu defines a main menu screen layout.
func MainMenu(th *material.Theme, state *state.State) Screen {
	var (
		add   widget.Clickable
		list  widget.Clickable
		quit  widget.Clickable
		image widget.Image
	)
	image.Src = draw()
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		// widgetcolor := th.ContrastBg                                 To change the widget's background color
		// widgetcolor.A = 0x45
		// max := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y)
		// paint.FillShape(gtx.Ops, widgetcolor, clip.Rect{Max: max}.Op())

		matAddBut := material.Button(th, &add, "Add student")
		matAddBut.Font = text.Font{Variant: "Mono", Weight: text.Bold, Style: text.Italic}
		matAddBut.Background = color.NRGBA{A: 0xff, R: 0x1e, G: 0x4d, B: 0x24}
		matListBut := material.Button(th, &list, "List students")
		matListBut.Font = text.Font{Variant: "Mono", Weight: text.Bold, Style: text.Italic}
		matQuitBut := material.Button(th, &quit, "Quit")
		matQuitBut.Font = text.Font{Variant: "Smallcaps", Style: text.Italic}
		matQuitBut.Background = color.NRGBA{A: 0xff, R: 0xc6, G: 0x28, B: 0x28}

		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(matAddBut.Layout)),
			layout.Rigid(rowInset(matListBut.Layout)),
			layout.Rigid(rowInset(matQuitBut.Layout)),
		)
		layout.SE.Layout(gtx, image.Layout)
		if add.Clicked() {
			return AddStudent(th, state), d
		}
		if list.Clicked() {
			return ListTable(th, state), d
		}
		if quit.Clicked() {
			state.Quit()
		}
		return nil, d
	}
}

func draw() paint.ImageOp {
	f, err := os.Open("sign-info-icon.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	src := paint.NewImageOp(img)
	return src
}
