# termios

This is a pure-go implementation of (most of) the functions described
in `termios(3)` on Linux (although you should probably read
`tcgetattr(3)` and `tcsendbreak(3)` from FreeBSD; at least to me they were
much clearer).

Only tested on Linux, and will need splitting into arch-specific bits
to work on \*BSD, but no strong reason for it not to work
there. Patches welcome!
