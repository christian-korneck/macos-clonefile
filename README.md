# macOS `clonefile` util

Create a Copy-On-Write clone of a file on MacOS.

## Details

`clonefile [src] [dest]` on MacOS causes the named file src to be cloned to the named file dest. The cloned file dest shares its data blocks with the src file but has its own copy of attributes, extended attributes and ACL's which are identical to those of the named file src.

This uses the `clonefile` Darwin syscall. See Apple Manpage for details and limitations:
- `man clonefile` on macOS or:
- https://www.manpagez.com/man/2/clonefile/
- https://github.com/apple-oss-distributions/xnu/blob/main/bsd/man/man2/clonefile.2

