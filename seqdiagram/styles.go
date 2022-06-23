package seqdiagram

import (
	"github.com/lmika/goseq/seqdiagram/graphbox"
)

// Diagram styles
type DiagramStyles struct {
	// Diagram margins
	Margin graphbox.Point

	// Styling of the actor box
	ActorBox     graphbox.ActorBoxStyle
	ActorIconBox graphbox.ActorIconBoxStyle

	// Styling of the note box
	NoteBox graphbox.NoteBoxStyle

	MultiNoteOverlap int

	// Styling of the activity line
	ActivityLine graphbox.ActivityLineStyle

	// Styling of arrow heads
	ArrowHeads map[ArrowHead]*graphbox.ArrowHeadStyle

	// Styling of the diagram title
	Title graphbox.TitleStyle

	// Block styling
	Block graphbox.BlockStyle

	// Styles of dividers
	Divider map[DividerType]graphbox.DividerStyle
}

// Fonts
var standardFont = mustLoadFont()

// The Default style
var DefaultStyle = &DiagramStyles{
	Margin: graphbox.Point{X: 8, Y: 8},
	ActorBox: graphbox.ActorBoxStyle{
		Font:     standardFont,
		FontSize: 16,
		Padding:  graphbox.Point{X: 16, Y: 8},
		Margin:   graphbox.Point{X: 8, Y: 8},
	},
	ActorIconBox: graphbox.ActorIconBoxStyle{
		Font:     standardFont,
		FontSize: 16,
		Padding:  graphbox.Point{X: 16, Y: 8},
		Margin:   graphbox.Point{X: 8, Y: 8},
		IconGap:  4,
	},
	NoteBox: graphbox.NoteBoxStyle{
		Font:     standardFont,
		FontSize: 14,
		Padding:  graphbox.Point{X: 8, Y: 4},
		Margin:   graphbox.Point{X: 8, Y: 8},
	},
	MultiNoteOverlap: 16,
	ActivityLine: graphbox.ActivityLineStyle{
		Font:          standardFont,
		FontSize:      14,
		SelfRefWidth:  48,
		SelfRefHeight: 24,
		Margin:        graphbox.Point{X: 16, Y: 8},
		TextGap:       4,
	},
	ArrowHeads: map[ArrowHead]*graphbox.ArrowHeadStyle{
		SolidArrowHead: {
			Xs:        []int{-9, 0, -9},
			Ys:        []int{-5, 0, 5},
			BaseStyle: "stroke:black;fill:black;stroke-width:2px;",
		},
		OpenArrowHead: {
			Xs:        []int{-9, 0, -9},
			Ys:        []int{-5, 0, 5},
			BaseStyle: "stroke:black;fill:none;stroke-width:2px;",
		},
		BarbArrowHead: {
			Xs:        []int{-11, 0},
			Ys:        []int{-7, 0},
			BaseStyle: "stroke:black;fill:black;stroke-width:2px;",
		},
		LowerBarbArrowHead: {
			Xs:        []int{-11, 0},
			Ys:        []int{7, 0},
			BaseStyle: "stroke:black;fill:black;stroke-width:2px;",
		},
	},
	Title: graphbox.TitleStyle{
		Font:     standardFont,
		FontSize: 20,
		Padding:  graphbox.Point{X: 4, Y: 16},
	},
	Block: graphbox.BlockStyle{
		Margin:           graphbox.Point{X: 8, Y: 8},
		TextPadding:      graphbox.Point{X: 4, Y: 4},
		MessagePadding:   graphbox.Point{X: 4, Y: 4},
		GapWidth:         4,
		PrefixExtraWidth: 4,

		Font:      standardFont,
		FontSize:  14,
		MidMargin: 4,
	},
	Divider: map[DividerType]graphbox.DividerStyle{
		DTGap: {
			Font:        standardFont,
			FontSize:    14,
			Padding:     graphbox.Point{X: 16, Y: 8},
			Margin:      graphbox.Point{X: 8, Y: 8},
			TextPadding: graphbox.Point{X: 0, Y: 0},
			Shape:       graphbox.DSFullRect,
		},
		DTFrame: {
			Font:        standardFont,
			FontSize:    14,
			Padding:     graphbox.Point{X: 16, Y: 8},
			Margin:      graphbox.Point{X: 8, Y: 8},
			TextPadding: graphbox.Point{X: 0, Y: 0},
			Shape:       graphbox.DSFramedRect,
		},
		DTLine: {
			Font:        standardFont,
			FontSize:    14,
			Padding:     graphbox.Point{X: 16, Y: 4},
			Margin:      graphbox.Point{X: 8, Y: 16},
			TextPadding: graphbox.Point{X: 4, Y: 2},
			Shape:       graphbox.DSFullLine,
		},
		DTSpacer: {
			Font:        standardFont,
			FontSize:    14,
			Padding:     graphbox.Point{X: 16, Y: 4},
			Margin:      graphbox.Point{X: 8, Y: 16},
			TextPadding: graphbox.Point{X: 0, Y: 0},
			Shape:       graphbox.DSSpacerRect,
		},
	},
}

