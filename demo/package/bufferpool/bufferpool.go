package bufferpool

// a custom  buffer pool package.
import (
	"errors"
	"math"
	"sync"
)

const (
	// the default size of buffer pool --- 4K.
	DefaultBufferSize = 2 << 11
	// the default timeout when write something to buffer pool.
	DefaultWriteTime = 5
)

var (
	COPYERROR       = errors.New("copy error!")
	TIMEOUTERROR    = errors.New("write timeout!")
	BUFFERNOTENOUGH = errors.New("need more free buffer!")
	TYPEERROR       = errors.New("put need string or []byte param!")
)

type BufferPool struct {
	Size          int64
	Data          []byte
	Used          int64
	Free          int64
	AutoIncrement bool
	sync.Mutex
}

// AutoIncrement means the buffer pool's max size is variable.
// it will increment buffer size when the free size is not enough.
// the additional buffer size equal to source buffer pool.
// new buffer pool is 2 times than before.
func NewBufferPool() *BufferPool {
	return &BufferPool{
		Size:          DefaultBufferSize,
		AutoIncrement: true,
		Free:          DefaultBufferSize,
		Data:          make([]byte, DefaultBufferSize),
	}
}

// disable buffer pool auto increment..
func (b *BufferPool) DisAbleAutoIncrement() {
	b.Lock()
	b.AutoIncrement = false
	b.Unlock()
}

// put data to buffer pool.
// if the free buffer size is smaller than the give data,
// and  will return BUFFERNOTENOUGH.
// if timeout when writing to the b, will return TIMEOUTERROR.
// if success, will return n (about the write size) and nil.
func (b *BufferPool) Put(data interface{}) (n int64, err error) {
	var putData []byte
	var putDataLength int64
	var blackspaceLength int64
	switch info := data.(type) {
	case string:
		putData = []byte(info)
		putDataLength = int64(len(putData))
	case []byte:
		putData = info
		putDataLength = int64(len(putData))
	default:
		return 0, TYPEERROR
	}

	b.Lock()
	// free buffer size smaller than data size which will be written to.
	if b.Free < int64(len(putData)) {
		addRate := math.Ceil(float64(putDataLength) / float64(b.Size))
		if addRate <= 1 {
			addRate = 2
		}
		if b.AutoIncrement == true {
			blackspaceLength = b.Size*int64(addRate) - b.Used - putDataLength
		} else {
			return 0, BUFFERNOTENOUGH
		}
	} else {
		blackspaceLength = b.Free - putDataLength
	}
	b.Data = append(b.Data[:b.Used], putData...)
	b.Data = append(b.Data, make([]byte, blackspaceLength)...)
	b.Used = b.Used + putDataLength
	b.Free = blackspaceLength
	b.Size = b.Used + b.Free
	b.Unlock()
	return putDataLength, nil
}

// get buffer pool's content.
func (b *BufferPool) GetContent() []byte {
	return b.Data
}

// flush the contents of buffer pool.
// flush will not be reset the buffer size.
// it's means the incremental buffer will not be recover.
func (b *BufferPool) Flush(interface{}) {
	b.Lock()
	b.Data = make([]byte, b.Size)
	b.Unlock()
}
