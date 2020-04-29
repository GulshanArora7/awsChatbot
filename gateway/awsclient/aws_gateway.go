package awsclient

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/GulshanArora7/awsChatbot/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/s3"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// AwsGetInstances func
func AwsGetInstances(awsRegion string, status string) []domain.EC2Dictionary {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ec2Svc := ec2.New(sess, &aws.Config{Region: aws.String(awsRegion)})
	_, cerr := ec2Svc.Config.Credentials.Get()
	if cerr != nil {
		log.Printf("ERROR..!!..Unable to find AWS Crendentials..Please Check..!!")
	}
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String(status),
				},
			},
		},
	}
	result, err := ec2Svc.DescribeInstances(input)
	checkError(err)
	ec2data := []domain.EC2Dictionary{}
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			if !strings.EqualFold(*instance.State.Name, "terminated") {
				ec2dict := domain.EC2Dictionary{"InstanceId": *instance.InstanceId, "InstanceType": *instance.InstanceType, "PrivateIP": *instance.PrivateIpAddress, "State": *instance.State.Name, "ImageID": *instance.ImageId, "PublicIP": *instance.PublicIpAddress}
				ec2data = append(ec2data, ec2dict)
			}
		}
	}
	return ec2data
}

// AwsGetInstancesTag func
func AwsGetInstancesTag(awsRegion string, tags map[string]string) []domain.EC2Dictionary {
	tagName := fmt.Sprintf("tag:%s", tags["TAGNAME"])
	tagValue := tags["TAGVALUE"]
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ec2Svc := ec2.New(sess, &aws.Config{Region: aws.String(awsRegion)})
	_, cerr := ec2Svc.Config.Credentials.Get()
	if cerr != nil {
		log.Printf("ERROR..!!..Unable to find AWS Crendentials..Please Check..!!")
	}
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(tagName),
				Values: []*string{
					aws.String(strings.Join([]string{"*", tagValue, "*"}, "")),
				},
			},
		},
	}
	result, err := ec2Svc.DescribeInstances(input)
	checkError(err)
	ec2data := []domain.EC2Dictionary{}
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			if !strings.EqualFold(*instance.State.Name, "terminated") {
				ec2dict := domain.EC2Dictionary{"InstanceId": *instance.InstanceId, "InstanceType": *instance.InstanceType, "PrivateIP": *instance.PrivateIpAddress, "State": *instance.State.Name, "ImageID": *instance.ImageId, "PublicIP": *instance.PublicIpAddress}
				ec2data = append(ec2data, ec2dict)
			}
		}
	}
	return ec2data
}

// AwsGetS3Buckets func
func AwsGetS3Buckets(awsRegion string) []domain.S3Dictionary {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	s3svc := s3.New(sess, &aws.Config{Region: aws.String(awsRegion)})
	_, cerr := s3svc.Config.Credentials.Get()
	if cerr != nil {
		log.Printf("ERROR..!!..Unable to find AWS Crendentials..Please Check..!!")
	}
	input := &s3.ListBucketsInput{}
	result, err := s3svc.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Printf(aerr.Error())
			}
		} else {
			log.Printf(err.Error())
		}
	}

	s3data := []domain.S3Dictionary{}
	for _, bucket := range result.Buckets {
		s3dict := domain.S3Dictionary{"BucketName": aws.StringValue(bucket.Name), "CreationDate": bucket.CreationDate}
		s3data = append(s3data, s3dict)
	}
	return s3data
}

// AwsGetS3BucketsTag func
func AwsGetS3BucketsTag(awsRegion string, tags map[string]string) []domain.S3Dictionary {
	bucketName := tags["BUCKETNAME"]
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	s3svc := s3.New(sess, &aws.Config{Region: aws.String(awsRegion)})
	_, cerr := s3svc.Config.Credentials.Get()
	if cerr != nil {
		log.Printf("ERROR..!!..Unable to find AWS Crendentials..Please Check..!!")
	}
	input := &s3.ListBucketsInput{}
	result, err := s3svc.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Printf(aerr.Error())
			}
		} else {
			log.Printf(err.Error())
		}
	}

	s3data := []domain.S3Dictionary{}
	for _, bucket := range result.Buckets {
		if strings.EqualFold(aws.StringValue(bucket.Name), bucketName) {
			s3dict := domain.S3Dictionary{"BucketName": aws.StringValue(bucket.Name), "CreationDate": bucket.CreationDate}
			s3data = append(s3data, s3dict)
		}
	}
	return s3data
}

