# bt-client-project
bittorrent client implementation in c++ as part of the 4730 project

```
bt-client <metadatafile> <download directory>
bt-tracker <port>
```

## Update: how we will vastly simplify this project

So, for this project we have to implement a bittorrent *like* protocol, not actual bittorrent. So, I have taken the creative liberty of
coming up with a much simpler p2p protocol, *inspired* by bt v1 but with some key differences. Looking at both the scope of the bt protocol and the time we have to complete, I've decided to *vastly* simplifiy the spec. 

Plans:

- Metadata files are now in json instead of bencode. 
- tracker implementation will be a simple go http server that listens on an arbitrary tcp port. It will be barebones, no db, no reverse proxy, just keep track of torrents and peers. 

- No directory support, only support single file sharing