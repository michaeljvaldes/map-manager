This simple app will perform the following main functions:
- Generate maps of your minecraft world by using unmined-cli (pre-installed on your machine)
- Package the maps as a simple, bare-html website
- Deploy the site to Netlify (pre-configured by user)


Run:
1. make build
2. bin/map-manager/map-manager.exe


Setup:

Unmined
1. download and install unmined (https://unmined.net/downloads/)
2. identify path to unmined-cli.exe

Netlify
1. create account on netlify (https://www.netlify.com/)
2. create empty site
3. get api-id for netlify site
4. generate an access token on netlify account settings

Config
1. fill out values in bin/map-manager/config.yml
