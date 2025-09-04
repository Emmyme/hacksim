package game

type Puzzle struct {
	Prompt    string
	Solution  string
	Completed bool
}

func GetPuzzles() []Puzzle {
	return []Puzzle{
		{
			Prompt:    "Decode this: 'khoor' (shifted by 3)",
			Solution:  "hello",
			Completed: false,
		},
		{
			Prompt:    "What comes next in the sequence: 2, 4, 8, 16, ?",
			Solution:  "32",
			Completed: false,
		},
		{
			Prompt:    "Rearrange the letters 'cahk' to form a hacking-related word.",
			Solution:  "hack",
			Completed: false,
		},
		{
			Prompt:    "If 'A' = 1, 'B' = 2, ..., what is the sum of the letters in 'CAB'?",
			Solution:  "6",
			Completed: false,
		},
	}
}
