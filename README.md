resize
======

a simple library to call opencv to resize the image
### Install
* install `opencv`

  ```
  Mac:
  	cd <opencv directory>
  	mkdir release
  	cd release
  	cmake -G "Unix Makefiles" -D BUILD_opencv_python=OFF ..
  	make -j9
  	sudo make install
  	
  Centos/Redhat:
    yum install opencv-devel
  ``` 

* `go get github.com/Centny/resize`

### Usage

```
err = Resize("t.png", "s.png", 100, 100)
err = Resize("t.jpg", "s.jpg", 100, 100)
```
