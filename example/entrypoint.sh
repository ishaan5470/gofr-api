# its not enough to say only depends on and we have to say wait  for for database to get started  and use this third prty package called  wait for 
wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"


#this package allows us to update our new.go file and it will automaticlaly update automatically when we make a change to it.
# as in go we need to cancel and then restart the terminal but with this package you dnt need to do it 

CompileDaemon --build="go build -o main main.go" --command=./main