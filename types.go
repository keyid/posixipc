package posixipc

type isemaphore interface {
	// Open establishes a connection between a name semaphore
	// and a process, LWP or thread
	Open(name string, oflag int) (*sem, error) // sem_open
	// Init is used to initialize the unnamed semaphores referred
	// too by *sem. The value of the initialized semaphore is `value`.
	// If `pshared` has a non-zero value, then the semaphore is shared
	// between processes; in this case any process that can
	// access the semaphore `sem` for performing
	// Wait(), TryWait(), Post(), and Destroy() operations
	Init(sem_t *sem, pshared int, value int) error // sem_init
	// Close is used to indicate that the calling process is finished
	// used the named semaphore indicated by `sem`. The sem_close() function
	// deallocates any system resources allocated by the system for use
	// by this process for this semaphore.
	Close(sem_t *sem) error // sem_close
	// Unlink removes the semaphore named by the string `name`. If the
	// semaphore named by `name` is currently referenced by other
	// processes, then sem_unlink has no effect on the state of the semaphore.
	Unlink(name *string) error // sem_unlink
	// Destory is used to destroy the unnamed semaphore indicated by `sem`.
	// Only a semaphore that was create using sem_init() may be destroyed via sem_destroy().
	Destroy(sem_t *sem) error // sem_destroy
	// GetValue updates the location referenced by the `*sval` argument
	// to have the value of the semaphore referenced by `sem`
	// without affecting the state of the semaphore.
	// The value `sval` may be 0 or positive. If `sval` is 0, there
	// may be other processes (or LWPs or threads) waiting for the
	// semaphore; if `sval` is positive, no processed is waiting.
	GetValue(sem_t *sem, sval *int) error
	// Wait locks the semaphore referenced by `sem` by performing a semaphore
	// lock operation on that semaphore. If the semaphore value is
	// currently zero, then the calling thread will not return from the call
	// to sem_wait() until it either locks the semaphore or the call is
	// interrupted by a signal. Upon success, the state of the semaphore
	// remains locked until sem_post() is executed successfully
	Wait(sem_t *sem)
	TryWait(sem_t *sem)
	// Post unlocks the semaphore referenced by `sem` by performing
	// a semaphore unlock operation on that semaphore
	Post(*sem) error
}

type imq interface {
	// Open establishes the connection between a process and a message queue
	// with a message queue descriptor. It creates an open message queue
	// description that refers to the new message queue, used by other functions
	// to refer to the message queue
	Open(name *string, oflag int) error // mq_open
	// Close removes the association between the message queue descriptor,
	// `mqdes`, and its message queue.
	Close(mqd_t mqdes) error // mq_close
	// Unlink removes the message queue named by the pathname `name`
	Unlink(name string) error // mq_unlink
	// Send adds the message pointed to by the argument `msg_ptr` to the
	// message queue specified by `mydes`. The `msg_len` argument specifies
	// the length of the message in bytes pointed to by `msg_ptr`. The value
	// of `msg_len` is less than or equal to the `mq_msgsize` or the call fails
	Send(mqd_t mqdes, msg_ptr uintptr, msg_len []byte, msg_prio int) error // mq_send
	// Recieve is used to recieve the oldest of the highest priority message(s)
	// from the message queue specified by `mqdes`. The message is remove from
	// the queue and copied to the buffer pointed to by `msg_ptr`
	Recieve(mqd_t mqdes, msg_ptr uintptr, msg_len []byte, int msg_prio) error // mq_recieve
	// Notify provides an asynchronous mechanism for processes to recieve notice
	// that messages are available in a message queue, rather than synchonously
	// blocking (waiting) in mq_recieve
	Notify(mqd_t mqdes, notification sigevent) error
	// SetAttr is used to set attributes associated with the open message queue description referenced by the message queue descriptor specified by mqdes.
	SetAttr(mqd_t mqdes, mqstat mq_attr, omqstat mq_attr) error
	// GetAttr is used to get attributes associated with the open message queue description referenced by the message queue descriptor specified by mqdes.
	GetAttr(mqd_t mqdes, mqstat mq_attr, omqstat mq_attr) error
}

type imem interface {
	Open(io.File)
	mmap()
}
