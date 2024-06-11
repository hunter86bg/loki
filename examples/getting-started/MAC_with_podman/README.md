# Getting Started on a MAC M-Series (arm64-based)

- Install 'brew'
    ```
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    ```

- Install dependencies
    ```
    brew install podman
    brew install podman-compose
    brew install git
    ```

- Validate installation and binary paths
    ```
    podman-compose --version
    ```

- Create a Directory that you will map to the podman machine
    ```
    mkdir ~/LOKI
    cd ~/LOKI
    ```


- Copy the following files from this directory to ~/LOKI:

  * alloy-local-config.yaml
  * docker-compose.yaml
  * loki-config.yaml


- Build flog as there is no arm-based image
    ```
    cd ~/Downloads
    git clone https://github.com/mingrammer/flog.git
    cd flog
    podman machine init -v ~/LOKI:/var/home/core/LOKI --memory 4096
    podman machine start
    podman build -t localhost/mingrammer/flog:latest .
    podman image ls
    cd ~/LOKI
    ```

- Bring the infrastructure up
    ```
    podman-compose up -d
    sleep 10 && podman ps
    ```

- Wait for the indrastructure is up and running and responding with 'ready'
    ```
    $ curl http://localhost:3101/ready
    ready
    $ curl http://localhost:3102/ready
    ready
    ```

- Open the portal (http://http://localhost:3000) and run several queries:

    ```
    {container="evaluate-loki-flog-1"}
    ```

    ```
    {container="evaluate-loki-grafana-1"}
    ```

    ```
    {container="evaluate-loki-flog-1",status="403"}
    ```

    ```
    {container="evaluate-loki-flog-1",method="GET"}
    ```

    ```
    sum by(container) (rate({container="evaluate-loki-flog-1",status="404"} [$__auto]))
    ```