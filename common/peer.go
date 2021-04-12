package common

import (
	"fmt"
	"net"
	"encoding/json"
	//"strings"
	//"strconv"
	"errors"
)
const PB_SIZE = 6 
type Peer struct {
	Address net.Addr `json:"Address"`
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
		p.Address = addr
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
func PeerMapUnmarshalJSON(jsbytes []byte, infohash string)(error) {
	var jsonpeers map[string]interface{}
	json.Unmarshal(jsbytes,&jsonpeers)
	if jsonpeers == nil {
		return errors.New("Unable to unmarshal json to peer map")
	}
	for i,_ := range(jsonpeers) {
		plist,ok := jsonpeers[i].([]interface{})
		if ok {
			for _,v := range(plist) {
				fmt.Println(v)
			}
		} else {
			return errors.New("Failure to umarshal peer struct list")
		}
	}
	return nil
}