#!/bin/bash
sudo docker pull a48zhang/melting:dev
sudo docker run -dp 65000:65000 --env-file ./melting.env a48zhang/melting:dev
