# LearningProtocolBuffers

## Setup Protoc for Golang on windows

- Download protoc-win32.zip from "https://developers.google.com/protocol-buffers/docs/downloads".
- Unzip the file then add the location of the protoc.exe to your PATH environment variable. In addition, it's not that much necessary, you can copy the protoc.exe to the GOPATH.
- Change/Navigate the CMD/Powershell path to the file then run protoc.exe --version from the command prompt to verify the procedure.
- Verify that your GOPATH environment variable is set.
- Run go get -u github.com/golang/protobuf/protoc-gen-go from the CMD/Powershell. This should install the binary to %GOPATH%/bin.
- Add the exact %GOPATH%/bin to your PATH environment variable.
- Open a new CMD/Powershell, navigate to your .proto file then run protoc --go_out=. *.proto
