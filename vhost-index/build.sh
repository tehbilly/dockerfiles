FILENAME=vhost-index

[ -e "$FILENAME" ] && rm $FILENAME

echo "Building ${FILENAME}"
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o $FILENAME
#echo "Compressing (using gzexe)"
#gzexe $FILENAME
#rm $FILENAME~
