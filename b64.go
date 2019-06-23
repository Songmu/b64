package b64

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const cmdName = "b64"

// Run the b64
func Run(argv []string, outStream, errStream io.Writer) error {
	log.SetOutput(errStream)
	fs := flag.NewFlagSet(
		fmt.Sprintf("%s (v%s rev:%s)", cmdName, version, revision), flag.ContinueOnError)
	fs.SetOutput(errStream)
	ver := fs.Bool("version", false, "display version")
	var decode bool
	fs.BoolVar(&decode, "d", false, "decode")
	fs.BoolVar(&decode, "D", false, "decode")
	if err := fs.Parse(argv); err != nil {
		return err
	}
	if *ver {
		return printVersion(outStream)
	}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	data = bytes.TrimSuffix(data, []byte("\r\n"))
	if !decode {
		_, err := fmt.Fprintln(outStream, base64.RawStdEncoding.EncodeToString(data))
		return err
	}
	b, err := base64.RawStdEncoding.DecodeString(string(data))
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(outStream, string(b))
	return err
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%s v%s (rev:%s)\n", cmdName, version, revision)
	return err
}
