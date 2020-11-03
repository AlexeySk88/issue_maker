package entities

import (
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
	if !strings.Contains(param, "labels=label+1") ||
		!strings.Contains(param, "weight=1") ||
		!strings.Contains(param, "milestone_id=123") ||
		!strings.Contains(param, "title=title+1") ||
		!strings.Contains(param, "description=desc-1") {
		t.Errorf("TestRequestParam failed, got %s", param)
	}

	param = r.RequestParam(1)
	if !strings.Contains(param, "title=title-2") ||
		!strings.Contains(param, "description=desc+2") ||
		!strings.Contains(param, "labels=label+3,label+4") ||
		!strings.Contains(param, "weight=2") ||
		!strings.Contains(param, "milestone_id=321") {
		t.Errorf("TestRequestParam failed, got %s", param)
	}
}

func TestIsCreateAndIsUpdate(t *testing.T) {
	i := Issue{
		Title:       "title-1",
		Description: "desc-1",
		Labels:      []string{"label1", "label2"},
		Weight:      1,
	}

	isCreate := i.IsCreate()
	isUpdate := i.IsUpdate()
	if !isCreate && isUpdate {
		t.Errorf("TestIsCreateAndIsUpdate failed, expected create - %t and update - %t, got %t and %t", isCreate, isUpdate, true, false)
	}

	i.Id = 123
	isCreate = i.IsCreate()
	isUpdate = i.IsUpdate()
	if isCreate && !isUpdate {
		t.Errorf("TestIsCreateAndIsUpdate failed, expected create - %t and update - %t, got %t and %t", isCreate, isUpdate, false, true)
	}
}

func TestGetMilestone(t *testing.T) {
	params := make(map[string]string)
	r := Request{
		milestoneIid: 123,
		Issues: []Issue{
			{
				Title: "title 1",
			},
			{
				milestoneIid: 321,
			},
		},
	}

	r.milestoneIdParam(0, &params)
	if _, ok := params["milestone_id"]; !ok {
		t.Error("TestGetMilestone failed, params not contains milestone_id key")
	}
	expect := "123"
	if params["milestone_id"] != expect {
		t.Errorf("TestGetMilestone failed, expected %s, got %s", expect, params["milestone_id"])
	}
	delete(params, "milestone_id")

	r.milestoneIdParam(1, &params)
	if _, ok := params["milestone_id"]; !ok {
		t.Error("TestGetMilestone failed, params not contains milestone_id key")
	}
	expect = "321"
	if params["milestone_id"] != expect {
		t.Errorf("TestGetMilestone failed, expected %s, got %s", expect, params["milestone_id"])
	}
}

func TestTitleParam(t *testing.T) {
	params := make(map[string]string)
	i := Issue{
		Title: "title 1",
	}

	i.titleParam(&params)
	if _, ok := params["title"]; !ok {
		t.Error("TestGetMilestone failed, params not contains title key")
	}
	expect := "title 1"
	if params["title"] != expect {
		t.Errorf("TestTitleParam failed, expected %s, got %s", expect, params["title"])
	}
}

func TestDescriptionParam(t *testing.T) {
	params := make(map[string]string)
	i := Issue{
		Description: "desc-1",
	}

	i.descriptionParam(&params)
	if _, ok := params["description"]; !ok {
		t.Error("TestGetMilestone failed, params not contains description key")
	}
	expect := "desc-1"
	if params["description"] != expect {
		t.Errorf("TestDescriptionParam failed, expected %s, got %s", expect, params["description"])
	}
}

func TestLabelsParam(t *testing.T) {
	params := make(map[string]string)
	i := Issue{
		Labels: []string{"label1"},
	}

	i.labelsParam(&params)
	if _, ok := params["labels"]; !ok {
		t.Error("TestGetMilestone failed, params not contains labels key")
	}
	expect := "label1"
	if params["labels"] != expect {
		t.Errorf("TestLabelsParam failed, expected %s, got %s", expect, params["labels"])
	}
	delete(params, "labels")

	i.Labels = append(i.Labels, "label 2")
	i.labelsParam(&params)
	if _, ok := params["labels"]; !ok {
		t.Error("TestGetMilestone failed, params not contains labels key")
	}
	expect = "label1,label 2"
	if params["labels"] != expect {
		t.Errorf("TestLabelsParam failed, expected %s, got %s", expect, params["labels"])
	}
}

func TestWeightParam(t *testing.T) {
	params := make(map[string]string)
	i := Issue{
		Weight: 2,
	}

	i.weightParam(&params)
	if _, ok := params["weight"]; !ok {
		t.Error("TestGetMilestone failed, params not contains weight key")
	}
	expect := "2"
	if params["weight"] != expect {
		t.Errorf("TestWeightParam failed, expected %s, got %s", expect, params["weight"])
	}
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

	res := &i == iCopy
	if res {
		t.Errorf("CopyForWrite failed, expected %t, got %t", false, res)
	}

	res = iCopy.Id != 123 || iCopy.Title != i.Title || iCopy.Description != i.Description || iCopy.Weight != i.Weight ||
		iCopy.Milestone != i.Milestone || iCopy.milestoneIid != 0 || !helpers.ArrayEquals(iCopy.Labels, i.Labels)
	if res {
		t.Errorf("CopyForWrite failed, expected %t, got %t", false, res)
	}
}
