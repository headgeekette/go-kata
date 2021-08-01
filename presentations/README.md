# The Presentations

These files can be read as is. But you can use pandoc to convert it to a [Reveal.js](https://revealjs.com/) file. 

I also have a pandoc Docker container that you can use.

```
docker run --rm -v `pwd`:/source headgeekette/pandoc:2.14.1 -t revealjs -s Functions.md -o Functions.html -V theme=moon
```