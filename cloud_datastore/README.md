
### Updating Indexes

```bash
$ gcloud config set project teraconnect-development
```

```bash
$ gcloud datastore indexes create index.yaml
```

### Deleting unused Indexes

```bash
$ gcloud datastore indexes cleanup index.yaml
```
