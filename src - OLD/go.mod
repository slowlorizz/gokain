module gokain/app

go 1.22.0

require (
	github.com/gizak/termui/v3 v3.1.0
	gokain/feat v0.0.0-00010101000000-000000000000
	gokain/tui v0.0.0-00010101000000-000000000000
	gokain/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/mattn/go-runewidth v0.0.2 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/nsf/termbox-go v0.0.0-20190121233118-02980233997d // indirect
	gokain/types v0.0.0-00010101000000-000000000000
)

replace gokain/worker => ../worker

replace gokain/utils => ./lib/utils

replace gokain/types => ./lib/types

replace gokain/feat => ./lib/feat

replace gokain/tui => ./lib/tui
