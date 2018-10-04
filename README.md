# go-file-ordering-bug

Behavior for Go 1.11 and later seems dependent on file name ordering:

```
chris:~/go/src/github.com/ccbrown/go-file-ordering-bug% go version
go version go1.11.1 darwin/amd64

chris:~/go/src/github.com/ccbrown/go-file-ordering-bug% go run main.go
func (github.com/ccbrown/go-file-ordering-bug/thepackage.A).A()
func (github.com/ccbrown/go-file-ordering-bug/thepackage.B).B()
func (github.com/ccbrown/go-file-ordering-bug/thepackage.SuperInterface).D()
func (github.com/ccbrown/go-file-ordering-bug/thepackage.SuperInterface).E()

chris:~/go/src/github.com/ccbrown/go-file-ordering-bug% go version    
go version go1.10.4 darwin/amd64

chris:~/go/src/github.com/ccbrown/go-file-ordering-bug% go run main.go
func (github.com/ccbrown/go-file-ordering-bug/thepackage.A).A()
func (github.com/ccbrown/go-file-ordering-bug/thepackage.B).B()
func (github.com/ccbrown/go-file-ordering-bug/thepackage.D).D()
func (github.com/ccbrown/go-file-ordering-bug/thepackage.E).E()
```
