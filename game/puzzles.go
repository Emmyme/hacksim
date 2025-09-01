package game

type Puzzle struct {
	Prompt      string
	Solution    string
	Description string
}

func GetPuzzles() []Puzzle {
	return []Puzzle{
		{
			Prompt:      "Decode this: 'khoor' (shifted by 3)",
			Solution:    "hello",
			Description: "A Caesar cipher puzzle. Each letter is shifted by 3.",
		},
		{
			Prompt:      "What comes next in the sequence: 2, 4, 8, 16, ?",
			Solution:    "32",
			Description: "Pattern recognition: powers of 2.",
		},
		{
			Prompt:      "Rearrange the letters 'cahk' to form a hacking-related word.",
			Solution:    "hack",
			Description: "Anagram puzzle.",
		},
		{
			Prompt:      "The password is the reverse of 'gnikcah'. What is it?",
			Solution:    "hacking",
			Description: "Reverse the string.",
		},
		{
			Prompt:      "You have 3 switches and 1 light bulb. Only one switch turns on the bulb. How do you find out which one, if you can only enter the room once?",
			Solution:    "Turn on one switch, wait, turn it off, turn on another, enter the room. If bulb is on, it's the second; if warm, it's the first; if off and cold, it's the third.",
			Description: "Classic logic puzzle.",
		},
	}
}
