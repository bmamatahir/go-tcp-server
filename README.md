# Go TCP Server
This is a simple Go TCP server that listens for incoming connections and prints received data.

### Building the Docker Image
```bash
docker build -t go-tcp-server .
```

### Running the Docker Container and enable verbose logging
```bash
docker run -p 8080:8080 go-tcp-server -port 8080 -v
```
This will print additional information, including timestamps and connection details.

### Connection to the server
```bash
nc <your_server_ip> <port>
```