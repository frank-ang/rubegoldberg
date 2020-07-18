#!/usr/bin/env python3

from aws_cdk import core

from data.data_stack import DataStack


app = core.App()
DataStack(app, "data")

app.synth()
