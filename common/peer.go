package common

import (
	"fmt"
	"net"
	"encoding/json"
	//"strings"
	//"strconv"
	"errors"
)

type Peer struct {
	Address string `json:"Address"`
	Port uint16 `json:Port`
}

var peers map[string][]*Peer

func ReturnLocalInterfacesAsPeers(port uint16)([] *Peer){
	addrs,err := net.InterfaceAddrs()
	var rtr []*Peer
	if err != nil {
		panic(err)
	}
	for _,addr := range(addrs) {
		var p Peer
		p.Address = addr.String()
		fmt.Println(addr.String())
		p.Port = port
		rtr = append(rtr, &p)

	}
	return rtr
}

func InitializePeerMap(infohash string, port uint16) {
	peers = make(map[string][]*Peer)
	local_peers := ReturnLocalInterfacesAsPeers(2655)
	for _,p := range(local_peers) {
		fmt.Println(p)
		peers[infohash] = append(peers[infohash],p)
	}
}

func PeerMapMarshalJSON()([]byte,error){
	jsonbytes,err := json.Marshal(peers)
	if err != nil {
		return nil,err
	}
	return jsonbytes,nil
}

func PeersPresentForInfohash(ihash string)(bool) {
	if peers[ihash] != nil {
		return true
	}
	return false
}

func GetPeerlistAtInfohash( ihash string)([]*Peer){
	if peers[ihash] != nil {
		return peers[ihash]
	}
	return nil
}
func PeerMapUnmarshalJSON(jsbytes []byte, infohash string)(error,[] *Peer ) {
	var jsonpeers map[string]interface{}
	var peerlist []*Peer
	err := json.Unmarshal(jsbytes,&jsonpeers)
	if err != nil {
		return errors.New("Unable to unmarshal json to peer map")
	}
	for i,_ := range(jsonpeers) {
		plist,ok := jsonpeers[i].([]interface{})
		if ok {
			for _,v := range(plist) {
				var tmpPeer Peer
				p,ok2 := v.(map[string]interface{})
				//fmt.Println(v)
				if ok2 {
					
					port,portok := p["Port"].(float64)
					if portok {
						tmpPeer.Port = uint16(port)
					}
					peer_addr_if,peer_addr_ok := p["Address"].(string)
					if peer_addr_ok {
						tmpPeer.Address = peer_addr_if
					}
				}

				append(peerlist,&tmpPeer)
			}

		} else {
			return errors.New("Failure to umarshal peer struct list")
		}
	}
	return nil,peerlist
}