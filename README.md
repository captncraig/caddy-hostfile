# caddy-hostfile
Local hostnames for caddy

Usage is simple, just add `hostfile` to your server block, and caddy will add the host(s) to your local hostfile on startup:

```
myapp.local, app.dev {
  hostfile
}
```

It will also remove the entries on a clean shutdown.
