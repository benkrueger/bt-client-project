# ltshare
Peer 2 Peer filesharing program based on homomorphic hashing.

Plans:

- Metadata files are in json 
- tracker implementation will be a simple go http server that listens on an arbitrary tcp port. It will be barebones, no db, no reverse proxy, just keep track of torrents and peers. 

- No directory support, only support single file sharing
- Figure out a protocol based on sharing lthash signature distance to share a file.
