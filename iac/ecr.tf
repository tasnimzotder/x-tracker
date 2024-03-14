locals {
  container_repos = [
    {
      name = "x-tracker-be"
    },
    {
      name = "x-tracker-fe"
    }
  ]
}

resource "aws_ecr_repository" "x-tracker-repo-fe" {
  count                = length(local.container_repos)
  name                 = local.container_repos[count.index].name
  image_tag_mutability = "MUTABLE"
  force_delete         = false

  image_scanning_configuration {
    scan_on_push = true
  }
}

output "ecr_repo_uri" {
  value = {
    for repo in aws_ecr_repository.x-tracker-repo-fe : repo.name => repo.repository_url
  }
}
