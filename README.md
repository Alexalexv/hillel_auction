# hillel_auction

I can`t use make file at local machine, so run project with \
<code>docker-compose up -d</code>

Init documentation
<code>swag init -g .\cmd\main.go</code>

To create migration
<code>goose -dir './db/migrations' postgres 'postgres://sasha:12345678@localhost:5432/auction' create MIGRATION_NAME sql</code> 

Migration up 
<code>goose -dir './db/migrations' postgres 'postgres://sasha:12345678@localhost:5432/auction' up</code>

Migration down 
<code>goose -dir './db/migrations' postgres 'postgres://sasha:12345678@localhost:5432/auction' down</code>