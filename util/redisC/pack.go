package redisC

type RedigoPack struct {
	String stringRds
	Hash   hashRds
	Key    keyRds
	List   listRds
	Set    setRds
	ZSet   zSetRds
	Bit    bitRds
}

var Conn = new(RedigoPack)
