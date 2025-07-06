package internal

import (
	"fmt"
	"slices"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type TaskStatus string

const (
	STATUS_TODO        TaskStatus = "todo"
	STATUS_IN_PROGRESS TaskStatus = "in-progress"
	STATUS_DONE        TaskStatus = "done"
)

var (
	purple         = lipgloss.Color("99")
	gray           = lipgloss.Color("245")
	lightGray      = lipgloss.Color("241")
	red            = lipgloss.Color("#FF0000")
	yellow         = lipgloss.Color("#FFCC66")
	White          = lipgloss.Color("#FFFFFF")
	headerStyle    = lipgloss.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
	cellStyle      = lipgloss.NewStyle().Padding(0, 1).Align(lipgloss.Center)
	RowStyle       = cellStyle.Foreground(White)
	statusColStyle = cellStyle.Foreground(red)
	errorStyle     = lipgloss.NewStyle().Bold(true).Padding(1, 0).Foreground(lipgloss.Color(red))
)

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdat"`
	UpdatedAt   time.Time  `json:"updatedat"`
}

func LpError(err error) {
	fmt.Println(errorStyle.Render(err.Error()))
}

func Newtask(id int, description string) *Task {
	return &Task{
		Id:          id,
		Description: description,
		Status:      STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func printTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println(lipgloss.NewStyle().Bold(true).Padding(1, 0).Foreground(lipgloss.Color(yellow)).Render("No tasks found"))
		return
	}

	headers := []string{"ID", "Description", "Status", "Created At", "Updated At"}

	var rows [][]string
	var TaskStatus []TaskStatus

	for _, task := range tasks {
		row := []string{
			fmt.Sprintf("%d", task.Id),
			task.Description,
			string(task.Status),
			task.CreatedAt.Format("2006-01-02 15:04"),
			task.UpdatedAt.Format("2006-01-02 15:04"),
		}

		rows = append(rows, row)
		TaskStatus = append(TaskStatus, task.Status)
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().
			Foreground(purple)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case col == 2:
				return statusColor(TaskStatus[row])
			default:
				return RowStyle
			}
		}).
		Headers(headers...).Rows(rows...)

	fmt.Println(t)
}

func statusColor(status TaskStatus) lipgloss.Style {
	switch status {
	case STATUS_TODO:
		return cellStyle.Foreground(red)
	case STATUS_IN_PROGRESS:
		return cellStyle.Foreground(yellow)
	case STATUS_DONE:
		return cellStyle.Foreground(lightGray)
	default:
		return cellStyle.Foreground(White)
	}
}

func ListTasks() {
	Tasks, err := ReadFromFile()
	if err != nil {
		LpError(err)
		return
	}
	printTasks(Tasks)
}

func AddTask(description string) {
	Tasks, err := ReadFromFile()
	if err != nil {
		LpError(err)
		return
	}

	var maxID int
	for _, t := range Tasks {
		if t.Id > maxID {
			maxID = t.Id
		}
	}
	newtask := Newtask(maxID+1, description)

	Tasks = append(Tasks, *newtask)

	err = SaveToFile(Tasks)
	if err != nil {
		LpError(err)
		return
	}
	printTasks(Tasks)
}

func DeleteTask(ID int) {
	Tasks, err := ReadFromFile()
	if err != nil {
		LpError(err)
		return
	}

	index := -1

	for i, t := range Tasks {
		if ID == t.Id {
			index = i
			break
		}
	}

	if index == -1 {
		LpError(fmt.Errorf("No Item found with this ID"))
		return
	}

	Tasks = slices.Delete(Tasks, index, index+1)

	err = SaveToFile(Tasks)
	if err != nil {
		LpError(err)
		return
	}
	printTasks(Tasks)
}
