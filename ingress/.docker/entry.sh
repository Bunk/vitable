#!/bin/sh

if [ "$ENV" = "production" ]; then
    exec "$@"
else
    # If we're running in dev mode, use
    # Fresh to rebuild / rerun the app on
    # file changes
    go get github.com/pilu/fresh;
    exec fresh;
fi
