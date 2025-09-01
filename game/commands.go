package game

type Command struct {
    Name        string
    Description string
    Execute     func(*State) string 
}

func GetCommands() []Command {
    return []Command{
        {
            Name: "scan",
            Description: "Scan the network for devices.",
            Execute: func(s *State) string {
                return "Scanned the network. Devices found: 3."
            },
        },
        {
            Name: "hack",
            Description: "Attempt to hack a device.",
            Execute: func(s *State) string {
                s.Score += 10
                return "Hack successful! Score increased."
            },
        },
        {
            Name: "cover",
            Description: "Cover your tracks to avoid detection.",
            Execute: func(s *State) string {
                return "Tracks covered. You are safe... for now."
            },
        },
    }
}