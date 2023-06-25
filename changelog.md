# CHANGELOG

## v0.5.0 BREAKING CHANGES
- very simple, add only method + file:line, not more all the stack


Add trace to return err without decoration
```
sed -i 's/return err/return werr.Trace(err)/g' *go
gofmt -w *.go
```

Find return err still not decorated
`ack 'return.*err$'


## v0.0.8
- remplace Stack par Wrap
- ajout MustWrap et MustWrapf (wrap et panic si err != nil)
