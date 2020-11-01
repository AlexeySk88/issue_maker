package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Request struct {
	AccessToken  string `yaml:"token"`
	ProjectId    int    `yaml:"project_id"`
	milestoneIid int
	Milestone    string
	Issues       []Issue
}

type Issue struct {
	Id           int
	Title        string
	Description  string
	Labels       []string
	milestoneIid int
	Milestone    string
	Weight       int
}

func (r *Request) validation() error {
	isMilestoneId := r.milestoneIid != 0
	if len(r.AccessToken) == 0 {
		return fmt.Errorf("не заполнено поле 'token'")
	}
	if r.ProjectId == 0 {
		return fmt.Errorf("не заполнено поле 'project_id'")
	}
	if len(r.Issues) == 0 {
		return fmt.Errorf("не добавлено ни одной задачи")
	}
	for _, issue := range r.Issues {
		if len(issue.Title) == 0 {
			return fmt.Errorf("не заполнено поле 'title'")
		}
		if len(issue.Description) == 0 {
			return fmt.Errorf("не заполнено поле 'description'")
		}
		if isMilestoneId && issue.milestoneIid == 0 {
			return fmt.Errorf("не заполнено поле 'milestone_id'")
		}
	}
	return nil
}

func (r *Request) requestParam(index int) string {
	issue := r.Issues[index]
	return fmt.Sprintf(
		"title=%s&description=%s%s%s%s",
		issue.getTitleParam(),
		issue.getDescriptionParam(),
		issue.getLabelsParam(),
		issue.getWeightParam(),
		r.getMilestoneIdParam(index),
	)
}

func (r *Request) getMilestoneIdParam(index int) string {
	baseStr := "&milestone_id="
	if r.Issues[index].milestoneIid != 0 {
		return baseStr + strconv.Itoa(r.Issues[index].milestoneIid)
	} else if r.milestoneIid != 0 {
		return baseStr + strconv.Itoa(r.milestoneIid)
	}
	return ""
}

func (r *Request) updateMilestone(milestones *[]Milestone) error {
	m := map[string]int{}
	for _, milestone := range *milestones {
		m[milestone.Title] = milestone.Id
	}

	iid, err := getMilestone(r.Milestone, m)
	if err != nil {
		return err
	}
	r.milestoneIid = iid

	for index, issue := range r.Issues {
		iid, err := getMilestone(issue.Milestone, m)
		if err != nil {
			return err
		}
		r.Issues[index].milestoneIid = iid
	}

	return nil
}

func getMilestone(key string, m map[string]int) (int, error) {
	if len(key) == 0 {
		return 0, nil
	}

	if val, ok := m[key]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("не найдет milestone с именем %s", key)
}

func (i *Issue) getTitleParam() string {
	return strings.ReplaceAll(i.Title, " ", "%20")
}

func (i *Issue) getDescriptionParam() string {
	return strings.ReplaceAll(i.Description, " ", "%20")
}

func (i *Issue) getLabelsParam() string {
	if len(i.Labels) > 0 {
		return strings.ReplaceAll("&labels="+strings.Join(i.Labels, ","), " ", "%20")
	}
	return ""
}

func (i *Issue) getWeightParam() string {
	if i.Weight != 0 {
		return "&weight=" + strconv.Itoa(i.Weight)
	}
	return ""
}

func (i *Issue) isCreate() bool {
	return i.Id == 0
}

func (i *Issue) isUpdate() bool {
	return i.Id != 0
}

func (i *Issue) copyForWrite(id int) *Issue {
	return &Issue{
		Id:          id,
		Title:       i.Title,
		Description: i.Description,
		Labels:      i.Labels,
		Milestone:   i.Milestone,
		Weight:      i.Weight,
	}
}
