# Use a base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Install Python and pip
RUN apk add --no-cache python3 py3-pip

# Copy the application files to the container
COPY hello-world.py /app

# Set the entrypoint command
CMD ["python3", "hello-world.py"]

