## Imageproxy

This is a repository to host the [imageproxy](https://github.com/willnorris/imageproxy) on Heroku. The setup of this is based on the [prototype-imageproxy](https://github.com/oreillymedia/prototype-imageproxy) but introduces a few changes. Since Go 1.5 it's no longer possible to import a program. This doesn't play well with Heroku since it needs all dependencies in the repo.

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
