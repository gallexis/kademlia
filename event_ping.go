package main

import (
	log "github.com/sirupsen/logrus"
	ds "kademlia/datastructure"
	"kademlia/message"
	"net"
)

// PING
func (d *DHT) onPingResponse(node ds.Node, ping *message.PingResponse, addr net.UDPAddr) {
	log.Info("onPingResponse", addr)
	d.routingTable.UpdateNodeStatus(ping.Id)
	d.PingPool <- node
}

func (d *DHT) onPingRequest(msg *message.PingRequest, addr net.UDPAddr) {
	log.Info("onPingRequest")
	pingResponse := message.PingResponse{
		T:  msg.T,
		Id: d.selfNodeID,
	}

	d.routingTable.UpdateNodeStatus(msg.Id)

	if _, err := d.conn.WriteToUDP(pingResponse.Encode(), &addr); err != nil {
		log.Error("Failed to send ping response")
	}
}

func (d *DHT) sendPingRequest(node ds.Node, tx message.TransactionId) {
	node.Send(d.conn, message.PingRequest{
		T:  tx,
		Id: d.selfNodeID,
	}.Encode())
}
