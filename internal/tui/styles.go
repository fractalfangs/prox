// Package tui provides a terminal-based management interface for CLIProxyAPI.
package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Color palette inspired by lio.cat: amber accents on near-black in dark
// terminals, cobalt blue accents in light terminals, no filled surfaces.
var (
	colorPrimary  = lipgloss.AdaptiveColor{Light: "#1447E6", Dark: "#FFB900"} // accent
	colorOnAccent = lipgloss.AdaptiveColor{Light: "#FBBF24", Dark: "#121212"} // text drawn on the accent
	colorSuccess  = lipgloss.AdaptiveColor{Light: "#15803D", Dark: "#86EFAC"}
	colorWarning  = lipgloss.AdaptiveColor{Light: "#C2410C", Dark: "#FDBA74"}
	colorError    = lipgloss.AdaptiveColor{Light: "#B91C1C", Dark: "#FCA5A5"}
	colorInfo     = colorPrimary
	colorMuted    = lipgloss.AdaptiveColor{Light: "#737373", Dark: "#6B6B6B"}
	colorText     = lipgloss.AdaptiveColor{Light: "#121212", Dark: "#FFFFFF"}
	colorSubtext  = lipgloss.AdaptiveColor{Light: "#404040", Dark: "#A3A3A3"}
	colorBorder   = lipgloss.AdaptiveColor{Light: "#A3A3A3", Dark: "#2E2E2E"}
)

// dotSeparator joins inline items, mirroring the "a ⋅ b ⋅ c" lines on lio.cat.
const dotSeparator = " ⋅ "

// Tab bar styles: tabs render as plain lowercase links, the active one underlined.
var (
	tabActiveStyle = lipgloss.NewStyle().
			Bold(true).
			Underline(true).
			Foreground(colorPrimary).
			Transform(strings.ToLower).
			MarginRight(3)

	tabInactiveStyle = lipgloss.NewStyle().
				Foreground(colorSubtext).
				Transform(strings.ToLower).
				MarginRight(3)

	tabBarStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingTop(1).
			PaddingBottom(1)
)

// Content styles
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorText).
			Transform(strings.ToLower).
			MarginBottom(1)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(colorSubtext)

	labelStyle = lipgloss.NewStyle().
			Foreground(colorSubtext).
			Width(24)

	valueStyle = lipgloss.NewStyle().
			Foreground(colorText)

	headingStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary).
			Transform(strings.ToLower)

	errorStyle = lipgloss.NewStyle().
			Foreground(colorError).
			Bold(true)

	successStyle = lipgloss.NewStyle().
			Foreground(colorSuccess)

	warningStyle = lipgloss.NewStyle().
			Foreground(colorWarning)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(colorMuted).
			PaddingLeft(1).
			PaddingRight(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(colorMuted)

	// accentBadgeStyle inverts the palette, e.g. amber block with dark text.
	accentBadgeStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(colorOnAccent).
				Background(colorPrimary).
				Padding(0, 1)
)

// Log level styles
var (
	logDebugStyle = lipgloss.NewStyle().Foreground(colorMuted)
	logInfoStyle  = lipgloss.NewStyle().Foreground(colorSubtext)
	logWarnStyle  = lipgloss.NewStyle().Foreground(colorWarning)
	logErrorStyle = lipgloss.NewStyle().Foreground(colorError)
)

// Table styles
var (
	tableHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(colorPrimary).
				Transform(strings.ToLower).
				BorderBottom(true).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(colorBorder)

	tableCellStyle = lipgloss.NewStyle().
			Foreground(colorText).
			PaddingRight(2)

	tableSelectedStyle = lipgloss.NewStyle().
				Foreground(colorOnAccent).
				Background(colorPrimary).
				Bold(true)
)

func logLevelStyle(level string) lipgloss.Style {
	switch level {
	case "debug":
		return logDebugStyle
	case "info":
		return logInfoStyle
	case "warn", "warning":
		return logWarnStyle
	case "error", "fatal", "panic":
		return logErrorStyle
	default:
		return logInfoStyle
	}
}
