package kagus

// NirvanaWriter implements io.WriteCloser and discards everything written to it
type NirvanaWriter struct{}

// NewNirvanaWriter returns a new NirvanaWriter
func NewNirvanaWriter() *NirvanaWriter {
	return new(NirvanaWriter)
}

// Write function for the io.Writer interface
func (nw *NirvanaWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

// Close doesn't do anything and returns the error nil
func (nw *NirvanaWriter) Close() error {
	return nil
}
