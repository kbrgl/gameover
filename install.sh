#!/bin/sh

set -e

tmpdir=$(mktemp -d 2>/dev/null || mktemp -d -t 'tmpdir')
rm -rf ~/.gameover
git clone https://github.com/kbrgl/gameover.git "${tmpdir}"
mv "${tmpdir}/.gameover" ~/.gameover

echo "  / _ \__ _ _ __ ___   ___  _____   _____ _ __ "
echo " / /_\/ _' | '_ ' _ \ / _ \/ _ \ \ / / _ \ '__|"
echo "/ /_\\ (_| | | | | | |  __/ (_) \ V /  __/ |   "
echo "\____/\__,_|_| |_| |_|\___|\___/ \_/ \___|_|"
echo "Add 'source ~/.gameover/env' to your ~/.zprofile or whatever you use"
