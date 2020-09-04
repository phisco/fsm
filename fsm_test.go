package fsm

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type On struct{}
type Off struct{}

func (o Off) next() (nextState State, err error) {
	log.Println("Off -> On")
	return &On{}, nil
}

func (o On) next() (nextState State, err error) {
	log.Println("On -> Off")
	return &Off{}, nil
}

func TestNewFsm(t *testing.T) {
	fsm := NewFSM(On{})
	err := fsm.tic()
	assert.Nil(t, err)
	assert.IsType(t, new(Off), fsm.getCurrentState())
	err = fsm.tic()
	assert.Nil(t, err)
	assert.IsType(t, new(On), fsm.getCurrentState())
}

type Open struct {
}
type Closed struct {
	coins int
}

func (c Closed) next() (nextState State, err error) {
	c.coins += 1
	if c.coins < 2 {
		log.Printf("Closed -> Closed (%d)\n", c.coins)
		return c, nil
	}
	log.Printf("Closed -> Open (%d -> 0)\n", c.coins)
	return Open{}, nil
}

func (o Open) next() (nextState State, err error) {
	log.Printf("Open -> Closed (%d)\n", 0)
	return Closed{}, nil
}

func TestNewTurnstile(t *testing.T) {
	fsm := NewFSM(Closed{0})
	err := fsm.tic()
	assert.Nil(t, err)
	assert.IsType(t, Closed{}, fsm.getCurrentState())
	err = fsm.tic()
	assert.Nil(t, err)
	assert.IsType(t, Open{}, fsm.getCurrentState())
	err = fsm.tic()
	assert.Nil(t, err)
	assert.IsType(t, Closed{}, fsm.getCurrentState())
	err = fsm.tic()
	assert.Nil(t, err)
	assert.IsType(t, Closed{}, fsm.getCurrentState())
	assert.Equal(t, 1, fsm.getCurrentState().(Closed).coins)
}
