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
		List:  make([]interface{}, 1),
		Count: 0,
	}

	timeTracker := TimeTracker{&mockRepository}
	timeTracker.CheckIn("roofimon")

	roofimon_checkin := mockRepository.Count
	//assert.Equal(1, roofimon_checkin, "they should be equal")
	if roofimon_checkin != 1 {
		t.Errorf("Expect one record but get %v", roofimon_checkin)
	}
}
