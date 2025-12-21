# Development Environment

## Go (latest version)

`sudo pacman go`

### Go Packages

- `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
- `go install github.com/aarondl/sqlboiler/v4@latest`
- `go install github.com/aarondl/sqlboiler/v4/drivers/sqlboiler-psql@latest`
- `go install github.com/aarondl/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest`

## Build Tools

- go-task: `sudo pacman -S go-task`
- npm: `sudo pacman -S npm`
- gcc: `sudo pacman -S gcc`
- libwebkit: `sudo pacman -S webkit2gtk`
- pkg-config: `sudo pacman -S pkg-config`
- docker: `sudo pacman -S docker`
- npm cli: `sudo npm install -g @go-task/cli`

## Git (GitHub)

- `git config --global user.name`
- `git config --global user.email`

Store the GitHub token in the `.env` file, in the root project (Station-Manager):

    GITHUB_TOKEN=my_github_token

# Serial Port

If you receive an error regarding serial port permissions, it is likely that you need need to
install the USB drivers for your transceiver. In the case of Yaesu, this will likely be the
**CP210x USB to UART Bridge VCP Driver**, which, for Linux is maintained in the kernel tree, so
simply switching on your Yaesu rig should suffice. If you are not ready for this, disable
the serial port in the `config.json` file. Remember, your PC should be on and you should be logged in
**before** connecting the Yaesu rig.

If you are still getting *permission errors* you will need to modify you system to allow access:

## Troubleshooting

From a console enter the following command:

     ls -al /dev/ttyUSB0

This should output something like this:

    crw-rw---- 1 root uucp 188, 0 Dec 11 14:41 /dev/ttyUSB0

The group `uucp` is the group that owns the device. You need to add your user to this group.
To do this, enter the following command:

     sudo usermod -a -G uucp $USER

The enter the following command to verify that you are now a member of the `uucp` group:

     id <your_username>

You should see `uucp` listed in the groups. If not, you may need to log out and log back in
for the changes to take effect.
