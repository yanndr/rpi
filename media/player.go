package media

import (
	"log"
	"os/exec"
	"sync"
)

type Player interface {
	Play(file string)
}

type OmxPlayer struct {
	paying bool
	mutex  sync.Mutex
}

func (p *OmxPlayer) Play(file string) {
	if !p.paying {
		p.mutex.Lock()
		defer p.mutex.Unlock()
		p.paying = true
		cmd := exec.Command("mplayer", file)
		err := cmd.Start()
		if err != nil {
			log.Println(err)
		}
		log.Printf("Waiting for command to finish...")
		err = cmd.Wait()
		log.Printf("Command finished with error: %v", err)
		p.paying = false
	}
}
