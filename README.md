# Union
Union is a program created to automate campo solar job.

## Screenshots:
![photo_2021-11-14_00-21-48](https://user-images.githubusercontent.com/69154350/141662907-bb48e4ac-97af-413c-95bc-be8845505550.jpg)


## Building - Using docker:
* First install docker.
* ```sudo docker run -it fedora /bin/bash```
* `yum install mingw64-gtk3 go glib2-devel`
* `bash -c "sed -i -e 's/-Wl,-luuid/-luuid/g' /usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig/gdk-3.0.pc"`
* Setup git in container to clone, this repo.
* `cd root`
* `mkdir go`
* `cd go`
* `mkdir src pkg lib bin`
* `cd src`
* Get program from github, and CD to it.
* `PKG_CONFIG_PATH=/usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go install -v github.com/gotk3/gotk3/gtk` #This will take about 8 minutes. 
* `CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui` #Compile
* `yes | cp -r /usr/x86_64-w64-mingw32/sys-root/mingw/*` . #Get gtk libs
* `sudo docker ps -alq` #get id from current session image
* `cp <image-id>:/root/go/src/Union/union` Documentos/union