// The Tight style.  Same horizontal dimensions as the normal
// style but slightly smaller vertical margins
var TightStyle = &DiagramStyles{
	Margin: graphbox.Point{X: 8, Y: 8},
	ActorBox: graphbox.ActorBoxStyle{
		Font:     standardFont,
		FontSize: 16,
		Padding:  graphbox.Point{X: 16, Y: 4},
		Margin:   graphbox.Point{X: 8, Y: 4},
	},
	ActorIconBox: graphbox.ActorIconBoxStyle{
		Font:     standardFont,
		FontSize: 16,
		Padding:  graphbox.Point{X: 16, Y: 8},
		Margin:   graphbox.Point{X: 8, Y: 4},
		IconGap:  4,
	},
	NoteBox: graphbox.NoteBoxStyle{
		Font:     standardFont,
		FontSize: 14,
		Padding:  graphbox.Point{X: 8, Y: 4},
		Margin:   graphbox.Point{X: 8, Y: 4},
	},
	MultiNoteOverlap: 16,
	ActivityLine: graphbox.ActivityLineStyle{
		Font:          standardFont,
		FontSize:      14,
		SelfRefWidth:  48,
		SelfRefHeight: 12,
		Margin:        graphbox.Point{X: 16, Y: 4},
		TextGap:       4,
	},
	ArrowHeads: map[ArrowHead]*graphbox.ArrowHeadStyle{
		SolidArrowHead: {
			Xs:        []int{-9, 0, -9},
			Ys:        []int{-5, 0, 5},
			BaseStyle: "stroke:black;fill:black;stroke-width:2px;",
		},
		OpenArrowHead: {
			Xs:        []int{-9, 0, -9},
			Ys:        []int{-5, 0, 5},
			BaseStyle: "stroke:black;fill:none;stroke-width:2px;",
		},
		BarbArrowHead: {
			Xs:        []int{-11, 0},
			Ys:        []int{-7, 0},
			BaseStyle: "stroke:black;fill:black;stroke-width:2px;",
		},
		LowerBarbArrowHead: {
			Xs:        []int{-11, 0},
			Ys:        []int{7, 0},
			BaseStyle: "stroke:black;fill:black;stroke-width:2px;",
		},
	},
	Title: graphbox.TitleStyle{
		Font:     standardFont,
		FontSize: 20,
		Padding:  graphbox.Point{X: 4, Y: 8},
	},
	Block: graphbox.BlockStyle{
		Margin:           graphbox.Point{X: 8, Y: 8},
		TextPadding:      graphbox.Point{X: 4, Y: 4},
		MessagePadding:   graphbox.Point{X: 4, Y: 4},
		GapWidth:         4,
		PrefixExtraWidth: 4,

		Font:      standardFont,
		FontSize:  14,
		MidMargin: 4,
	},
	Divider: map[DividerType]graphbox.DividerStyle{
		DTGap: {
			Font:        standardFont,
			FontSize:    14,
			Padding:     graphbox.Point{X: 16, Y: 8},
			Margin:      graphbox.Point{X: 8, Y: 8},
			TextPadding: graphbox.Point{X: 0, Y: 0},
			Shape:       graphbox.DSFullRect,
		},
		DTFrame: {
			Font:        standardFont,
			FontSize:    14,
			Padding:     graphbox.Point{X: 16, Y: 8},
			Margin:      graphbox.Point{X: 8, Y: 8},
			TextPadding: graphbox.Point{X: 0, Y: 0},
			Shape:       graphbox.DSFramedRect,
		},
		DTLine: {
			Font:        standardFont,
			FontSize:    14,
			Padding:     graphbox.Point{X: 16, Y: 4},
			Margin:      graphbox.Point{X: 8, Y: 16},
			TextPadding: graphbox.Point{X: 4, Y: 2},
			Shape:       graphbox.DSFullLine,
		},
		DTSpacer: {
			Font:        standardFont,
			FontSize:    14,
			Padding:     graphbox.Point{X: 16, Y: 4},
			Margin:      graphbox.Point{X: 8, Y: 16},
			TextPadding: graphbox.Point{X: 0, Y: 0},
			Shape:       graphbox.DSSpacerRect,
		},
	},
}

