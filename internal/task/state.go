/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package task

type (
	// TODO: un pool ?
	State struct {
		Name         string
		Schedule     string
		NextRun      string
		LastSuccess  string
		LastFailure  string
		ErrMsg       string
		SuccessCount int
		FailureCount int
		Disabled     bool
	}
)

func NewState(task Task, nextRun string) *State {
	return &State{
		Name:     task.Name(),
		Disabled: task.Disabled(),
		Schedule: task.Schedule(),
		NextRun:  nextRun,
	}
}

func (s *State) AfterRun(r *Result) {
	if r.Success {
		s.SuccessCount += 1
		s.LastSuccess = r.DateTime
	} else {
		s.FailureCount += 1
		s.LastFailure = r.DateTime
	}

	s.NextRun = r.NextRun
}

func (s *State) Clone() *State {
	clone := *s
	return &clone
}

type (
	// TODO: un pool ?
	Result struct {
		Error    error
		DateTime string
		NextRun  string
		Success  bool
	}
)

/*
####### END ############################################################################################################
*/
