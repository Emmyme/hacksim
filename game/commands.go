package game

type Command struct {
	Name        string
	Description string
	Execute     func(*State)
}

func GetCommands() []Command {
	return []Command{
		{
			Name:        "scan",
			Description: "Scan the network for devices.",
			Execute:     func(s *State) {},
		},
		{
			Name:        "hack",
			Description: "Attempt to hack a device.",
			Execute:     func(s *State) {},
		},
		{
			Name:        "cover",
			Description: "Cover your tracks to avoid detection.",
			Execute:     func(s *State) {},
		},
		{
			Name:        "decrypt",
			Description: "Decrypt a file.",
			Execute:     func(s *State) {},
		},
	}
}
