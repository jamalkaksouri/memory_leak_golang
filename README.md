# Memory Leak checking in golang service

##Dependencies
- Gin framework
- Graphviz

## Installation Graphviz in windows
1. Visit the download location 2.0k
2. Download and run the 32-bit 224 or 64-bit 3.9k exe file.
3. Ignore any security warnings you might get.
4. During the installation, make sure you select Add Graphviz top the system PATH for current user.
5. When the installation is finished, start CMD as an administrator
6. Restart any applications or Command prompts where you want to use Graphviz.

**Note**
Make sure that during installation checked "Add Graphviz to the system PATH for current user"

![](../754680b6a7f66af5318b6deed62a3e8f5c0d34f2.png)

## Usage

# one time install
go install github.com/google/pprof@latest
# Dump heap to a file
curl http://<HOSTNAME>:<PORT>/debug/pprof/heap > heap.out
# Use pprof to interact with heap
go tool pprof heap.out
# Inside the new command prompt
png
# for other output type file using help command