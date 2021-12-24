# Make file

# Build docker image
.PHONE: docker
docker-build-%: Dockerfile
	@docker build -f Dockerfile -t mxudong/object-mocker:v$* .

output:
	@echo 'mkdir output'