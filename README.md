# Server-Monitor

##Requirements:
    Golang >1.13
    GCC compiler
    LINUX Operating System
    Basic commands that comes Default with LINUX:
        * top
        * uname
        * lscpu
        * free 
        * netstat
    Comes out of the box:
        * ip link show (sudo apt install net-tools)
        * nmcli (sudo apt-get install network-manager)


<a href="https://github.com/manoj2210/Server-Monitor-Backend.git">`git clone https://github.com/manoj2210/Server-Monitor-Backend.git`</a>

Next, cd to the directory of the project.

`$ cd Server-Monitor-Backend`

Next, install the dependencies for the project using the following command:

`$ go run cmd/main.go` 

Keep the backend running and run front end in another terminal. 