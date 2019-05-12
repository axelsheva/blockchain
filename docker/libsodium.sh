curl https://download.libsodium.org/libsodium/releases/libsodium-1.0.17-stable.tar.gz --output libsodium.tar.gz
tar -xvzf libsodium.tar.gz
cd libsodium-stable
sh configure
make && make check
sudo make install
