# goDownloader
godownloader is a CLI written in `go` where you can download file from the web concurrently.

## Installation
Install go. You can use this [link](https://go.dev/dl/) for installation.

Clone this repository and you are good to `go`.

## Usage
Compile the code using the following command
```
go build -o godownloader
```

Create a text file that contains the address and output directory seperated by spaces in separate lines.
```
https://images.financemagnates.com/images/DogeCoin%20header-min_id_9b93abb9-4711-4f3b-9b9f-666b164647de_size900.jpg ./doge.jpg
https://cloudfront-us-east-1.images.arcpublishing.com/coindesk/L4V344Y2EBFTBBNIEN7WUYUIBY.jpg ./original_doge.jpg

```

Then run the binary
```
./godownloader -file=/path/to/file.txt
```

Your files will be downloaded in the location given in the text file

Additional flags
```
./godownloader -file=/path/ -concurrent=2 -log=true
```
- `concurrent` - number of concurrent threads to use (default value is 10)
- `log` - prints log if set true (default is false)
