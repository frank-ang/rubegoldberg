#!/usr/bin/env python3

from aws_cdk import core

from data.data_stack import DataStack
env_SIN = core.Environment(region="ap-southeast-1")

app = core.App()
DataStack(app, "service-data", env=env_SIN)

app.synth()
