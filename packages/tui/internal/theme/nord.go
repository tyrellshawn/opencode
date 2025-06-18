package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// NordTheme implements the Theme interface with Nord colors.
// It provides both list and dark variants based on the Nord palette.
type NordTheme struct {
	BaseTheme
}

// NewNordTheme creates a new instance of the Nord theme.
func NewNordTheme() *NordTheme {
	// Nord color palette from https://www.nordtheme.com/docs/colors-and-palettes
	polarNight0 := "#2E3440"
	polarNight1 := "#3B4252"
	polarNight2 := "#434C5E"
	polarNight3 := "#4C566A"
	snowStorm0 := "#D8DEE9"
	snowStorm1 := "#E5E9F0"
	snowStorm2 := "#ECEFF4"
	frost0 := "#8FBCBB"
	frost1 := "#88C0D0"
	frost2 := "#81A1C1"
	frost3 := "#5E81AC"
	aurora0 := "#BF616A"
	// aurora1 := "#D08770"
	aurora2 := "#EBCB8B"
	aurora3 := "#A3BE8C"
	aurora4 := "#B48EAD"

	theme := &NordTheme{}

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{Dark: frost1, Light: frost1}
	theme.SecondaryColor = lipgloss.AdaptiveColor{Dark: frost2, Light: frost2}
	theme.AccentColor = lipgloss.AdaptiveColor{Dark: frost0, Light: frost0}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{Dark: aurora0, Light: aurora0}
	theme.WarningColor = lipgloss.AdaptiveColor{Dark: aurora2, Light: aurora2}
	theme.SuccessColor = lipgloss.AdaptiveColor{Dark: aurora3, Light: aurora3}
	theme.InfoColor = lipgloss.AdaptiveColor{Dark: frost0, Light: frost0}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{Dark: snowStorm2, Light: polarNight0}
	theme.TextMutedColor = lipgloss.AdaptiveColor{Dark: snowStorm0, Light: polarNight3}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{Dark: aurora2, Light: aurora2}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{Dark: polarNight0, Light: snowStorm2}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{Dark: polarNight1, Light: snowStorm1}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{Dark: polarNight2, Light: snowStorm0}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{Dark: polarNight3, Light: snowStorm1}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{Dark: frost3, Light: frost3}
	theme.BorderDimColor = lipgloss.AdaptiveColor{Dark: polarNight2, Light: snowStorm0}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{Dark: aurora3, Light: aurora3}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{Dark: aurora0, Light: aurora0}
	theme.DiffContextColor = lipgloss.AdaptiveColor{Dark: polarNight3, Light: snowStorm0}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{Dark: frost3, Light: frost3}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{Dark: frost3, Light: frost3}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{Dark: aurora0, Light: aurora0}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{Dark: polarNight2, Light: snowStorm0}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{Dark: polarNight2, Light: snowStorm0}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{Dark: polarNight1, Light: snowStorm1}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{Dark: polarNight3, Light: snowStorm1}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: polarNight2, Light: snowStorm0}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{Dark: polarNight2, Light: snowStorm0}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{Dark: snowStorm2, Light: polarNight0}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{Dark: frost3, Light: frost3}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{Dark: frost0, Light: frost0}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{Dark: frost1, Light: frost1}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{Dark: aurora3, Light: aurora3}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{Dark: aurora2, Light: aurora2}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{Dark: aurora2, Light: aurora2}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{Dark: aurora0, Light: aurora0}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{Dark: polarNight3, Light: snowStorm1}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{Dark: polarNight2, Light: snowStorm1}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{Dark: polarNight2, Light: snowStorm1}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{Dark: frost1, Light: frost1}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{Dark: frost1, Light: frost1}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{Dark: polarNight2, Light: snowStorm0}

	// Syntax highsnowStorming colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{Dark: polarNight3, Light: snowStorm2}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{Dark: aurora4, Light: aurora4}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{Dark: frost1, Light: frost1}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{Dark: frost0, Light: frost0}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{Dark: aurora2, Light: aurora2}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{Dark: aurora2, Light: aurora2}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{Dark: aurora3, Light: aurora3}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{Dark: aurora4, Light: aurora4}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{Dark: polarNight2, Light: snowStorm2}

	return theme
}

func init() {
	RegisterTheme("nord", NewNordTheme())
}
