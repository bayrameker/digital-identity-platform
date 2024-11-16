provider "aws" {
  region = "us-east-1"
}

resource "aws_eks_cluster" "k8s" {
  name     = "k8s-cluster"
  role_arn = var.cluster_role_arn
  # Diğer ayarlar...
}

# Diğer kaynaklar (VPC, IAM, Node Group vb.)
