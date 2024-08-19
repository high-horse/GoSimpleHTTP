.PHONY = server

server :
	cd cmd && go run .

curl-call:
	curl -v http://localhost:8000/test