package arguments

type CommandOpts struct {
	recursivelyOpt       bool
	nameBeginWithADotOpt bool
	longFormatOpt        bool
	sortOpt              bool
	reverseArrayOpt      bool
}

var Options *CommandOpts

func (opts *CommandOpts) RecursivelyOpt() bool {
	return opts.recursivelyOpt
}

func (opts *CommandOpts) NameBeginWithADotOpt() bool {
	return opts.nameBeginWithADotOpt
}

func (opts *CommandOpts) LongFormatOpt() bool {
	return opts.longFormatOpt
}

func (opts *CommandOpts) SortOpt() bool {
	return opts.sortOpt
}

func (opts *CommandOpts) ReverseArrayOpt() bool {
	return opts.reverseArrayOpt
}

func GetFlags(R *bool, a *bool, l *bool, S *bool, r *bool) *CommandOpts {
	Options = &CommandOpts{*R, *a, *l, *S, *r}
	return Options
}
