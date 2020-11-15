package managers

import (
	"github.com/stretchr/testify/require"
	"issue_maker/entities"
	"testing"
)

func TestConsoleManager_ReadConsole(t *testing.T) {
	input := " /alexey/go "
	output := getConsoleManager(input).ReadConsole()
	require.Equal(t, "/alexey/go", output)
}

func TestConsoleManager_CheckRequest(t *testing.T) {
	r := getRequest()
	err := getConsoleManager("y").CheckRequest(r)
	require.NoError(t, err)

	err = getConsoleManager("n").CheckRequest(r)
	require.Error(t, err)
}

func getConsoleManager(output string) *ConsoleManager {
	return NewConsoleManager(func() (string, error) {
		return output, nil
	})
}

func getRequest() *entities.Request {
	return &entities.Request{
		AccessToken:  "token",
		Issues: []entities.Issue{
			{
				Title:       "title 1",
				Description: "desc-1",
				Labels:      []string{"label 1"},
				Milestone:   "milestone_1",
				Weight:      1,
			},
			{
				Title:        "title-2",
				Description:  "desc 2",
				Milestone:    "milestone_2",
				Labels:       []string{"label 3", "label 4"},
				Weight:       2,
			},
		},
	}
}