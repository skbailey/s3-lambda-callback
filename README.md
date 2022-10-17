## S3 Lambda Callback

Prepare the [go binary for deployment](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html).

```bash
GOOS=linux CGO_ENABLED=0 go build main.go
```

Then zip it for upload
```bash
zip function.zip main
```

Upload the package
```bash
aws lambda update-function-code --function-name uploadCallback --zip-file fileb://function.zip
```
