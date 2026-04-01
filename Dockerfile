# 1. Use an official Go image that uses Debian (so we can easily install Python)
FROM golang:1.24-bookworm

# 2. Install Python 3 and pip
RUN apt-get update && apt-get install -y python3 python3-pip python3-venv

# 3. Create a virtual environment and install the Kociemba engine
RUN python3 -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"
RUN pip install kociemba

# 4. Set the working directory for your Go app
WORKDIR /app

# 5. Copy your Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# 6. Copy all your project files (Go code, static folder, etc.)
COPY . .

# 7. Build the Go backend
RUN go build -o rubiks-server cmd/api/main.go

# 8. Start the server when the container boots
CMD ["./rubiks-server"]