# what this is
go backend api for dnd monster manual

simple MVP receiving API get requests:

localhost:8088/monster?name=demogorgon

returns JSON with monster stat block

<hr />

# example setup #####################################

# make host dir
  mkdir radappname; cd radappname

# make share dir to bridge between container and host
  mkdir share 

# spin up image
  docker build -t gomm-image --no-cache .

# run container
  docker run -d -it -p {some port number on your local machine e.g. 8000}:3000 -v /run/host-services/ssh-auth.sock:/ssh-agent -e SSH_AUTH_SOCK='/ssh-agent' -v ./share:/root/share -v ~/.gitconfig:/etc/gitconfig --name gomm-container gomm-image

# NOTES
this assumes that you're running ssh-agent to forward ssh keys (for git)
it also assumes that you have a global ~/.gitconfig on the host machine
