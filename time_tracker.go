package time_tracker

type Inserter interface {
	Insert(docs ...interface{}) error
}

type TimeTracker struct {
	repository Inserter
}

func (timeTracker TimeTracker) CheckIn(username string) {
	timeTracker.repository.Insert(map[string]string{"username": username})
}
