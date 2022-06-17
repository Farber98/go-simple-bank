#######             #######
####### MULTI STAGE #######
#######             #######

####### BUILD STAGE ####### 

# Base image
FROM golang:1.18.2-alpine AS builder
# Working directory inside the image
WORKDIR /app
# Copy everything from the root of our project to workdir
COPY . .
# Build our app to a binary single executable file.
RUN go build -o main main.go


####### RUN STAGE ####### 

FROM alpine
WORKDIR /app
# Copy binary from builder stage (/app/main) to run stage (/app)
COPY --from=builder /app/main .
# Copy env file to load configuration from builder stage (/app/app.env) to run stage (/app)
COPY --from=builder /app/app.env .


# Expose ports used
EXPOSE 8080
# Define default CMD to run when the container starts
CMD [ "/app/main" ]

## OUTPUT: 23MB

