# policy - ecr_read_access

data "aws_iam_policy_document" "ecr_read_access" {
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

resource "aws_iam_policy" "ecr_read_access" {
  name        = "EC2-ECR-Access"
  description = "Provides access to read & download data from ECR service"
  policy      = data.aws_iam_policy_document.ecr_read_access.json
}

data "aws_iam_policy_document" "ec2_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

# policy - timestream_read_access

data "aws_iam_policy_document" "timestream_read_access" {
  statement {
    effect = "Allow"
    actions = [
      "timestream:*",
      "kms:DescribeKey",
      "kms:CreateGrant",
      "kms:Decrypt",
      "dbqms:CreateFavoriteQuery",
      "dbqms:DescribeFavoriteQueries",
      "dbqms:UpdateFavoriteQuery",
      "dbqms:DeleteFavoriteQueries",
      "dbqms:GetQueryString",
      "dbqms:CreateQueryHistory",
      "dbqms:UpdateQueryHistory",
      "dbqms:DeleteQueryHistory",
      "dbqms:DescribeQueryHistory",
      "s3:ListAllMyBuckets"
    ]

    resources = ["*"]
  }
}

resource "aws_iam_policy" "timestream_read_access" {
  name        = "EC2-Timestream-Access"
  description = "Provides access to read data from Timestream service"
  policy      = data.aws_iam_policy_document.timestream_read_access.json
}

# role - ec2_ecr_timestream_access

resource "aws_iam_role" "ec2_ecr_timestream_access" {
  name               = "EC2-to-ECR-Timestream-Access"
  assume_role_policy = data.aws_iam_policy_document.ec2_assume_role_policy.json
}

resource "aws_iam_role_policy_attachment" "ec2_ecr_access" {
  role       = aws_iam_role.ec2_ecr_timestream_access.name
  policy_arn = aws_iam_policy.ecr_read_access.arn
}

resource "aws_iam_role_policy_attachment" "ec2_ecr_timestream_access" {
  role       = aws_iam_role.ec2_ecr_timestream_access.name
  policy_arn = aws_iam_policy.timestream_read_access.arn
}


# attach role to ec2 instance profile
resource "aws_iam_instance_profile" "ec2_ecr_timestream_access_profile" {
  name = "EC2-ECR-Access-Profile"
  role = aws_iam_role.ec2_ecr_timestream_access.name
}