// The small style.  This has narrower margins and font sizes and
// is used to produce smaller diagrams.
var SmallStyle = &DiagramStyles{
	Margin: graphbox.Point{X: 4, Y: 4},
	ActorBox: graphbox.ActorBoxStyle{
		Font:     standardFont,
		FontSize: 14,
		Padding:  graphbox.Point{X: 12, Y: 6},
		Margin:   graphbox.Point{X: 8, Y: 8},
	},
	ActorIconBox: graphbox.ActorIconBoxStyle{
		Font:     standardFont,
		FontSize: 14,
		Padding:  graphbox.Point{X: 12, Y: 6},
		Margin:   graphbox.Point{X: 8, Y: 8},
		IconGap:  2,
	},
	NoteBox: graphbox.NoteBoxStyle{
		Font:     standardFont,
		FontSize: 12,
		Padding:  graphbox.Point{X: 6, Y: 3},
		Margin:   graphbox.Point{X: 6, Y: 6},
	},
	MultiNoteOverlap: 8,
	ActivityLine: graphbox.ActivityLineStyle{
		Font:          standardFont,
		FontSize:      12,
		Margin:        graphbox.Point{X: 8, Y: 8},
		TextGap:       4,
		SelfRefWidth:  32,
		SelfRefHeight: 12,
	},
	ArrowHeads: map[ArrowHead]*graphbox.ArrowHeadStyle{
		SolidArrowHead: {
			Xs:        []int{-7, 0, -7},
			Ys:        []int{-4, 0, 4},
			BaseStyle: "stroke:black;fill:black;stroke-width:2px;",
		},
		OpenArrowHead: {
			Xs:        []int{-7, 0, -7},
			Ys:        []int{-4, 0, 4},
			BaseStyle: "stroke:black;fill:none;stroke-width:2px;",
		},
		BarbArrowHead: {
			Xs:        []int{-9, 0},
			Ys:        []int{-5, 0},
			BaseStyle: "stroke:black;fill:black;stroke-width:2px;",
		},
		LowerBarbArrowHead: {
			Xs:        []int{-9, 0},
			Ys:        []int{5, 0},
			BaseStyle: "stroke:black;fill:black;stroke-width:2px;",
		},
	},
	Title: graphbox.TitleStyle{
		Font:     standardFont,
		FontSize: 18,
		Padding:  graphbox.Point{X: 2, Y: 8},
	},
	Block: graphbox.BlockStyle{
		Margin:           graphbox.Point{X: 5, Y: 5},
		TextPadding:      graphbox.Point{X: 3, Y: 2},
		MessagePadding:   graphbox.Point{X: 3, Y: 2},
		GapWidth:         3,
		PrefixExtraWidth: 3,

		Font:      standardFont,
		FontSize:  12,
		MidMargin: 2,
	},
	Divider: map[DividerType]graphbox.DividerStyle{
		DTGap: {
			Font:        standardFont,
			FontSize:    12,
			Padding:     graphbox.Point{X: 12, Y: 6},
			Margin:      graphbox.Point{X: 6, Y: 6},
			TextPadding: graphbox.Point{X: 0, Y: 0},
			Shape:       graphbox.DSFullRect,
		},
		DTFrame: {
			Font:        standardFont,
			FontSize:    12,
			Padding:     graphbox.Point{X: 12, Y: 6},
			Margin:      graphbox.Point{X: 6, Y: 6},
			TextPadding: graphbox.Point{X: 0, Y: 0},
			Shape:       graphbox.DSFramedRect,
		},
		DTLine: {
			Font:        standardFont,
			FontSize:    12,
			Padding:     graphbox.Point{X: 12, Y: 6},
			Margin:      graphbox.Point{X: 6, Y: 12},
			TextPadding: graphbox.Point{X: 2, Y: 1},
			Shape:       graphbox.DSFullLine,
		},
		DTSpacer: {
			Font:        standardFont,
			FontSize:    12,
			Padding:     graphbox.Point{X: 12, Y: 6},
			Margin:      graphbox.Point{X: 6, Y: 12},
			TextPadding: graphbox.Point{X: 2, Y: 1},
			Shape:       graphbox.DSSpacerRect,
		},
	},
}

func StyleByName(name string) *DiagramStyles {
	switch name {
	case "tight":
		return TightStyle
	case "small":
		return SmallStyle
	case "default":
		return DefaultStyle
	default:
		return DefaultStyle
	}
}
