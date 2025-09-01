package game

type Mission struct {
    Title         string
    Description   string
    RequiredSteps []string
    CurrentStep   int
    Completed     bool
}

func GetMissions() []Mission {
    return []Mission{
        {
            Title: "First Hack",
            Description: "Successfully hack into the school Wi-Fi.",
            RequiredSteps: []string{"scan", "hack", "cover"},
            CurrentStep: 0,
            Completed: false,
        },
        {
            Title: "Data Heist",
            Description: "Steal the secret file from the admin server.",
            RequiredSteps: []string{"scan", "hack", "cover"},
            CurrentStep: 0,
            Completed: false,
        },
        {
            Title: "Cover Your Tracks",
            Description: "Erase all logs of your activity.",
            RequiredSteps: []string{"cover"},
            CurrentStep: 0,
            Completed: false,
        },
    }
}


func (m *Mission) Progress(command string) {
    if m.Completed || m.CurrentStep >= len(m.RequiredSteps) {
        return
    }
    if command == m.RequiredSteps[m.CurrentStep] {
        m.CurrentStep++
        if m.CurrentStep == len(m.RequiredSteps) {
            m.Completed = true
        }
    }
}