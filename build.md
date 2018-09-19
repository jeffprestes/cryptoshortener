# Build for Linux server

```bash
env GOOS=linux GOARCH=amd64 go build -o cryptoshortener 
```

# To copy the new binary to the server 

```bash  
scp cryptoshortener <<user>>@xxxxxxx:/srv/cryptoshortener
```

```bash
scp public/templates/index.html root@xxxxxxxx:/srv/cryptoshortener/public/templates/index.html
scp -r public root@xxxxxxx:/srv/cryptoshortener
```