// AwsGetSecGroup func
func AwsGetSecGroup(awsRegion string) []domain.SGDictionary {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ec2Svc := ec2.New(sess, &aws.Config{Region: aws.String(awsRegion)})
	_, cerr := ec2Svc.Config.Credentials.Get()

	if cerr != nil {
		log.Printf("ERROR..!!..Unable to find AWS Crendentials..Please Check..!!")
	}

	descSGInput := &ec2.DescribeSecurityGroupsInput{
		MaxResults: aws.Int64(500),
	}

	descSGOut, err := ec2Svc.DescribeSecurityGroups(descSGInput)
	if err != nil {
		log.Println("Unable to get security groups. Err:", err)
		os.Exit(1)
	}

	sgdata := []domain.SGDictionary{}
	for _, secgroup := range descSGOut.SecurityGroups {
		for _, ingress := range secgroup.IpPermissions {
			if len(ingress.IpRanges) > 0 {
				for _, ipRange := range ingress.IpRanges {
					if *ipRange.CidrIp == "0.0.0.0/0" {
						var fromPort, toPort int64
						if ingress.FromPort != nil {
							fromPort = *ingress.FromPort
						}

						if ingress.ToPort != nil {
							toPort = *ingress.ToPort
						}
						sgdict := domain.SGDictionary{"SecurityGroupID": *secgroup.GroupId, "SecurityGroupName": *secgroup.GroupName, "FromPort": fromPort, "ToPort": toPort, "IngressCIDR": *ipRange.CidrIp}
						sgdata = append(sgdata, sgdict)
					}
				}
			}
		}
	}
	return sgdata
}

// AwsGetSecGroupTag func
func AwsGetSecGroupTag(awsRegion string, tags map[string]string) []domain.SGDictionary {

	tagName := fmt.Sprintf("tag:%s", tags["TAGNAME"])
	tagValue := tags["TAGVALUE"]

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ec2Svc := ec2.New(sess, &aws.Config{Region: aws.String(awsRegion)})
	_, cerr := ec2Svc.Config.Credentials.Get()

	if cerr != nil {
		log.Printf("ERROR..!!..Unable to find AWS Crendentials..Please Check..!!")
	}

	descSGInput := &ec2.DescribeSecurityGroupsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(tagName),
				Values: []*string{
					aws.String(strings.Join([]string{"*", tagValue, "*"}, "")),
				},
			},
		},
	}

	descSGOut, err := ec2Svc.DescribeSecurityGroups(descSGInput)
	if err != nil {
		log.Println("Unable to get security groups. Err:", err)
		os.Exit(1)
	}

	sgdata := []domain.SGDictionary{}
	for _, secgroup := range descSGOut.SecurityGroups {
		for _, ingress := range secgroup.IpPermissions {
			if len(ingress.IpRanges) > 0 {
				for _, ipRange := range ingress.IpRanges {
					if *ipRange.CidrIp == "0.0.0.0/0" {
						var fromPort, toPort int64
						if ingress.FromPort != nil {
							fromPort = *ingress.FromPort
						}

						if ingress.ToPort != nil {
							toPort = *ingress.ToPort
						}
						sgdict := domain.SGDictionary{"SecurityGroupID": *secgroup.GroupId, "SecurityGroupName": *secgroup.GroupName, "FromPort": fromPort, "ToPort": toPort, "IngressCIDR": *ipRange.CidrIp}
						sgdata = append(sgdata, sgdict)
					}
				}
			}
		}
	}
	return sgdata
}

