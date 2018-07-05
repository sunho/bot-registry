package com.botregistry.core

case class Build(repo: Repo, step: Int, tasks: List[BuildTask])

object Build {
  def apply(options: BuildTaskOptions, repo: Repo): Build = {
    val tasks = BuildTaskFactory(options, repo)
    new Build(repo, 0, tasks)
  }
}
