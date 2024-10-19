package main

import (
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/log"
)

type Styles struct{
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

type model struct {
	index int
 	questions []string
 	width 	int
  	height    int
  	answerField  textinput.Model
	styles *Styles
}

func DefaultStyles() *Styles{
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).
		BorderStyle(lipgloss.NormalBorder()).Padding(2).Width(100);
	return s
}

func Next(m *model)  {
	if m.index < len(m.questions) - 1 {
         m.index++
	} else {

	}
}

func New(questions []string) *model {
	ti := textinput.New()
	ti.Placeholder = "Type your answer..."

	ti.Focus()

	styles := DefaultStyles()
	return &model{
		questions:   questions,
		width:       100,
		height:      100,
		answerField: ti,
		styles: styles,
	}
}

func (m model) Init() tea.Cmd{
	log.Info("Init")
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit // Quit the program
		case "enter":
			//log.Warn("Press eneter -> {} " + m.answerField.Value() + "\r")
			m.index++
			m.answerField.SetValue("")
			return m, nil
	    default:
			// Update text input
			var cmd tea.Cmd
			m.answerField, cmd = m.answerField.Update(msg)
			return m, cmd
		}
	default:
		return m, nil
	}
}

func (m model) View() string {
    if m.index == len(m.questions){

    // Create a new style on the renderer.
    style := lipgloss.NewStyle().
		Padding(6).
		Foreground(lipgloss.Color("1")).
		Width(100).
		Background(lipgloss.AdaptiveColor{Light: "228", Dark: "36"})

		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Center,
				style.Render("Done!"),
			),

		)
	}
    if m.width == 0 {
		return "loading..."
	}
	// stack some left-aligned strings together in the center of the window
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.questions[m.index],
			m.styles.InputField.Render(m.answerField.View()),
		),
	)
}

func main() {

	f, err := tea.LogToFile("debug.log","debug");

	if(err != nil) {
		log.Error("Error ", err)
	}

	questions := []string{
		"What's your name?",
		"What's your pet's name?",
	}

	m := New(questions)

	p := tea.NewProgram(m,tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	defer f.Close();
}
