## P04- Sequential vs. Concurrent Image Downloader
### Delton Hughes 

### Description:

- In this program we will make two functions which will download 
images in a `sequential` and `concurrent` fasion. We will also 
time said function and include proper error handling if images 
do not download correctly. 


### Files
|   #   | Files/ Folders | Description                      |
| :---: | --------- | -------------------------------- |
|   1   | [main.go](./main.go)  | Main program file. |
|   2   | [go.mod](./go.mod/)| Init file        |                  
|   3   | [conImg1.jpg](./conImg1.jpg/)| Concurrent Image   |  
|   4   | [conImg2.jpg](./conImg2.jpg/)| Concurrent Image   | 
|   5   | [conImg3.jpg](./conImg3.jpg/)| Concurrent Image   | 
|   6   | [conImg4.jpg](./conImg4.jpg/)| Concurrent Image   | 
|   7   | [conImg5.jpg](./conImg5.jpg/)| Concurrent Image   | 
|   8   | [seqImg1.jpg](./seqImg1.jpg/)| Sequential Image   | 
|   9   | [seqImg2.jpg](./seqImg2.jpg/)| Sequential Image   | 
|   10  | [seqImg3.jpg](./seqImg3.jpg/)| Sequential Image   | 
|   11  | [seqImg4.jpg](./seqImg4.jpg/)| Sequential Image   | 
|   12  | [seqImg5.jpg](./seqImg5.jpg/)| Sequential Image   | 

### Speed Conclusion
- Concurrent download took: 293.4561ms
- Sequential download took: 2.3523046s

### Example Commands
- Only command needed is to use `go run main.go`
```
                 //Need to cd into the correct folder first look for correct folder. 
$ ls             //this will list the folders where you are located  //to go into the according folder 
$ cd /y/4143-PLC-Hughes/assignments/p03/ImageMod  //brings you to the folder with the program we need to run.
```