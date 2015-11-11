## Imageproxy

This is a repository to host [imageproxy](https://github.com/willnorris/imageproxy) on Heroku. The setup of this is based on [prototype-imageproxy](https://github.com/oreillymedia/prototype-imageproxy) but introduces a few changes. Since Go 1.5 it's no longer possible to import a program. This doesn't play well with Heroku since it needs all dependencies in the repo.

### Usage

```shell
godep save
git add .; git commit -m 'updated dependencies'
```

### Deployment
Depending on your git origin setup you'll have to adjust the origin name before pushing the changes to Heroku.

```shell
git push heroku master
```
Make sure to set the following environment variables

#### `WHITELIST`
A comma-separated list of domains that are allowed image sources.

#### `BASEURL` (optional)
The domain to be used if only a relative path to the image is provided, such as:
```
http://imageproxy.example.com/400x400/image.jpg
```

If `BASEURL` is set to `http://example.com/static/`, `imageproxy` will
automatically expand the `image.jpg` path to the full URL `http://example.com/static/image.jpg`.
