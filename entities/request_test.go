package entities

import "testing"

func TestRequestParam(t *testing.T) {
	expectParam1 := "title=title-1&description=desc-1&labels=label1,label2&weight=1&milestone_id=123"
	expectParam2 := "title=title-2&description=desc-2&labels=label3,label4&weight=2&milestone_id=321"

	r := Request{
		AccessToken:  "token",
		milestoneIid: 123,
		Issues: []Issue{
			{
				Title:       "title-1",
				Description: "desc-1",
				Labels:      []string{"label1", "label2"},
				Weight:      1,
			},
			{
				Title:        "title-2",
				Description:  "desc-2",
				milestoneIid: 321,
				Labels:       []string{"label3", "label4"},
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
		t.Errorf("TestIsCreate failed, expected create - %t and update - %t, got %t and %t", isCreate, isUpdate, true, false)
	}

	i.Id = 123
	isCreate = i.IsCreate()
	isUpdate = i.IsUpdate()
	if isCreate && !isUpdate {
		t.Errorf("TestIsCreate failed, expected create - %t and update - %t, got %t and %t", isCreate, isUpdate, false, true)
	}
}
