package logger

import "os"

func (l logger) Debug(msg string) {
	l.log.Debug(msg)
}
func (l logger) Info(msg string) {
	l.log.Info(msg)
}
func (l logger) Warn(msg string) {
	l.log.Warn(msg)
}
func (l logger) Error(msg string) {
	l.log.Error(msg)
}
func (l logger) Fatal(msg string) {
	l.log.Error(msg)
	os.Exit(1)
}
