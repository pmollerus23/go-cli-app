package task

import "encoding/json"

type TaskStatus int

const (
	Pending TaskStatus = iota
	InProgress
	Completed
	Cancelled
)

var statusStrings = map[TaskStatus]string{
	Pending:    "pending",
	InProgress: "in_progress",
	Completed:  "completed",
	Cancelled:  "cancelled",
}

var stringToStatus = map[string]TaskStatus{
	"pending":     Pending,
	"in_progress": InProgress,
	"completed":   Completed,
	"cancelled":   Cancelled,
}

func (s TaskStatus) String() string {
	if str, ok := statusStrings[s]; ok {
		return str
	}
	return "unknown"
}

func (s TaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *TaskStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	if status, ok := stringToStatus[str]; ok {
		*s = status
		return nil
	}
	*s = Pending
	return nil
}
