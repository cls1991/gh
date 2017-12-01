#!/usr/bin/env python

import sys

commit_msg_filepath = sys.argv[1]
with open(commit_msg_filepath, 'wb') as f:
    f.write("# Please include a useful commit message!")

