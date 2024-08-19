# Readme
## Introduction to gRPC
What is **gRPC**?
gRPC is a high-performance, open-source *Remote Procedure Call* (RPC) framework initially developed by **Google**. It allows you to define services using Protocol Buffers (protobuf), a language-neutral and platform-neutral way of serializing structured data, and then automatically generates client and server code in multiple programming languages.

## Key Features of gRPC ✅:
1. **Language Agnostic**:
    gRPC supports multiple programming languages, including Go, Java, Python, C++, Ruby, and more. This allows you to build a service in one language and have clients in another.
2. **Protocol Buffers (Protobuf)**: 
    gRPC uses Protocol Buffers as its Interface Definition Language (IDL) by default, enabling efficient serialization of messages and strong data types.
3. **Automatic Code Generation**:
    gRPC generates client and server code from the .proto files, making it easy to build services with consistent APIs across different languages.
4. **Bi-Directional Streaming**:
    gRPC supports streaming in multiple forms:
5. **Unary RPC**: A single request followed by a single response.
6. **HTTP/2 Based**:
    gRPC is built on top of HTTP/2, which provides features like multiplexing, flow control, header compression, and low-latency communication, making it highly efficient.
7. **Built-in Authentication**:
    gRPC supports various authentication mechanisms, including SSL/TLS for transport security and token-based authentication.
8. **Cross-Platform**:
    gRPC works across different environments, including cloud, on-premises, mobile, and web applications.
9. **Load Balancing and Pluggable Mechanisms**:
    gRPC has built-in support for load balancing and can be extended with custom mechanisms for authentication, authorization, and more.

## Running the server 🚀
1. To install **Go** in your device, visit https://go.dev/doc/install and follow the steps.
2. To install **Protocol Buffer** compiler *protoc*, visit https://grpc.io/docs/protoc-installation/ and follow the steps corresponding to your OS.

After the installation is complete, run the following in your current working directory:
1. Install the protocol compiler plugins for Go using the following commands:
    ```
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```
2. Update your `PATH` so that the `protoc` compiler can find the plugins:
    ```
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```
3. To modify the RPC (*optional*), make required changes to the `proto/invoicer.proto` file and run the following command to modify the generated files accordingly:
    ```
    protoc --go_out=. --go_opt=paths=source_relative \    
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/invoicer.proto
    ```
4. To start the server, run:
    ```
    go run main.go
    ```
    This shall start a TCP connection on the port 8080.

## Testing the server 🧪
Once you successfully run the `go run` command, you won't see anything popping up in your terminal and might think that my server has hanged.😟

No worries! Fortunately this is the known behaviour 😮‍💨
#### But how to test your RPC?
There are a variety of tools such as:

1. BloomRPC: https://appimage.github.io/BloomRPC/
2. Kreya: https://kreya.app/
3. Postman: https://www.postman.com/downloads/
and many more...

I am using Kreya to test my gRPC server. To download and install Kreya, visit: https://kreya.app/downloads/.

### Steps to test:
1. Open Kreya; You will see the below screen:

    <img src="https://i.ibb.co/ZBG6J5f/Screenshot-2024-08-19-at-3-52-51-PM.png" width="600" height="auto" alt="step-1"/>

2. Click on Open project and browse to `Learn-Go/go-grpc/kreya` and select the `go-grpc.krproj` file. Once you do so, the following screen will appear:

    <img src="https://i.ibb.co/zsRNM4T/Screenshot-2024-08-19-at-4-08-07-PM.png" width="600" height="auto" alt="step-2">

3. Click on the arrow beside the `Invoicer` tab, a dropdown will appear which shall contain the RPCs to test. 

4. Click on the RPC you want to test. You should see a screen like this (you won't see the right panel before calling):
    <img src="https://i.ibb.co/W37DzC3/Screenshot-2024-08-19-at-4-11-15-PM.png" width="600" height="auto" alt="step-3">
5. Clicking on send shall give you a result similar to the one shown on the right side of the image.

### HAPPY CODING! 💻


