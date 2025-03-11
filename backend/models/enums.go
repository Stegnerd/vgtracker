package models

type PresetConfig string

const (
	Aura PresetConfig = "Aura"
	Lara PresetConfig = "Lara"
	Nora PresetConfig = "Nora"
)

var AllPresetConfigs = []struct {
	Value  PresetConfig
	TSName string
}{
	{Aura, "Aura"},
	{Lara, "Lara"},
	{Nora, "Nora"},
}

type PaletteColor string

const (
	Noir    PaletteColor = "noir"
	Emerald PaletteColor = "emerald"
	Green   PaletteColor = "green"
	Lime    PaletteColor = "lime"
	Orange  PaletteColor = "orange"
	Amber   PaletteColor = "amber"
	Yellow  PaletteColor = "yellow"
	Teal    PaletteColor = "teal"
	Cyan    PaletteColor = "cyan"
	Sky     PaletteColor = "sky"
	Blue    PaletteColor = "blue"
	Indigo  PaletteColor = "indigo"
	Violet  PaletteColor = "violet"
	Purple  PaletteColor = "purple"
	Fuchsia PaletteColor = "fuchsia"
	Pink    PaletteColor = "pink"
	Rose    PaletteColor = "rose"
)

var AllPaletteColors = []struct {
	Value  PaletteColor
	TSName string
}{
	{Noir, "noir"},
	{Emerald, "emerald"},
	{Green, "green"},
	{Lime, "lime"},
	{Orange, "orange"},
	{Amber, "amber"},
	{Yellow, "yellow"},
	{Teal, "teal"},
	{Cyan, "cyan"},
	{Sky, "sky"},
	{Blue, "blue"},
	{Indigo, "indigo"},
	{Violet, "violet"},
	{Purple, "purple"},
	{Fuchsia, "fuschia"},
	{Pink, "pink"},
	{Rose, "rose"},
}

type SufaceColor string

const (
	Slate   SufaceColor = "slate"
	Gray    SufaceColor = "gray"
	Zinc    SufaceColor = "zinc"
	Neutral SufaceColor = "neutral"
	Stone   SufaceColor = "stone"
	Soho    SufaceColor = "soho"
	Viva    SufaceColor = "viva"
	Ocean   SufaceColor = "ocean"
)

var AllSurfaceColors = []struct {
	Value  SufaceColor
	TSName string
}{
	{Slate, "slate"},
	{Gray, "gray"},
	{Zinc, "zinc"},
	{Neutral, "neutral"},
	{Stone, "stone"},
	{Soho, "soho"},
	{Viva, "viva"},
	{Ocean, "ocean"},
}
