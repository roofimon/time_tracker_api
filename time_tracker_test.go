package time_tracker

import (
	"testing"
)

type MockRepository struct {
	List  []interface{}
	Count int
}

func (m *MockRepository) Insert(data interface{}) {
	m.List[0] = data
	m.Count = 1
}

func TestItShouldCreateNewRecordWhenCheckInFirstTime(t *testing.T) {
	mockRepository := MockRepository{
		make([]interface{}, 1),
		0,
	}

	timeTracker := TimeTracker{&mockRepository}
	timeTracker.CheckIn("roofimon")

	roofimon_checkin := mockRepository.Count
	if roofimon_checkin != 1 {
		t.Errorf("Expect one record but get %v", roofimon_checkin)
	}
}
