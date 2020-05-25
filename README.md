# addition
To add two number and get result

Here I used Cobra to create a command line interface.

Later I added command named addition and added some code to get the addition result. 

How to install cobra ? 

1. go get github.com/spf13/cobra/cobra to install cobra package in your gopath. 

2. Create a file name with the Command line tool name you want and inside it : 
     
     cobra init --pkg-name sum

3. It will create main.go file and cmd folder which will have all the commands we give. 

4. to add commands: cobra add addition

5. After adding : go install 

6. To get the addition to work we  should use:  sum addition 2 4    


