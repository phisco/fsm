# fsm
A simple FSM Go library 

This is intentionally as simple as possible, more complex implementations can be found for example [here](https://github.com/looplab/fsm).

I will probably add more FSM implementations in the near future, always focusing on simplicity.

# How to use

Some simple test cases can be found [here](fsm_test.go), here just a basic example:

```go
import (
	"log"
	"testing"
	"github.com/stretchr/testify/assert"
)

type On struct {}
type Off struct {}

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
	assert.IsType(t,new(Off), fsm.getCurrentState())
	err = fsm.tic()
	assert.Nil(t, err)
	assert.IsType(t,new(On), fsm.getCurrentState())
}
```