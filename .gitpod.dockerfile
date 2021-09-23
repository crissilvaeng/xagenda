FROM gitpod/workspace-postgres

RUN sudo apt-get update && \
    sudo rm -rf /var/lib/apt/lists/*
