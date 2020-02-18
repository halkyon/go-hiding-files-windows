# What is this?

A quick example on using Golang to set the hidden file attribute on a file.

This could be abstracted into a "HideFile" function of some sort, which could move the file
into a file prefixed with "." if you're on a UNIX based operating system, and set this
hidden attribute if on Windows. This would probably need to use Go build constraints so it
doesn't complain about the Windows syscalls library not existing.
