# go-learn
Learning golang

### command for git
git pull origin master --allow-unrelated-histories
git merge origin origin/master


## learn oop 
```go
type deck []string

func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}
```
