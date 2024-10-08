####> This option file is used in:
####>   podman build, farm build
####> If file is edited, make sure the changes
####> are applicable to all of those.
#### **--security-opt**=*option*

Security Options

- `apparmor=unconfined` : Turn off apparmor confinement for the container
- `apparmor=alternate-profile` : Set the apparmor confinement profile for the
container

- `label=user:USER`     : Set the label user for the container processes
- `label=role:ROLE`     : Set the label role for the container processes
- `label=type:TYPE`     : Set the label process type for the container processes
- `label=level:LEVEL`   : Set the label level for the container processes
- `label=filetype:TYPE` : Set the label file type for the container files
- `label=disable`       : Turn off label separation for the container
- `no-new-privileges`   : Not supported

- `seccomp=unconfined` : Turn off seccomp confinement for the container
- `seccomp=profile.json` :  White listed syscalls seccomp Json file to be used as a seccomp filter
