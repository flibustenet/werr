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

Wrap or create the error with one of the werr functions, it'll add the
function/method name and the file:line in the string. Nothing more.

Just print the error.

## How to adapt a legacy app

Replace `fmt.Errorf` with `werr.Errorf`
```
sed -i 's/fmt.Errorf/werr.Errorf/g' *go
gofmt -w *.go
```

Replace `errors.New` with `werr.New`
```
sed -i 's/errors.New/werr.New/g' *go
gofmt -w *.go
```

Add trace to return err without decoration
```
sed -i 's/return err/return werr.Trace(err)/g' *go
gofmt -w *.go
```

Find return err still not decorated
`ack 'return.*err$'
