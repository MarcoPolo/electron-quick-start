# Net Fail Windows

Reproducing windows naped pipe connection errors from renderers

## To Use

To clone and run this repository you'll need [Git](https://git-scm.com) and [Node.js](https://nodejs.org/en/download/) (which comes with [npm](http://npmjs.com)), and [golang](https://golang.org/dl) installed on your computer. From your command line:


```bash
# Clone this repository
git clone -b net-fail-windows https://github.com/marcopolo/electron-quick-start net-fail-windows

# Go into the repository
cd net-fail-windows

# Setup go listening server
cd npipe


# get dependencies
go get
go build
./npipe.exe --server # starts the server

# Install dependencies and run the app
npm install && npm start
```

You'll see the main process says `connected` in the terminal, but the renderer hangs on `starting connection`.
The renderer will hang until you run `windowsHack()`. Why that works, I have no idea


#### License [CC0 (Public Domain)](LICENSE.md)
