package binary

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"

	"github.com/projecteru2/core/log"
)

// calls the plugin and gets json response
func (p Plugin) call(ctx context.Context, cmd string, req interface{}, resp interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, p.config.ResourcePlugin.CallTimeout)
	defer cancel()
	logger := log.WithFunc("resource.binary.call")

	command := exec.CommandContext(ctx, p.path, cmd) // nolint
	command.Dir = p.config.ResourcePlugin.Dir
	logger.Infof(ctx, "call %s", command.String())

	out, err := p.execCommand(ctx, command, req)
	if err != nil {
		logger.Error(ctx, err, string(out))
		return err
	}

	if len(out) == 0 {
		return nil
	}
	return json.Unmarshal(out, resp)
}

func (p Plugin) execCommand(ctx context.Context, cmd *exec.Cmd, req interface{}) ([]byte, error) {
	logger := log.WithFunc("resource.binary.execCommand")
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	logger.WithField("in", string(b)).Debug(ctx, "call params")
	cmd.Stdin = bytes.NewBuffer(b)
	return cmd.CombinedOutput()
}