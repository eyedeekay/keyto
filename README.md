keyto
=====

Simple application for getting information from plain-text i2p destination keys.
It's goal is to be short, understandable to anyone, and just do this one thing
in a reliable, easy to install way that doesn't even require an I2P router
installed to use.

setup
-----

If you have go and git installed, you can just

        go get -u github.com/eyedeekay/keyto

Or, you can use the binary by downloading and unzipping this file:
[keyto](https://github.com/eyedeekay/keyto/archive/master.zip), then running the
appropriate application for your OS and Platform. To install it into your PATH,
you can use:

        make

to build it for your platform, and

        sudo make install

to install it to /usr/bin/

use
---

The first argument to keyto should be a full destination, in plain text format.
If you've got a file like local_dest.key from i2p-bote, you can redirect the
file into the first argument like so:

        keyto $(cat local_dest.key)
