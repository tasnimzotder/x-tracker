data "aws_iam_policy_document" "ec2_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "ec2_ecr_access" {
  name               = "EC2-ECR-Access"
  assume_role_policy = data.aws_iam_policy_document.ec2_assume_role_policy.json
}

data "aws_iam_policy_document" "ec2_ecr_access" {
  statement {
    effect = "Allow"
    actions = [
      "ecr:GetAuthorizationToken",
      "ecr:BatchCheckLayerAvailability",
      "ecr:GetDownloadUrlForLayer",
      "ecr:GetRepositoryPolicy",
      "ecr:DescribeRepositories",
      "ecr:ListImages",
      "ecr:DescribeImages",
      "ecr:BatchGetImage",
      "ecr:GetLifecyclePolicy",
      "ecr:GetLifecyclePolicyPreview",
      "ecr:ListTagsForResource",
      "ecr:DescribeImageScanFindings"
    ]

    resources = ["*"]
  }
}

resource "aws_iam_policy" "ec2_ecr_access" {
  name   = "EC2-ECR-Access"
  policy = data.aws_iam_policy_document.ec2_ecr_access.json
}

resource "aws_iam_role_policy_attachment" "ec2_ecr_access" {
  role       = aws_iam_role.ec2_ecr_access.name
  policy_arn = aws_iam_policy.ec2_ecr_access.arn
}

# attach role to ec2 instance profile
resource "aws_iam_instance_profile" "ec2_ecr_access" {
  name = "EC2-ECR-Access-Profile"
  role = aws_iam_role.ec2_ecr_access.name
}