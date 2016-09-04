Small example of using https://github.com/GeertJohan/go.rice

Things to note about go.rice:

- Using the `append` workflow allows you to distribute a binary and
  have the user of that binary append their own html/css/js files
- Love the way that "append" scans the binary and only appends the
  directories which are actually referenced.
- `append` is efficient for compiling/serving a large set of files ..
  the files are zipped thereby not bloating the binary unnecessarily
- default `index.html` will be appended to a directory name, so with this
  example you can just load http://localhost:8080
- you can extract the assets from an "appended" binary just using `unzip` (although,
  it seemed to create files instead of directories .. seems to work if you pre-create
  the directories and don't allow unzip to replace the dirs with files)

To use this example:

- go get / git clone
- glide install
- make