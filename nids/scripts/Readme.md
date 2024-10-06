sudo pacman -S base-devel cmake make gcc flex bison libpcap libmaxminddb openssl python3 swig zlib
sudo pacman -Sy zeek (will take a long time to compile)
To initialize zeek,
sudo zeekctl deploy
sudo nano /opt/zeek/etc/node.cfg
[zeek]
type=standalone
host=localhost
interface=wlp1s0  # Change it with your actual interface name


give executable permission for the scripts in this folder
