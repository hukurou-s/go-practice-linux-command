package arguments

type CommandOpts struct {
	RecursivelyOpt       bool
	NameBeginWithADotOpt bool
	LongFormatOpt        bool
	SortOpt              bool
	ReverseArrayOpt      bool
}

var Options *CommandOpts

func GetFlags(R *bool, a *bool, l *bool, S *bool, r *bool) *CommandOpts {
	Options = &CommandOpts{*R, *a, *l, *S, *r}
	return Options
}
