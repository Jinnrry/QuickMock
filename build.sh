# linux x86_64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/release/QuickMock_linux_x86_64 main.go

# linux x86
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/release/QuickMock_linux_x86 main.go

# linux arm
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o bin/release/QuickMock_linux_arm main.go

# linux arm64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/release/QuickMock_linux_arm64 main.go

# linux ppc64
CGO_ENABLED=0 GOOS=linux GOARCH=ppc64 go build -o bin/release/QuickMock_linux_ppc64 main.go

# linux ppc64le
CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -o bin/release/QuickMock_linux_ppc64le main.go

# linux mips
CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -o bin/release/QuickMock_linux_mips main.go

# linux mipsle
CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -o bin/release/QuickMock_linux_mipsle main.go

# linux mips64
CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -o bin/release/QuickMock_linux_mips64 main.go

# linux mips64le
CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -o bin/release/QuickMock_linux_mips64le main.go

# windows x86_64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/release/QuickMock_windows_x86_64.exe main.go

# windows x86
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/release/QuickMock_windows_x86.exe main.go

#mac os x86_64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/release/QuickMock_macos_x86_64 main.go

