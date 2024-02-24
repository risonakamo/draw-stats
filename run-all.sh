# run server and web build. need tmux already open

set -ex
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

tmux rename-window spawn

# run server
tmux new-window -n run -c $HERE
tmux send "bash run-server.sh" Enter

# run web build
tmux split-window -h -c $HERE/time-stats-web
tmux send "pnpm watch" Enter