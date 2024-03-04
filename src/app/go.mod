module gokain/app

go 1.22.0

require (
	github.com/slowlorizz/gokain-logs v1.0.3
	gokain/worker v0.0.0-00010101000000-000000000000
)

replace gokain/worker => ../worker
