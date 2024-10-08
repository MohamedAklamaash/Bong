sudo pacman -Sy mailutils
sudo pacman -Sy msmtp
sudo pacman -Sy libpcap

g++ -o ../../cmd/nids_monitor usingrawsockets.cpp
