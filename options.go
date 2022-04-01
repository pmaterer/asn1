package asn1

import (
	"fmt"
	"strconv"
	"strings"
)

type options struct {
	private     bool
	optional    bool
	application bool
	explicit    bool
	set         bool
	enumerated  bool
	choice      bool
	timeType    Tag
	stringType  Tag
	tag         *Tag
}

func newDefaultOptions() *options {
	return &options{
		stringType: TagUtf8String,
		timeType:   TagUtcTime,
	}
}

func parseOptions(optionsString string) (*options, error) {
	opts := newDefaultOptions()
	for _, token := range strings.Split(optionsString, ",") {
		args := strings.Split(strings.TrimSpace(token), ":")
		err := parseOption(opts, args)
		if err != nil {
			return nil, err
		}
	}

	return opts, nil
}

func parseOption(opt *options, args []string) error {
	switch args[0] {
	// string types
	case "ia5":
		opt.stringType = TagIa5String
	case "printable":
		opt.stringType = TagPrintableString
	case "numeric":
		opt.stringType = TagNumericString
	case "utf8":
		opt.stringType = TagUtf8String
	case "octet":
		opt.stringType = TagOctetString
	// time types
	case "utc":
		opt.timeType = TagUtcTime
	case "generalized":
		opt.timeType = TagGeneralizedTime
	// everything else
	case "private":
		opt.private = true
	case "optional":
		opt.optional = true
	case "application":
		opt.application = true
	case "explicit":
		opt.explicit = true
	case "set":
		opt.set = true
	case "enumerated":
		opt.enumerated = true
	case "choice":
		opt.choice = true
	case "tag":
		err := parseTagOption(opt, args)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseTagOption(opts *options, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("no value given for tag")
	}
	value, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid tag value '%s'", args[1])
	}
	tag := Tag(value)
	opts.tag = &tag
	return nil
}
