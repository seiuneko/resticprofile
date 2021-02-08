//+build !darwin,!windows

package crond

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const (
	startMarker = "### this content was generated by resticprofile, please leave this line intact ###\n"
	endMarker   = "### end of resticprofile content, please leave this line intact ###\n"
)

type Crontab struct {
	entries []Entry
}

func NewCrontab(entries []Entry) *Crontab {
	return &Crontab{
		entries: entries,
	}
}

func (c *Crontab) Update(source string, addEntries bool, w io.StringWriter) error {
	var err error

	before, crontab, after, sectionFound := extractOwnSection(source)

	if sectionFound && len(c.entries) > 0 {
		for _, entry := range c.entries {
			crontab, _, err = deleteLine(crontab, entry)
			if err != nil {
				return err
			}
		}
	}

	_, err = w.WriteString(before)
	if err != nil {
		return err
	}

	if !sectionFound {
		// add a new line at the end of the file before adding our stuff
		_, err = w.WriteString("\n")
		if err != nil {
			return err
		}
	}

	_, err = w.WriteString(startMarker)
	if err != nil {
		return err
	}

	if sectionFound {
		_, err = w.WriteString(crontab)
		if err != nil {
			return err
		}
	}

	if addEntries {
		err = c.Generate(w)
		if err != nil {
			return err
		}
	}

	_, err = w.WriteString(endMarker)
	if err != nil {
		return err
	}

	if sectionFound {
		_, err = w.WriteString(after)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Crontab) Generate(w io.StringWriter) error {
	var err error
	if len(c.entries) > 0 {
		for _, entry := range c.entries {
			err = entry.Generate(w)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Crontab) LoadCurrent() (string, error) {
	buffer := &strings.Builder{}
	cmd := exec.Command("crontab", "-l")
	cmd.Stdout = buffer
	cmd.Stderr = buffer
	err := cmd.Run()
	if err != nil && strings.HasPrefix(buffer.String(), "no crontab for ") {
		// it's ok to be empty
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("%w: %s", err, buffer.String())
	}
	return cleanupCrontab(buffer.String()), nil
}

func (c *Crontab) Rewrite() error {
	crontab, err := c.LoadCurrent()
	if err != nil {
		return err
	}
	input := &bytes.Buffer{}
	err = c.Update(crontab, true, input)
	if err != nil {
		return err
	}

	cmd := exec.Command("crontab", "-")
	cmd.Stdin = input
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (c *Crontab) Remove() error {
	crontab, err := c.LoadCurrent()
	if err != nil {
		return err
	}
	buffer := &bytes.Buffer{}
	err = c.Update(crontab, false, buffer)
	if err != nil {
		return err
	}

	cmd := exec.Command("crontab", "-")
	cmd.Stdin = buffer
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func cleanupCrontab(crontab string) string {
	// this pattern detects if a header has been addded to the output of "crontab -l"
	pattern := regexp.MustCompile(`^# DO NOT EDIT THIS FILE[^\n]*\n#[^\n]*\n#[^\n]*\n`)
	// and removes it if found
	return pattern.ReplaceAllString(crontab, "")
}

// extractOwnSection returns before our section, inside, and after if found.
// - It is not returning both start and end markers.
// - If not found, it returns the file content in the first string
func extractOwnSection(crontab string) (string, string, string, bool) {
	start := strings.Index(crontab, startMarker)
	end := strings.Index(crontab, endMarker)
	if start == -1 || end == -1 {
		return crontab, "", "", false
	}
	return crontab[:start], crontab[start+len(startMarker) : end], crontab[end+len(endMarker):], true
}

func deleteLine(crontab string, entry Entry) (string, bool, error) {
	// should match a line like:
	// 00,15,30,45 * * * *	/home/resticprofile --no-ansi --config config.yaml --name profile --log backup.log backup
	search := fmt.Sprintf(`(?m)^[^#][^\n]+resticprofile[^\n]+--config %s --name %s[^\n]* %s\n`,
		regexp.QuoteMeta(entry.configFile),
		regexp.QuoteMeta(entry.profileName),
		regexp.QuoteMeta(entry.commandName),
	)
	pattern, err := regexp.Compile(search)
	if err != nil {
		return crontab, false, err
	}
	if pattern.MatchString(crontab) {
		// al least one was found
		return pattern.ReplaceAllString(crontab, ""), true, nil
	}
	return crontab, false, nil
}
