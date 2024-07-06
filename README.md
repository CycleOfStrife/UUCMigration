# U<sup>2</sup>CMigration: User-Unaware Container Live Migration Strategy
U<sup>2</sup>CMigration is a User-Unaware Container live Migration strategy for containerized workloads. It employs a lightweight and autonomous two-phase prediction by analyzing container memory pages across various workloads. We utilize the data shift prediction for stable memory pages (Phase-I). For unstable memory pages (Phase-II), we develop an attention-based prediction that jointly considers the spatio-temporal characteristics of memory pages and system-level features. Guided by dirty page predictions, we further develop a container live migration strategy that judiciously decides the optimal stop-and-copy iteration with the minimum amount of memory dirty pages. We have implemented an open-source prototype of U<sup>2</sup>CMigration based on the CRIU (Checkpoint/Restore In Userspace) project. Extensive prototype experiments demonstrate that U<sup>2</sup>CMigration reduces the container migration duration by 26.1%−47.9% and the downtime by 21.3%−32.6% compared with the state-of-the-art solutions.

U<sup>2</sup>CMigration is currently under the review process of Journal of Computer Science and Technology.

# Getting Started
## Dependencies

1. golang >= 1.19
2. protobuf 
3. protobuf-python
4. protobuf-c
5. protobuf-c-devel
6. protobuf-compiler
7. protobuf-devel
8. Python 3.8.8
4. numpy
5. pandas
6. sklearn
9. pytorch 2.0.1

## Installation

```shell

git clone https://github.com/CycleOfStrife/UUCMigration.git
cd criu-3.16.1
make && make install
cd runc 
make && make install
cd Podman
make && make install
```
## Run the U<sup>2</sup>CMigration System


```shell
#linux01
podman run -itd --name mybox docker.io/busybox
podman conatienr checkpoint --live-migration --predict-mode="SSPD" --dirty-file="test.csv" --ip=<linux02IP> --path="/migration" mybox

```

```
#linux02
mkdir /migration
cd /migration
podman container restore -i <the last file> mybox

```


