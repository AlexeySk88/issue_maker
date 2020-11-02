package entities

import (
	"issue_maker/helpers"
	"testing"
)

func TestRequestParam(t *testing.T) {
	expectParam1 := "title=title%201&description=desc-1&labels=label%201&weight=1&milestone_id=123"
	expectParam2 := "title=title-2&description=desc%202&labels=label%203,label%204&weight=2&milestone_id=321"

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
	if param != expectParam1 {
		t.Errorf("TestRequestParam failed, expected %s, got %s", expectParam1, param)
	}

	param = r.RequestParam(1)
	if param != expectParam2 {
		t.Errorf("TestRequestParam failed, expected %s, got %s", expectParam2, param)
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
	var params []string
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
	r.milestoneIdParam(1, &params)

	expect := "milestone_id=123"
	if params[0] != expect {
		t.Errorf("TestGetMilestone failed, expected %s, got %s", expect, params[0])
	}

	expect = "milestone_id=321"
	if params[1] != expect {
		t.Errorf("TestGetMilestone failed, expected %s, got %s", expect, params[1])
	}
}

func TestTitleParam(t *testing.T) {
	var params []string
	i := Issue{
		Title: "title 1",
	}
	i.titleParam(&params)

	expect := "title=title 1"
	if params[0] != expect {
		t.Errorf("TestTitleParam failed, expected %s, got %s", expect, params[0])
	}
}

func TestDescriptionParam(t *testing.T) {
	var params []string
	i := Issue{
		Description: "desc-1",
	}
	i.descriptionParam(&params)

	expect := "description=desc-1"
	if params[0] != expect {
		t.Errorf("TestDescriptionParam failed, expected %s, got %s", expect, params[0])
	}
}

func TestLabelsParam(t *testing.T) {
	var params []string
	i := Issue{
		Labels: []string{"label1"},
	}
	i.labelsParam(&params)

	expect := "labels=label1"
	if params[0] != expect {
		t.Errorf("TestLabelsParam failed, expected %s, got %s", expect, params[0])
	}

	i.Labels = append(i.Labels, "label 2")
	i.labelsParam(&params)
	expect = "labels=label1,label 2"
	if params[1] != expect {
		t.Errorf("TestLabelsParam failed, expected %s, got %s", expect, params[1])
	}
}

func TestWeightParam(t *testing.T) {
	var params []string
	i := Issue{
		Weight: 2,
	}
	i.weightParam(&params)

	expect := "weight=2"
	if params[0] != expect {
		t.Errorf("TestWeightParam failed, expected %s, got %s", expect, params[0])
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
