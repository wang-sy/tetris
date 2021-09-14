package keyboard

import (
	"sync"

	"github.com/eiannone/keyboard"
	"github.com/pkg/errors"
)

const bufferSize = 10

type KeyboardEventBuffer struct {
	mu  sync.Mutex
	buf []keyboard.KeyEvent

	eventChan <-chan keyboard.KeyEvent
}

// NewKeyboardEventBuffer create KeyboardEventBuffer and init.
func NewKeyboardEventBuffer() (*KeyboardEventBuffer, error) {
	eventChan, err := keyboard.GetKeys(bufferSize)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	keyboardEventBuffer := &KeyboardEventBuffer{
		buf:       make([]keyboard.KeyEvent, 0),
		eventChan: eventChan,
	}

	go keyboardEventBuffer.watch()

	return keyboardEventBuffer, nil
}

// ListAndClear buffer.
func (b *KeyboardEventBuffer) ListAndClear() []keyboard.KeyEvent {
	b.mu.Lock()
	defer b.mu.Unlock()

	bufCopy := make([]keyboard.KeyEvent, len(b.buf))
	copy(bufCopy, b.buf)

	b.buf = make([]keyboard.KeyEvent, 0)

	return bufCopy
}

// watch event chan.
func (b *KeyboardEventBuffer) watch() {
	for {
		event := <-b.eventChan

		b.mu.Lock()
		b.buf = append(b.buf, event)
		b.mu.Unlock()
	}
}
