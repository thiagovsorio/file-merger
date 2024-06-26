# File Merger
This project does take two files that you pass and merge then into one destination file, copying line by line from the two sources, alternatively.

### What it does
```mermaid
flowchart LR
    f1("`
    FileOne.txt
      line 1
      line 2
      line 3
    `") -.-> File-merger
    f2("`
    FileTwo.txt
      linha 1
      linha 2
      linha 3
    `") -.-> File-merger
    File-merger -.-> 
    d("`
    DestionationFile.txt
      line 1
      linha 1
      line 2
      linha 2
      line 3
      linha 3
    `")
```

### Purpose

The initial purpose was my learning of Latin. So I wanted to take two .tsv files of the Bible, one in English and one in Latin and merge them line by line.

Then I used [this repo](https://github.com/LukeSmithxyz/kjv/tree/master) that reads tsv files, and used it for my Latin/kjv one.

## Usage (Linux and Docker)
With the two input files in this directory:  
Build and run:
```sh
# Build
make 

# Run
FILE_ONE=yourFileOne.txt \
FILE_TWO=yourFileTwo.txt \
FILE_DEST=destinationFile.txt \
make run
```

## Usage (Linux and Go)
It compiles in Go, you can fastly install it [here](https://go.dev/doc/install).

Build and run:
```sh
# Build
make install

# Run
./file-merger \
  -source1=yourFileOne.txt \
  -source2=yourFileTwo.txt \
  -destination=yourDestinationFile.txt
```

