FROM golang as firststage
WORKDIR /work
ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .
#
# Step #2: Copy the executable into a minimal image (less than 5MB) 
#         which doesn't contain the build tools and artifacts
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=firststage /work/myapp .
CMD ["./Obsidian"]  
