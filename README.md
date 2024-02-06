# Bitwarden CLI Helper

A simple CLI to support a local workflow with the offical Bitwarden CLI. As Bitwarden Desktop is not available (yet) on Linux ARM64 systems, I have decided to use the Bitwarden CLI for password management on my local machine (especially for non-browser based logins). The official Bitwarden CLI however misses some fundamental features in my opinion:

* Automatic setting of the BW_SESSION environment variable. Each time the Vault is unlocked, additional input is required to either set the BW_SESSION token in your environment, or copy the token to your clipboard. This helper automatically stores the token which can be used for further reference.

* Automatic copy to clipboard of login items. Once a login item is retrieved via the vault, you need to manually copy the output for use. This helper automatically copies any login item to your clipboard.

## Overview

In order to unlock the Vault and retrieve a session token, the CLI makes use of Bitwarden's [Vault Management API](https://bitwarden.com/help/vault-management-api/). The Vault Management API can be accessed by running `bw serve` on your local machine. By default this spins up on `http://localhost:8087`. During the initialization of BW Helper you are able to change these values.

BW Helper will automatically run `bw serve` on the background when trying to unlock the vault, so there's no need to manually run this command.

## Requirements

This CLI makes use of the [clipboard](https://github.com/atotto/clipboard) package. Therefore, on Linux it requires either the `xclip` or `xsel` package to be installed.

For all operating systems it is required to have the [Bitwarden CLI](https://bitwarden.com/help/cli/) installed and configured (e.g. the correct server is setup etc.)

## Install

To install, run:

`go install github.com/frankleef/bw-helper/cmd/bw-helper@latest`

## Usage

This CLI assumes you have already set your server and are logged in using the offical Bitwarden CLI.

### Initialization

This command creates the necessary configuration folders and files, allowing you to unlock your vault and retrieve logins.

`bw-helper init --password <my-password>`

**Note**
* The password field is required
* Currently, the password is stored as plain text in `$HOME/.bw-helper/config.yaml`.

```
NAME:
   bw-helper init

USAGE:
   bw-helper init [command [command options]] 

DESCRIPTION:
   Initialize the helper.

OPTIONS:
   --password value  Password to login into Bitwarden
   --host value      Host of Vault Management API. Default http://localhost
   --port value      Vault Management API port. Default 8087
   --help, -h        show help (default: false)
```

### Logging in

To unlock your vault, run `bw-helper login`. This will perform the following actions:

* Call the Vault Management API with your provided password to retrieve a session token
* Stores this token in `$HOME/.bw-helper/.token`

```
NAME:
   bw-helper login

USAGE:
   bw-helper login [command [command options]] 

DESCRIPTION:
   Login to your Bitwarden Vault

OPTIONS:
   --help, -h  show help (default: false)

```

### Getting items from your Vault

Since it is not possible to store environment variables in a parent process, nor does Bitwarden CLI support reading the token from file, we need to use `bw-helper` as a wrapper for using the Bitwarden CLI. The exact same command structure supported by the Bitwarden CLI is supported by `bw-helper`.

***Example: retrieving your LinkedIn password***

`bw-helper get password linkedin`

***Example: getting your LinkedIn username***

`bw-helper get username linkedin`

**Note**

**Any** command supported by the Bitwarden CLI is suppoerted by `bw-helper`, as the command arguments are passed directly to Bitwarden CLI if it's an unknown command for `bw-helper`. As an example, calling `bw-helper gibberish` will end up calling `bw gibberish`.