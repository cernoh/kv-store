# KV store in Go using Raft

## Overview

I've started this project to experiment with concurrency, databases, and how to
work around failure.

Why go? Its robust failure handling is really nice for databases as you can
easily bump out edge cases, and with its automatic garbage collection any memory
leaks are easily swept away. Raft helps to handle with concurrency and make sure
that failures dont take down your system.
