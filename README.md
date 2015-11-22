envoy
============

Make a hosts file from Chef Server nodes.

Example command:

`envoy create -f output.txt -n node-name -s https://chef-server.example.com -e production -k key.pem`

This replaces a Ruby script that did this pretty well until we hit about 900 nodes. After that it used way too much memory and became a problem.

Binaries are available on the [releases page](https://github.com/darron/envoy/releases).

Also available as a [Docker image](https://hub.docker.com/r/darron/envoy/).
