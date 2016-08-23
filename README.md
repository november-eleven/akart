# Akart

An auto-complete pet project in order to play with an **adaptive radix tree**.

## Requirements

 * **go v1.6.2** _([Installation instructions](https://golang.org/doc/install))_

## Install source on local filesystem

```bash
cd $GOPATH
mkdir -p src/github.com/november-eleven/
cd src/github.com/november-eleven/
git clone git@github.com:november-eleven/akart.git
cd akart
```


## Compile

```bash
go build
```

## Testing

```bash
go test
```

## Launch

```bash
./akart prog
```

## Dependencies

**Akart** use a fork of [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree), you can find it
in the `art/` folder.

## Evolutions ?

 * If the list of keywords begins to be large (several millions), we could use a solution like
 [`mmap`](https://en.wikipedia.org/wiki/Mmap).

 * If the auto-complete engine must match any portion of the keywords, we should use an adaptive suffix tree instead.

