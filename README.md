envoy
============

Make a hosts file from Chef Server nodes.

Example command:

`envoy create -f output.txt -n node-name -s https://chef-server.example.com -e production -k key.pem`

Replacing a Ruby script that did this well until we hit about 1000 nodes and then used up way too much memory.
