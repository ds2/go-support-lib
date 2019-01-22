package common

type State int

const (
	Unknown  State = 0
	New      State = 1
	Inactive       = 2
	Prepared State = 3
	Active   State = 4
	Removed  State = 5
)

func (s State) String() string {
	names := [...]string{
		"Unknown",
		"New",
		"Inactive",
		"Prepared",
		"Active",
		"Removed",
	}
	return names[s]
}
