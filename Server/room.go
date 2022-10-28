package main

import "net"

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {
	for addr, m := range r.members {
		if addr != sender.conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}

func (r *room) membersList(sender *client) []string {
	var list []string
	for addr, m := range r.members {
		if addr != sender.conn.RemoteAddr() {
			list = append(list, m.name)
		}
	}
	return list
}
