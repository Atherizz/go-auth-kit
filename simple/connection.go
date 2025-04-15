package simple

import "fmt"

type Connection struct {
	*File
}

func NewConnection (file *File) (*Connection, func()) {
	conn := &Connection{
		File: file,
	}
	return conn, func() {
		conn.Close()
	}
}

func (c *Connection) Close() {
	fmt.Println("close connection", c.File.Name)
}