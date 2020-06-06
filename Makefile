build:
	sudo docker build -t russian-doll .

run:
	sudo docker run -p 8080:8080 russian-doll &