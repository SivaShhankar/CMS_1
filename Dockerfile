# golang image where workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/SivaShhankar/CMS_1

# Setting up working directory
WORKDIR /go/src/github.com/SivaShhankar/CMS_1

# Get godeps for managing and restoring dependencies
RUN go get github.com/tools/godep

# Restore godep dependencies
RUN godep restore 

# Build the taskmanager command inside the container.
RUN go install github.com/SivaShhankar/CMS_1

# Run the taskmanager command when the container starts.
ENTRYPOINT /go/bin/CMS_1

# Service listens on port 8080.
EXPOSE 8080