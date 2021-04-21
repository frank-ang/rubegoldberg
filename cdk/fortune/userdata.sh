#!/bin/sh

# Patch
yum -y update

# install httpd
yum install httpd -y

# enable and start httpd
systemctl enable httpd
systemctl start httpd
# dummy web server page, should be disabled by subsequent app deployment.
echo "<html><head><title> Example Web Server</title></head>" >  /var/www/html/index.html
echo "<body>" >>  /var/www/html/index.html
echo "<div><center><h2>Welcome AWS $(hostname -f) </h2>" >>  /var/www/html/index.html
echo "<hr/>" >>  /var/www/html/index.html
curl http://169.254.169.254/latest/meta-data/instance-id >> /var/www/html/index.html
echo "</center></div></body></html>" >>  /var/www/html/index.html

# CodeDeploy
yum install -y ruby
cd /home/ec2-user
curl -O https://aws-codedeploy-ap-southeast-1.s3.amazonaws.com/latest/install
chmod +x ./install
./install auto

# CloudWatch Agent
yum install -y amazon-cloudwatch-agent
## Auto-start the CW agent. Depends on the CloudWatch configuration in SSM parameter store
/opt/aws/amazon-cloudwatch-agent/bin/amazon-cloudwatch-agent-ctl -a fetch-config -m ec2 -s -c ssm:AmazonCloudWatch-linux-fortune

# XRay Agent
curl https://s3.ap-southeast-1.amazonaws.com/aws-xray-assets.ap-southeast-1/xray-daemon/aws-xray-daemon-3.x.rpm -o /home/ec2-user/xray.rpm
yum install -y /home/ec2-user/xray.rpm
