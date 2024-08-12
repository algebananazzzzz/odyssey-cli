package models

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type finishedTaskMsg string
type erroredTaskMsg error

type model struct {
	tasks   []Task
	index   int
	width   int
	height  int
	spinner spinner.Model
	done    bool
}

type Task struct {
	Name string
	Run  func() error
}

var (
	currentTaskNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle            = lipgloss.NewStyle().Margin(1, 2)
	checkMark            = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	errorMark            = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).SetString("✗")
)

func NewModel(tasks []Task) model {
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return model{
		tasks:   tasks,
		spinner: s,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(runTask(m.tasks[m.index]), m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case finishedTaskMsg:
		task := m.tasks[m.index]
		if m.index >= len(m.tasks)-1 {
			// Task completed. We're done!
			m.done = true
			return m, tea.Sequence(
				tea.Printf("%s %s", checkMark, task.Name), // print the last success message
				tea.Quit, // exit the program
			)
		}

		m.index++

		return m, tea.Batch(
			tea.Printf("%s %s", checkMark, task.Name), // print success message above our program
			runTask(m.tasks[m.index]),                 // download the next package
		)
	case erroredTaskMsg:
		task := m.tasks[m.index]
		return m, tea.Sequence(
			tea.Printf("%s %s: %s", errorMark, task.Name, msg.Error()), // print the last error message
			tea.Quit, // exit the program
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	n := len(m.tasks)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.done {
		return doneStyle.Render(fmt.Sprintf("Done! Completed %d tasks.\n", n))
	}

	taskCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n)

	spin := m.spinner.View() + " "
	cellsAvail := max(0, m.width-lipgloss.Width(spin+taskCount))

	taskName := currentTaskNameStyle.Render(m.tasks[m.index].Name)
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Running Task: " + taskName)

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+taskCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + taskCount
}

func runTask(task Task) tea.Cmd {
	return func() tea.Msg {
		err := task.Run()
		if err != nil {
			return erroredTaskMsg(err)
		}
		return finishedTaskMsg(task.Name)
	}
	// d := time.Millisecond * time.Duration(rand.Intn(500)) //nolint:gosec
	// return tea.Tick(d, func(t time.Time) tea.Msg {
	// 	return finishedTaskMsg(task.Name)
	// })
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
