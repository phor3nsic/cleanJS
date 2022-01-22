## CleanJS

The cleanJS has the function of analyzing urls of javascript files provided via stdin, checking if a url is valid and if the .map file is visible, with that it returns the url to the user, to be integrated in the next step!

### Example

```
echo http://google.com/ | getJS --complete | cleanJS | sort -u
```