// AwsGetELBv1 func
func AwsGetELBv1(awsRegion string) []domain.ELBv1Dictionary {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	elbSvc := elb.New(sess, &aws.Config{Region: aws.String("eu-central-1")})
	_, cerr := elbSvc.Config.Credentials.Get()

	if cerr != nil {
		log.Printf("ERROR..!!..Unable to find AWS Crendentials..Please Check..!!")
	}

	elbInput := &elb.DescribeLoadBalancersInput{
		PageSize: aws.Int64(100),
	}

	elbout, err := elbSvc.DescribeLoadBalancers(elbInput)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elb.ErrCodeAccessPointNotFoundException:
				fmt.Println(elb.ErrCodeAccessPointNotFoundException, aerr.Error())
			case elb.ErrCodeDependencyThrottleException:
				fmt.Println(elb.ErrCodeDependencyThrottleException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	elbv1data := []domain.ELBv1Dictionary{}
	for _, elbdetails := range elbout.LoadBalancerDescriptions {
		elbv1dict := domain.ELBv1Dictionary{"Elbv1Name": *elbdetails.LoadBalancerName, "Elbv1DNSName": *elbdetails.DNSName, "Elbv1Scheme": *elbdetails.Scheme, "Elbv1CreationDate": aws.Time(*elbdetails.CreatedTime)}
		elbv1data = append(elbv1data, elbv1dict)
	}
	return elbv1data
}

// AwsGetELBv1Tag func
func AwsGetELBv1Tag(awsRegion string, tags map[string]string) []domain.ELBv1Dictionary {

	elbName := tags["ELBNAME"]
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	elbSvc := elb.New(sess, &aws.Config{Region: aws.String("eu-central-1")})
	_, cerr := elbSvc.Config.Credentials.Get()

	if cerr != nil {
		log.Printf("ERROR..!!..Unable to find AWS Crendentials..Please Check..!!")
	}

	elbInput := &elb.DescribeLoadBalancersInput{
		LoadBalancerNames: []*string{
			aws.String(elbName),
		},
	}

	elbout, err := elbSvc.DescribeLoadBalancers(elbInput)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elb.ErrCodeAccessPointNotFoundException:
				fmt.Println(elb.ErrCodeAccessPointNotFoundException, aerr.Error())
			case elb.ErrCodeDependencyThrottleException:
				fmt.Println(elb.ErrCodeDependencyThrottleException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	elbv1data := []domain.ELBv1Dictionary{}
	for _, elbdetails := range elbout.LoadBalancerDescriptions {
		elbv1dict := domain.ELBv1Dictionary{"Elbv1Name": *elbdetails.LoadBalancerName, "Elbv1DNSName": *elbdetails.DNSName, "Elbv1Scheme": *elbdetails.Scheme, "Elbv1CreationDate": aws.Time(*elbdetails.CreatedTime)}
		elbv1data = append(elbv1data, elbv1dict)
	}
	return elbv1data
}

// AwsGetELBv2 func
func AwsGetELBv2(awsRegion string) []domain.ELBv2Dictionary {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	elbv2Svc := elbv2.New(sess, &aws.Config{Region: aws.String("eu-central-1")})
	_, cerr := elbv2Svc.Config.Credentials.Get()

	if cerr != nil {
		log.Printf("ERROR..!!..Unable to find AWS Crendentials..Please Check..!!")
	}

	elbv2Input := &elbv2.DescribeLoadBalancersInput{
		PageSize: aws.Int64(100),
	}

	elbv2out, err := elbv2Svc.DescribeLoadBalancers(elbv2Input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elb.ErrCodeAccessPointNotFoundException:
				fmt.Println(elb.ErrCodeAccessPointNotFoundException, aerr.Error())
			case elb.ErrCodeDependencyThrottleException:
				fmt.Println(elb.ErrCodeDependencyThrottleException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	elbv2data := []domain.ELBv2Dictionary{}
	for _, elbv2details := range elbv2out.LoadBalancers {
		elbv2dict := domain.ELBv2Dictionary{"Elbv2Name": *elbv2details.LoadBalancerName, "Elbv2DNSName": *elbv2details.DNSName, "Elbv2Scheme": *elbv2details.Scheme, "ELBv2Status": *elbv2details.State.Code, "Elbv2CreationDate": aws.Time(*elbv2details.CreatedTime)}
		elbv2data = append(elbv2data, elbv2dict)
	}
	return elbv2data
}

// AwsGetELBv2Tag func
func AwsGetELBv2Tag(awsRegion string, tags map[string]string) []domain.ELBv2Dictionary {

	elbv2Name := tags["ELBNAME"]
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	elbv2Svc := elbv2.New(sess, &aws.Config{Region: aws.String("eu-central-1")})
	_, cerr := elbv2Svc.Config.Credentials.Get()

	if cerr != nil {
		log.Printf("ERROR..!!..Unable to find AWS Crendentials..Please Check..!!")
	}
	elbv2Input := &elbv2.DescribeLoadBalancersInput{
		PageSize: aws.Int64(100),
	}

	elbv2out, err := elbv2Svc.DescribeLoadBalancers(elbv2Input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elb.ErrCodeAccessPointNotFoundException:
				fmt.Println(elb.ErrCodeAccessPointNotFoundException, aerr.Error())
			case elb.ErrCodeDependencyThrottleException:
				fmt.Println(elb.ErrCodeDependencyThrottleException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	elbv2data := []domain.ELBv2Dictionary{}
	for _, elbv2details := range elbv2out.LoadBalancers {
		if *elbv2details.LoadBalancerName == elbv2Name {
			elbv2Input1 := &elbv2.DescribeLoadBalancersInput{
				LoadBalancerArns: []*string{
					aws.String(*elbv2details.LoadBalancerArn),
				},
			}
			result, err1 := elbv2Svc.DescribeLoadBalancers(elbv2Input1)
			if err1 != nil {
				if aerr1, ok := err1.(awserr.Error); ok {
					switch aerr1.Code() {
					case elb.ErrCodeAccessPointNotFoundException:
						fmt.Println(elb.ErrCodeAccessPointNotFoundException, aerr1.Error())
					case elb.ErrCodeDependencyThrottleException:
						fmt.Println(elb.ErrCodeDependencyThrottleException, aerr1.Error())
					default:
						fmt.Println(aerr1.Error())
					}
				} else {
					fmt.Println(err1.Error())
				}
			}
			for _, details := range result.LoadBalancers {
				elbv2dict := domain.ELBv2Dictionary{"Elbv2Name": *details.LoadBalancerName, "Elbv2DNSName": *details.DNSName, "Elbv2Scheme": *details.Scheme, "ELBv2Status": *details.State.Code, "Elbv2CreationDate": aws.Time(*details.CreatedTime)}
				elbv2data = append(elbv2data, elbv2dict)
			}
		}
	}
	return elbv2data
}
