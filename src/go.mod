module gokain/app

go 1.22.0

require (
	github.com/slowlorizz/gokain-logs v1.0.3
	gokain/feat v0.0.0-00010101000000-000000000000
	gokain/utils v0.0.0-00010101000000-000000000000
)

require gokain/types v0.0.0-00010101000000-000000000000 // indirect

replace gokain/worker => ../worker

replace gokain/utils => ./lib/utils

replace gokain/types => ./lib/types

replace gokain/feat => ./lib/feat
