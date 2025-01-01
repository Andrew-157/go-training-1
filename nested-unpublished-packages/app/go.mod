module example.com/app

go 1.23.4

replace example.com/server => ../server/

replace example.com/utils => ../utils/

require example.com/server v0.0.0-00010101000000-000000000000

require example.com/utils v0.0.0-00010101000000-000000000000 // indirect
