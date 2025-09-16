# HackSim

A terminal-based hacking simulator game built with Go.

## Screenshots

![HackSim Mission](assets/image1.png)

![HackSim Won](assets/image2.png)

## Installation

```bash
git clone https://github.com/Emmyme/hacksim.git
cd hacksim
go mod tidy
go run .
```

## Dependencies

This project uses the following Go modules:

### Direct Dependencies
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** v1.3.6 - Terminal UI framework for building interactive command-line applications
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** v1.1.0 - Terminal styling library for beautiful CLI interfaces

### Additional Requirements
- **Go 1.25.0** or higher


## How to Play

Complete missions by entering commands in the correct order, then solve puzzles.

### Commands
- `scan` - Scan network
- `hack` - Hack device  
- `cover` - Cover tracks
- `decrypt` - Decrypt file

### Mission Sequences
1. **First Hack**: `scan` → `hack` → `cover`
2. **Data Heist**: `scan` → `hack` → `decrypt` → `cover`  
3. **Cover Tracks**: `cover`
4. **Final Breach**: `scan` → `hack` → `decrypt` → `hack` → `cover` 

**Note**: Wrong command order resets the game!

### Controls
- Type and press Enter to submit
- Backspace to delete
- `q` or Ctrl+C to quit

## Winning

Complete all missions and puzzles to win!
