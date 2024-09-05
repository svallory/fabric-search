package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
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

type model struct {
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

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		m.viewport.Width = msg.Width - 4
		m.viewport.Height = msg.Height - 10
		m.list.SetWidth(msg.Width / 2)
		m.list.SetHeight(msg.Height - 10)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.selectedFolder = i.title
				m.readmeContent = readMarkdownFile(filepath.Join(os.Getenv("HOME"), ".config", "fabric", "patterns", strings.ReplaceAll(strings.ToLower(i.title), " ", "_"), "README.md"))
				m.systemContent = readMarkdownFile(filepath.Join(os.Getenv("HOME"), ".config", "fabric", "patterns", strings.ReplaceAll(strings.ToLower(i.title), " ", "_"), "system.md"))
				m.showingReadme = true
			}
		case "tab":
			m.showingReadme = !m.showingReadme
		case "n":
			fmt.Println(m.selectedFolder)
		case "p":
			fmt.Println(filepath.Join(os.Getenv("HOME"), ".config", "fabric", "patterns", strings.ReplaceAll(strings.ToLower(m.selectedFolder), " ", "_")))
		case "r":
			fmt.Println(m.readmeContent)
		case "s":
			fmt.Println(m.systemContent)
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return "Goodbye!"
	}

	var content string
	if m.selectedFolder != "" {
		if m.showingReadme {
			content = m.readmeContent
		} else {
			content = m.systemContent
		}
	}

	if content == "" {
		content = "No content available."
	}

	m.viewport.SetContent(content)

	return appStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.list.View(),
			lipgloss.JoinVertical(
				lipgloss.Left,
				titleStyle.Render(fmt.Sprintf("Viewing: %s (%s)", m.selectedFolder, map[bool]string{true: "README", false: "system"}[m.showingReadme])),
				m.viewport.View(),
				"\nPress 'tab' to toggle between README and system",
				"\nKeybindings: n (name), p (path), r (readme), s (system), q (quit)",
			),
		),
	)
}

func main() {
	items := []list.Item{}

	pattern := filepath.Join(os.Getenv("HOME"), ".config", "fabric", "patterns", "*")
	folders, _ := filepath.Glob(pattern)

	for _, folder := range folders {
		name := filepath.Base(folder)
		title := strings.ReplaceAll(strings.Title(strings.ReplaceAll(name, "_", " ")), "Md", "MD")
		items = append(items, item{title: title, desc: ""})
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Folders"

	m := model{list: l, viewport: viewport.New(0, 0)}
	m.viewport.Style = lipgloss.NewStyle().Padding(1, 2)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func readMarkdownFile(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}

	rendered, _ := glamour.Render(string(content), "dark")
	return rendered
}