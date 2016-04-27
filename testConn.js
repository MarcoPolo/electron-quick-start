const net = require('net')
const getenv = require('getenv')
const path = require('path')

function buildWin32SocketRoot () {
  let appdata = getenv('APPDATA', '')
  // Remove leading drive letter e.g. C:
  if (/^[a-zA-Z]:/.test(appdata)) {
    appdata = appdata.slice(2)
  }

  const runMode = getenv('KEYBASE_RUN_MODE', 'prod')

  // Handle runModes, prod has no extension.
  let extension = ''

  if (runMode !== 'prod') {
    extension = runMode.charAt(0).toUpperCase() + runMode.substr(1)
  }
  let path = `\\\\.\\pipe\\kbservice${appdata}\\Keybase${extension}`
  return path
}

function testConn () {
  console.log('starting connection')
  // const p = '\\\\.\\pipe\\kbservice\\Users\\IEUser\\AppData\\Roaming\\Keybase\\keybased.sock'
  const socketRoot = buildWin32SocketRoot()
  const socketName = 'keybased.sock'
  const socketPath = path.join(socketRoot, socketName)

  const s = net.connect({path: socketPath}, () => { console.log('connected') })

  s.on('end', () => {
    console.log('disconnected from server')
  }).on('error', (e) => {
    console.log('error from server', e)
  }).on('lookup', (l) => {
    console.log('lookup: ', l)
  })
}

module.exports = testConn
