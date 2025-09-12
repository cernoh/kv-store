# KV store in Go using Raft

## Overview

I've started this project to experiment with concurrency, databases, and how to
work around failure.

## Why go?

Its robust failure handling is really nice for databases as you can easily bump
out edge cases, and with its automatic garbage collection any memory leaks are
easily swept away. Raft helps to handle with concurrency and make sure that
failures dont take down your system.

## Why devbox?

I love how it works! I ended up switching away from it to using nix itself,
which is what devbox uses under the hood. The fact its immutable and allows you
to make projects entirely atomic makes it really easy to work on it from
multiple devices. It also makes it so that if I want other collaborators to work
on it, it's incredibly easy for them to get set up, and allows me to pin the
version of Go to make sure there aren't any compatibility issues.
