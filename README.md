# werr: opt-in traces in errors.

Some error helper for personal use and with my coworkers.

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
