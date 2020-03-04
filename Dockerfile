# Telling to use Docker's golang ready image
FROM golang
# Name and Email of the author
MAINTAINER Iago Agualuza
# Create app folder
RUN mkdir /app
# Copy our file in the host contianer to our contianer
ADD . /app
# Set /app to the go folder as workdir
WORKDIR /app
# Generate binary file from our /app
RUN go build
# Expose the port 80
EXPOSE 80:80
# Run the app binarry file
CMD ["./app"]