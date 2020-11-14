package entities

import (
	"github.com/stretchr/testify/require"
	"issue_maker/helpers"
	"strings"
	"testing"
)

func TestRequest_RequestParam(t *testing.T) {
	r := getRequest()

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

func TestIssue_CopyForWrite(t *testing.T) {
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
	require.Equal(t, 123, iCopy.Id)
	require.Equal(t, i.Title, iCopy.Title)
	require.Equal(t, i.Description, iCopy.Description)
	require.Equal(t, i.Weight, iCopy.Weight)
	require.Equal(t, i.Milestone, iCopy.Milestone)
	require.Equal(t, 0, iCopy.milestoneIid)
	require.True(t, helpers.ArrayEquals(iCopy.Labels, i.Labels))
}

func TestRequest_UpdateMilestone(t *testing.T) {
	r := getRequest()
	ms := getMilestones()

	err := r.UpdateMilestone(&ms)
	require.NoError(t, err)
	require.Equal(t, 1, r.Issues[0].milestoneIid)
	require.Equal(t, 2, r.Issues[1].milestoneIid)
}

func TestRequest_Validation(t *testing.T) {
	r := Request{}
	require.Error(t, r.Validation())

	r.ProjectId = 123
	require.Error(t, r.Validation())

	r.AccessToken = "abc"
	require.Error(t, r.Validation())

	r.Issues = append(r.Issues, Issue{})
	require.Error(t, r.Validation())

	r.Issues[0].Id = 321
	require.NoError(t, r.Validation())

	r.Issues[0].Id = 0
	r.Issues[0].Title = "title"
	require.Error(t, r.Validation())

	r.Issues[0].Description = "desc"
	require.Error(t, r.Validation())

	r.Issues[0].Milestone = "milestone_1"
	require.NoError(t, r.Validation())
}

func getRequest() Request {
	return Request{
		AccessToken:  "token",
		milestoneIid: 123,
		Issues: []Issue{
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
				milestoneIid: 321,
				Labels:       []string{"label 3", "label 4"},
				Weight:       2,
			},
		},
	}
}

func getMilestones() []Milestone {
	return []Milestone{
		{1, 1, "milestone_1"},
		{2, 2, "milestone_2"},
	}
}
