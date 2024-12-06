package service

import (
	"github.com/jtieri/habbgo/legacy/game/item"
	"github.com/jtieri/habbgo/legacy/game/navigator"
	"github.com/jtieri/habbgo/legacy/game/player"
	"github.com/jtieri/habbgo/legacy/game/room"
	"github.com/jtieri/habbgo/legacy/game/types"
)

type Proxies struct {
	Rooms     *room.RoomServiceProxy
	Items     *item.ItemServiceProxy
	Navigator *navigator.NavigatorServiceProxy
	Players   *player.PlayerServiceProxy
}

func (p *Proxies) RoomService() types.Proxy {
	return p.Rooms
}
func (p *Proxies) ItemService() types.Proxy {
	return p.Items
}

func (p *Proxies) NavigatorService() types.Proxy {
	return p.Navigator
}

func (p *Proxies) PlayerService() types.Proxy {
	return p.Players
}
