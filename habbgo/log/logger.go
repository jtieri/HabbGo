package log

import (
	"github.com/jtieri/HabbGo/habbgo/game/player"
	"github.com/jtieri/HabbGo/habbgo/protocol/packets"
	"log"
	"reflect"
	"runtime"
	"strings"
)

func PrintOutgoingPacket(playerAddr string, p *packets.OutgoingPacket) {
	log.Printf("[%v] [OUTGOING] [%v - %v]: %v ", playerAddr, p.Header, p.HeaderId, p.Payload.String())
}

func PrintIncomingPacket(playerAddr string, handler func(*player.Player, *packets.IncomingPacket), p *packets.IncomingPacket) {
	hName := getHandlerName(runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name())
	log.Printf("[%v] [INCOMING] [%v - %v|%v]: %v ", playerAddr, hName, p.Header, p.HeaderId, p.Payload.String())
}

func PrintUnknownPacket(playerAddr string, p *packets.IncomingPacket) {
	log.Printf("[%v] [UNK] [%v - %v]: %v ", playerAddr, p.Header, p.HeaderId, p.Payload.String())
}

func getHandlerName(handler string) string {
	sp := strings.Split(handler, "/") // e.g. github.com/jtieri/HabbGo/habbgo/protocol/handlers.GenerateKey
	s2 := sp[len(sp)-1]               // e.g. handlers.GenerateKey
	return strings.Split(s2, ".")[1]  // e.g. GenerateKey
}
