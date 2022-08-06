package docker

// `open --background -a Docker` to start docker from command line
// death by ports shell `lsof -t -i tcp:80 | xargs kill`
// rm all docker containers `docker ps -a --format "{{.ID}}" | xargs docker rm`
