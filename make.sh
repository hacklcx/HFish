xgo --targets=darwin/amd64 -ldflags='-w -s' .
xgo --targets=darwin/386 -ldflags='-w -s' .
xgo --targets=linux/arm64 -ldflags='-w -s' .
xgo --targets=linux/arm -ldflags='-w -s' .
xgo --targets=linux/amd64 -ldflags='-w -s' .
xgo --targets=linux/386 -ldflags='-w -s' .
xgo --targets=windows/amd64 -ldflags='-w -s' . 
xgo --targets=windows/386 -ldflags='-w -s' .
