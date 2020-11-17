package linkedin

// Cursor is interface for all nodes that have pagination/cursor
type Cursor interface {
	// SetLinkedin set Linkedin client
	SetLinkedin(ln *Linkedin)
	// Next fetch next page content
	Next() bool
	// Error return error occured while cursoring
	Error() error
}
