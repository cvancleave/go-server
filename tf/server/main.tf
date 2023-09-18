resource "aws_ecs_cluster" "ecs_cluster" {
  name = "${var.name}-ecs-cluster"
}

resource "aws_ecs_task_definition" "ecs_task_definition" {
  cpu                      = "128"
  family                   = "service"
  memory                   = "256"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]

  container_definitions = jsonencode([
    {
      name  = "${var.name}-ecs-container"
      image = "${var.image}"
      portMappings = [
        {
          containerPort = var.port
          hostPort      = var.port
          protocol      = "tcp"
        }
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-group         = "${aws_cloudwatch_log_group.log_group.name}"
          awslogs-region        = "${data.aws_region.current.name}"
          awslogs-stream-prefix = "ecs"
        }
      }
    }
  ])
}

resource "aws_ecs_service" "ecs_service" {
  name            = "${var.name}-ecs-service"
  cluster         = aws_ecs_cluster.ecs_cluster.id
  desired_count   = 1
  task_definition = aws_ecs_task_definition.ecs_task_definition.id
  launch_type     = "FARGATE"

  network_configuration {
    assign_public_ip = false
    subnets = [
      // TODO - add subnets here
    ]
    security_groups = [
      // TODO - add security groups here
    ]
  }

  load_balancer {
    container_name   = "${var.name}-ecs-container"
    container_port   = var.port
    target_group_arn = aws_lb_target_group.target_group.arn
  }
}

resource "aws_cloudwatch_log_group" "log_group" {
  name = "${var.name}-log-group"
}

resource "aws_lb_target_group" "target_group" {
  name                 = "${var.name}-target-group"
  port                 = var.port
  protocol             = "HTTP"
  target_type          = "ip"
  deregistration_delay = 180
}

resource "aws_lb" "load_balancer" {
  name               = "${var.name}-load-balancer"
  load_balancer_type = "application"
  subnets= [
      // TODO - add subnets here
  ]
  security_groups = [
      // TODO - add security groups here
  ]
}

resource "aws_lb_listener" "http_listener" {
  port              = 80
  protocol          = "HTTP"
  load_balancer_arn = aws_lb.load_balancer.arn

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.target_group.arn
  }
}

resource "aws_lb_listener" "https_listener" {
  port              = 443
  protocol          = "HTTPS"
  load_balancer_arn = aws_lb.load_balancer.arn
  ssl_policy        = "ELBSecurityPolicy-TLS-1-2-2017-01"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.target_group.arn
  }
}