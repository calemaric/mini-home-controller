

const { app, BrowserWindow } = require('electron')

const createWindow = () => {
        const win = new BrowserWindow({
                width: 1200,
                height: 800,
                acceptFirstMouse: true,
                autoHideMenuBar: true,
                useContentSize: true,
        })

        win.loadURL('http://localhost:8080');
}

app.whenReady().then(() => {
        createWindow()
})

// Quit when all windows are closed.
app.on('window-all-closed', function () {
        // On OS X it is common for applications and their menu bar
        // to stay active until the user quits explicitly with Cmd + Q
        if (process.platform !== 'darwin') {
                app.quit()
        }
})

app.on('activate', function () {
        // On OS X it's common to re-create a window in the app when the
        // dock icon is clicked and there are no other windows open.
        if (mainWindow === null) {
                createWindow()
        }
})

