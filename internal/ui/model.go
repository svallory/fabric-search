package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/svallory/fabric-search/internal/config"
)

var (
	appStyle = lipgloss.NewStyle().Margin(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFDF5")).
				Render
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	list            list.Model
	viewport        viewport.Model
	selectedFolder  string
	readmeContent   string
	systemContent   string
	showingReadme   bool
	quitting        bool
	windowWidth     int
	windowHeight    int
}

func NewModel() Model {
	items := []list.Item{}

	pattern := filepath.Join(config.GetConfigDir(), "patterns", "*")
	folders, _ := filepath.Glob(pattern)

	for _, folder := range folders {
		name := filepath.Base(folder)
		title := strings.ReplaceAll(strings.Title(strings.ReplaceAll(name, "_", " ")), "Md", "MD")
		items = append(items, item{title: title, desc: ""})
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Folders"

	m := Model{list: l, viewport: viewport.New(0, 0)}
	m.viewport.Style = lipgloss.NewStyle().Padding(1, 2)

	return m
}

// ... (rest of the Model methods)