module my_project/web_server_two

go 1.13

replace my_project/logger => ../logger

require my_project/logger v0.0.0-00010101000000-000000000000 // indirect
