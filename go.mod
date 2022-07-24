module github.com/elgs/wsld

go 1.18

replace github.com/elgs/wsl => ../wsl

replace github.com/elgs/gorediscache => ../gorediscache

require (
	github.com/elgs/gosqljson v0.0.0-20220712125658-2f85b34a6a73
	github.com/elgs/gostrgen v0.0.0-20220325073726-0c3e00d082f6
	github.com/elgs/wsl v0.0.0-20220716111606-cb56e5c6a4fc
	github.com/go-sql-driver/mysql v1.6.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/elgs/gorediscache v0.0.0-20220717011252-82d76361e241 // indirect
	github.com/elgs/gosplitargs v0.0.0-20161028071935-a491c5eeb3c8 // indirect
	github.com/gomodule/redigo v1.8.9 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
