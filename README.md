# mdir

move files to md5-style path

## install

```
$ go get github.com/turnon/mdir
```

## usage

```go
cmd := mdir.Cmd{
    Src:      "/path/to/src",
    Dest:     "/path/to/dest",
    Segments: []int{2, 2, 2},
    Force:    true, // actually move files
    CopyFile: true, // copy instead move
}

cmd.MvFiles()
```