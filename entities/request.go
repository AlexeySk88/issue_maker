package entities

import (
	"fmt"
	"issue_maker/helpers"
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

func (r *Request) Validation() error {
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
		if issue.Id != 0 {
			continue
		}
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

func (r *Request) RequestParam(index int) string {
	issue := r.Issues[index]
	var params []string
	issue.titleParam(&params)
	issue.descriptionParam(&params)
	issue.labelsParam(&params)
	issue.weightParam(&params)
	r.milestoneIdParam(index, &params)
	param := strings.Join(params, "&")
	return helpers.ReplaceForRestParam(param)
}

func getMilestone(key string, m map[string]int) (int, error) {
	if len(key) == 0 {
		return 0, nil
	}

	if val, ok := m[key]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("не найден milestone с именем %s", key)
}

func (i *Issue) titleParam(params *[]string) {
	if len(i.Title) > 0 {
		*params = append(*params, "title="+i.Title)
	}
}

func (i *Issue) descriptionParam(params *[]string) {
	if len(i.Description) > 0 {
		*params = append(*params, "description="+i.Description)
	}
}

func (i *Issue) labelsParam(params *[]string) {
	if len(i.Labels) > 0 {
		*params = append(*params, "labels="+strings.Join(i.Labels, ","))
	}
}

func (i *Issue) weightParam(params *[]string) {
	if i.Weight != 0 {
		*params = append(*params, "weight="+strconv.Itoa(i.Weight))
	}
}

func (r *Request) milestoneIdParam(index int, params *[]string) {
	baseStr := "milestone_id="
	if r.Issues[index].milestoneIid != 0 {
		*params = append(*params, baseStr+strconv.Itoa(r.Issues[index].milestoneIid))
	} else if r.milestoneIid != 0 {
		*params = append(*params, baseStr+strconv.Itoa(r.milestoneIid))
	}
}

func (r *Request) UpdateMilestone(milestones *[]Milestone) error {
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

func (i *Issue) IsCreate() bool {
	return i.Id == 0
}

func (i *Issue) IsUpdate() bool {
	return i.Id != 0
}

func (i *Issue) CopyForWrite(id int) *Issue {
	return &Issue{
		Id:          id,
		Title:       i.Title,
		Description: i.Description,
		Labels:      i.Labels,
		Milestone:   i.Milestone,
		Weight:      i.Weight,
	}
}
