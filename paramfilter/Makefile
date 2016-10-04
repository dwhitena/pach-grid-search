all: compile docker push

compile:
	CGO_ENABLED=0 go build -o paramfilter

docker:
	sudo docker build --no-cache=true --force-rm=true -t dwhitena/paramfilter .

push:
	sudo docker push dwhitena/paramfilter
