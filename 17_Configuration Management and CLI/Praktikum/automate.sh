#!/bin/bash

mkdir "$1 at $(date)"

cd "$1 at $(date)"

mkdir about_me
mkdir about_me/personal
mkdir about_me/professional
mkdir my_friends
mkdir my_system_info

echo "https://www.facebook.com/$2" > about_me/personal/facebook.txt
echo "https://www.linkedin.com/in/$3" > about_me/professional/linkedin.txt

curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt > my_friends/list_of_my_friends.txt

echo "My username: $(whoami)" > my_system_info/about_this_laptop.txt
echo "With host: $(uname -a)" >> my_system_info/about_this_laptop.txt

ping -c 3 google.com >my_system_info/internet_connection.txt
