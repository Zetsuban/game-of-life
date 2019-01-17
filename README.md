# Game Of Life
Terminal version of Conway's Game Of Life in go using termbox-go

## Features
- Runs in terminal !!
- Presets via flags
- Step by step or running

## Install
Use go get to install this
```
go get github.com/zetsuban/game-of-life
```

## How to Use

### Launching
Use `life` to launch it
The following **flags** are supported :

flag | description
-----|------------
-h   | print a little help message
-c   | set custom column number
-r   | set custom row number
-g   | glider preset
-t   | tumbler preset
-e   | exploder preset
-S   | toggles set by step
-s   | time in milliseconds between each generations
-cmd | launch the game without using termbox-go

### Keymap
To be used while the game is running

key      | description
---------|------------
spc      | toggles step by step
p        | while step by step go to next gen
esc      | quit the game
ctrl + q | quit the game

## TODO

- [] Implement more presets
- [] Restart World without quiting the game
- [] Change game speed within game
- [] Themes (???)
- [x] ~~Use termbox for the display~~