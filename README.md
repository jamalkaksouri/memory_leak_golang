# Memory Leak checking in golang service

Sometimes, we find pods restarting several times a day without any error. The memory consumption keeps going up, until it reaches the memory limit.\
We have a process to diagnose memory leak for Go services. Tools such as pprof and minikube can help us finding the root cause.

## Dependencies
- Gin framework
- Graphviz

## Installation Graphviz in windows
1. Visit the [download location](https://gitlab.com/graphviz/graphviz/-/releases)
2. Download and run the [32-bit](https://gitlab.com/api/v4/projects/4207231/packages/generic/graphviz-releases/6.0.2/windows_10_cmake_Release_graphviz-install-6.0.2-win32.exe) or [64-bit](https://gitlab.com/api/v4/projects/4207231/packages/generic/graphviz-releases/6.0.2/windows_10_cmake_Release_graphviz-install-6.0.2-win64.exe) exe file.
3. Ignore any security warnings you might get.
4. During the installation, make sure you select Add Graphviz top the system PATH for current user.
5. When the installation is finished, start CMD as an administrator
6. Run command `dot`

**Note**
Make sure that during installation checked "Add Graphviz to the system PATH for current user"

![](../754680b6a7f66af5318b6deed62a3e8f5c0d34f2.png)

# Usage

1. One time install
```bash 
go install github.com/google/pprof@latest
```
2. Dump heap to a file(ensure that application is run!-using new terminal)
```bash 
curl http://<HOSTNAME>:<PORT>/debug/pprof/heap > heap.out
```
3. Use pprof to interact with heap
```bash 
go tool pprof heap.out
```
4. Inside the new command prompt
```bash 
png
```

💡 Everytime for reporting new result you should repeat levels 2,3,4\
💡 After run `go run main.go` go to [localhost](http://localhost:8080/append-slice) on browser then repeat previous level

# routes

🔗 **[http://localhost:8080/append-slice](http://localhost:8080/append-slice)**\
🔗 **[http://localhost:8080/hanging](http://localhost:8080/hanging)**\
🔗 **[http://localhost:8080/streams](http://localhost:8080/streams)**


# Tips
In the diagram, each box is a function, and each arrow means a function call. The bigger the box, the higher memory usage. From the graph above, the blame goes to runtime function “allocm” (the largest box near the bottom).
Once we found which function is causing the problem, we can check how the memory is leaked.
