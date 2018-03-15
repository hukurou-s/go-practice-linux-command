package arguments

type CommandOpts struct {
	recursivelyOpt       bool
	nameBeginWithADotOpt bool
	longFormatOpt        bool
	sortOpt              bool
	reverseArrayOpt      bool
}

var CommandOpts *CommandOpts

func GetFlags(R *bool, a *bool, l *bool, S *bool, r *bool) *CommandOpts {
	commandOpts := &CommandOpts{*R, *a, *l, *S, *r}
	return commandOpts
}
