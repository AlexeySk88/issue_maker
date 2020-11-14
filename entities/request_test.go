package entities

import (
	"github.com/stretchr/testify/require"
	"issue_maker/helpers"
	"strings"
	"testing"
)

func TestRequestParam(t *testing.T) {
	r := Request{
		AccessToken:  "token",
		milestoneIid: 123,
		Issues: []Issue{
			{
				Title:       "title 1",
				Description: "desc-1",
				Labels:      []string{"label 1"},
				Weight:      1,
			},
			{
				Title:        "title-2",
				Description:  "desc 2",
				milestoneIid: 321,
				Labels:       []string{"label 3", "label 4"},
				Weight:       2,
			},
		},
	}

	param := r.RequestParam(0)
	require.True(t, strings.Contains(param, "labels=label+1"))
	require.True(t, strings.Contains(param, "weight=1"))
	require.True(t, strings.Contains(param, "milestone_id=123"))
	require.True(t, strings.Contains(param, "title=title+1"))
	require.True(t, strings.Contains(param, "description=desc-1"))

	param = r.RequestParam(1)
	require.True(t, strings.Contains(param, "title=title-2"))
	require.True(t, strings.Contains(param, "description=desc+2"))
	require.True(t, strings.Contains(param, "labels=label+3,label+4"))
	require.True(t, strings.Contains(param, "weight=2"))
	require.True(t, strings.Contains(param, "milestone_id=321"))
}

func TestIsCreateAndIsUpdate(t *testing.T) {
	i := Issue{
		Title:       "title-1",
		Description: "desc-1",
		Labels:      []string{"label1", "label2"},
		Weight:      1,
	}
	require.True(t, i.IsCreate())
	require.False(t, i.IsUpdate())

	i.Id = 123
	require.False(t, i.IsCreate())
	require.True(t, i.IsUpdate())
}

func TestCopyForWrite(t *testing.T) {
	i := Issue{
		Title:        "title-2",
		Description:  "desc 2",
		milestoneIid: 321,
		Milestone:    "mil-321",
		Labels:       []string{"label 3", "label 4"},
		Weight:       2,
	}
	iCopy := i.CopyForWrite(123)
	require.False(t, &i == iCopy)
	require.True(t, iCopy.Id == 123)
	require.True(t, iCopy.Title == i.Title)
	require.True(t, iCopy.Description == i.Description)
	require.True(t, iCopy.Weight == i.Weight)
	require.True(t, iCopy.Milestone == i.Milestone)
	require.True(t, iCopy.milestoneIid == 0)
	require.True(t, helpers.ArrayEquals(iCopy.Labels, i.Labels))
}
