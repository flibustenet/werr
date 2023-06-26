Disclaimer: private use and with my coworkers. It will break! feel free to copy-paste.


# werr: opt-in traces in errors.

Add the name of the function/method and the file:line opt-in on each error
to print a minimalist traceback with decorations.

Example:

```
> gotodo/vue.Issue() issue.go:29
read issue:
> gotodo/model.(*Project).GetIssueStr() issue.go:99
Erreur en lecture issue x: strconv.Atoi: parsing "x": invalid syntax
```

## Usage

`go get go.flibuste.net/werr`

Wrap or create the error with one of the werr functions, it'll add the
function/method name and the file:line in the string. Nothing more.

`Wrap` and `Wrapf` will wrap the error, like `%w`.
`Errorf` is like `fmt.Errorf` with trace
`New` is like `errors.New` with trace

Just print the error.

## How to adapt a legacy app

Replace `fmt.Errorf` with `werr.Errorf`
```
sed -i 's/fmt.Errorf/werr.Errorf/g' *go
```

Replace `errors.New` with `werr.New`
```
sed -i 's/errors.New/werr.New/g' *go
```

Add trace + wrap to return err without decoration, like `%w`
```
sed -i 's/return err/return werr.Wrap(err)/g' *go
```

Add `import go.flibuste.net/werr`
```
goimports -w *.go
```

Find return err still not decorated
`ack 'return.*err$'
