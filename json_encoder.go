package flume

import "go.uber.org/zap/zapcore"

// NewJSONEncoder just hides the zap json encoder, to avoid exporting zap
func NewJSONEncoder(cfg EncoderConfig) Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig(cfg))
}
