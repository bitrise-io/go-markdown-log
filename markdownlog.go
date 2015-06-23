package markdownlog

import (
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
)

func ClearLogFile() error {
	var w io.Writer

	pth := os.Getenv("BITRISE_STEP_FORMATTED_OUTPUT_FILE_PATH")
	if pth != "" {
		f, err := os.OpenFile(pth, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		defer func() {
			err := f.Close()
			if err != nil {
				return err
			}
		}()

		w = io.MultiWriter(f, os.Stdout)
	} else {
		log.Error("No BITRISE_STEP_FORMATTED_OUTPUT_FILE_PATH defined")
	}

	_, err := w.Write([]byte(""))
	if err != nil {
		return err
	}

	log.Info("Log file cleared")

	return nil
}

func ErrorMessageToOutput(msg string) error {
	var w io.Writer

	pth := os.Getenv("BITRISE_STEP_FORMATTED_OUTPUT_FILE_PATH")
	if pth != "" {
		f, err := os.OpenFile(pth, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		defer func() {
			err := f.Close()
			if err != nil {
				return err
			}
		}()

		w = io.MultiWriter(f, os.Stderr)
	} else {
		w = io.MultiWriter(os.Stderr)
		log.Errorln("No BITRISE_STEP_FORMATTED_OUTPUT_FILE_PATH defined")
	}

	_, err := w.Write([]byte(msg))
	if err != nil {
		return err
	}

	log.Errorln(msg)

	return nil
}

func ErrorSectionToOutput(section string) error {
	msg := "\n" + section + "\n"

	return ErrorMessageToOutput(msg)
}

func ErrorSectionStartToOutput(section string) error {
	msg := section + "\n"

	return ErrorMessageToOutput(msg)
}

func MessageToOutput(msg string) error {
	var w io.Writer

	pth := os.Getenv("BITRISE_STEP_FORMATTED_OUTPUT_FILE_PATH")
	if pth != "" {
		f, err := os.OpenFile(pth, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		defer func() {
			err := f.Close()
			if err != nil {
				return err
			}
		}()

		w = io.MultiWriter(f, os.Stdout)
	} else {
		w = io.MultiWriter(os.Stdout)
		log.Error("No BITRISE_STEP_FORMATTED_OUTPUT_FILE_PATH defined")
	}

	_, err := w.Write([]byte(msg))
	if err != nil {
		return err
	}

	log.Infoln(msg)

	return nil
}

func SectionToOutput(section string) error {
	msg := "\n" + section + "\n"

	return MessageToOutput(msg)
}

func SectionStartToOutput(section string) error {
	msg := section + "\n"

	return MessageToOutput(msg)
}
