#!/usr/bin/env python3
import os

from aws_cdk import core as cdk
#from step.step_stack import StepStack
from aws_cdk import aws_lambda as lambda_
from aws_cdk import aws_stepfunctions as sfn
from aws_cdk import aws_stepfunctions_tasks as tasks

# For consistency with TypeScript code, `cdk` is the preferred import name for
# the CDK's core module.  The following line also imports it as `core` for use
# with examples from the CDK Developer's Guide, which are in the process of
# being updated to use `cdk`.  You may delete this import if you don't need it.
from aws_cdk import core

from fortune_serverless.fortune_serverless_stack import FortuneServerlessStack


app = core.App()
stack = FortuneServerlessStack(app, "FortuneServerlessStack",
    # If you don't specify 'env', this stack will be environment-agnostic.
    # Account/Region-dependent features and context lookups will not work,
    # but a single synthesized template can be deployed anywhere.

    # Uncomment the next line to specialize this stack for the AWS Account
    # and Region that are implied by the current CLI configuration.

    #env=core.Environment(account=os.getenv('CDK_DEFAULT_ACCOUNT'), region=os.getenv('CDK_DEFAULT_REGION')),

    # Uncomment the next line if you know exactly what Account and Region you
    # want to deploy the stack to. */

    #env=core.Environment(account='123456789012', region='us-east-1'),

    # For more information, see https://docs.aws.amazon.com/cdk/latest/guide/environments.html
    )

#  StepStack(app, construct_id="fortune-step")
hello_function = lambda_.Function(stack, "MyLambdaFunction",
                                    code=lambda_.Code.from_inline("""
                        exports.handler = (event, context, callback) => {
                            callback(null, "Hello World!");
                        }"""),
                                    runtime=lambda_.Runtime.NODEJS_12_X,
                                    handler="index.handler",
                                    timeout=cdk.Duration.seconds(25))

state_machine = sfn.StateMachine(stack, "MyStateMachine",
                                    definition=tasks.LambdaInvoke(stack, "MyLambdaTask",
                                    lambda_function=hello_function).next(
                                    sfn.Succeed(stack, "GreetedWorld")))
app.synth()
