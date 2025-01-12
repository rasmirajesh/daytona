// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package selection

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/daytonaio/daytona/internal/util"
	"github.com/daytonaio/daytona/pkg/apiclient"
	"github.com/daytonaio/daytona/pkg/views"
	list_view "github.com/daytonaio/daytona/pkg/views/workspace/list"
)

func generateWorkspaceList(workspaces []apiclient.WorkspaceDTO, isMultipleSelect bool, action string) []list.Item {

	// Initialize an empty list of items.
	items := []list.Item{}

	// Populate items with titles and descriptions from workspaces.
	for _, workspace := range workspaces {
		var projectsInfo []string

		if len(workspace.Projects) == 0 {
			continue
		}

		if len(workspace.Projects) == 1 {
			projectsInfo = append(projectsInfo, util.GetRepositorySlugFromUrl(workspace.Projects[0].Repository.Url, true))
		} else {
			for _, project := range workspace.Projects {
				projectsInfo = append(projectsInfo, project.Name)
			}
		}

		// Get the time if available
		uptime := ""
		createdTime := ""
		if workspace.Info != nil && workspace.Info.Projects != nil && len(workspace.Info.Projects) > 0 {
			createdTime = util.FormatTimestamp(workspace.Info.Projects[0].Created)
		}
		if len(workspace.Projects) > 0 && workspace.Projects[0].State != nil {
			if workspace.Projects[0].State.Uptime == 0 {
				uptime = "STOPPED"
			} else {
				uptime = fmt.Sprintf("up %s", util.FormatUptime(workspace.Projects[0].State.Uptime))
			}
		}

		newItem := item[apiclient.WorkspaceDTO]{
			title:          workspace.Name,
			id:             workspace.Id,
			desc:           strings.Join(projectsInfo, ", "),
			createdTime:    createdTime,
			uptime:         uptime,
			target:         workspace.Target,
			choiceProperty: workspace,
		}

		if isMultipleSelect {
			newItem.isMultipleSelect = true
			newItem.action = action
		}

		items = append(items, newItem)
	}

	return items
}

func getWorkspaceProgramEssentials(modelTitle string, actionVerb string, workspaces []apiclient.WorkspaceDTO, footerText string, isMultipleSelect bool) tea.Model {

	items := generateWorkspaceList(workspaces, isMultipleSelect, actionVerb)

	d := ItemDelegate[apiclient.WorkspaceDTO]{}

	l := list.New(items, d, 0, 0)

	l.Styles.FilterPrompt = lipgloss.NewStyle().Foreground(views.Green)
	l.Styles.FilterCursor = lipgloss.NewStyle().Foreground(views.Green)

	l.FilterInput.PromptStyle = lipgloss.NewStyle().Foreground(views.Green)
	l.FilterInput.TextStyle = lipgloss.NewStyle().Foreground(views.Green)

	m := model[apiclient.WorkspaceDTO]{list: l}

	m.list.Title = views.GetStyledMainTitle(modelTitle + actionVerb)
	m.list.Styles.Title = lipgloss.NewStyle().Foreground(views.Green).Bold(true)
	m.footer = footerText

	p, err := tea.NewProgram(m, tea.WithAltScreen()).Run()

	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	return p
}

func selectWorkspacePrompt(workspaces []apiclient.WorkspaceDTO, actionVerb string, choiceChan chan<- *apiclient.WorkspaceDTO) {
	list_view.SortWorkspaces(&workspaces, true)

	p := getWorkspaceProgramEssentials("Select a Workspace To ", actionVerb, workspaces, "", false)
	if m, ok := p.(model[apiclient.WorkspaceDTO]); ok && m.choice != nil {
		choiceChan <- m.choice
	} else {
		choiceChan <- nil
	}
}

func GetWorkspaceFromPrompt(workspaces []apiclient.WorkspaceDTO, actionVerb string) *apiclient.WorkspaceDTO {
	choiceChan := make(chan *apiclient.WorkspaceDTO)

	go selectWorkspacePrompt(workspaces, actionVerb, choiceChan)

	return <-choiceChan
}

func selectWorkspacesFromPrompt(workspaces []apiclient.WorkspaceDTO, actionVerb string, choiceChan chan<- []*apiclient.WorkspaceDTO) {
	list_view.SortWorkspaces(&workspaces, true)

	footerText := lipgloss.NewStyle().Bold(true).PaddingLeft(2).Render(fmt.Sprintf("\n\nPress 'x' to mark workspace.\nPress 'enter' to %s the current/marked workspaces.", actionVerb))
	p := getWorkspaceProgramEssentials("Select Workspaces To ", actionVerb, workspaces, footerText, true)

	m, ok := p.(model[apiclient.WorkspaceDTO])
	if ok && m.choices != nil {
		choiceChan <- m.choices
	} else if ok && m.choice != nil {
		choiceChan <- []*apiclient.WorkspaceDTO{m.choice}
	} else {
		choiceChan <- nil
	}
}

func GetWorkspacesFromPrompt(workspaces []apiclient.WorkspaceDTO, actionVerb string) []*apiclient.WorkspaceDTO {
	choiceChan := make(chan []*apiclient.WorkspaceDTO)

	go selectWorkspacesFromPrompt(workspaces, actionVerb, choiceChan)

	return <-choiceChan
}
