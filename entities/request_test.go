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
