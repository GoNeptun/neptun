# Contributing to Neptune

We welcome contributions to Neptune of any kind including documentation, themes,
organization, tutorials, blog posts, bug reports, issues, feature requests,
feature implementations, pull requests, answering questions on the forum,
helping to manage issues, etc.

*Changes to the codebase **and** related documentation, e.g. for a new feature, should still use a single pull request.*

## Table of Contents

* [Reporting Issues](#reporting-issues)
* [Submitting Patches](#submitting-patches)
  * [Code Contribution Guidelines](#code-contribution-guidelines)
  * [Git Commit Message Guidelines](#git-commit-message-guidelines)
  * [Fetching the Sources From GitHub](#fetching-the-sources-from-github)
  * [Building Neptune with Your Changes](#building-neptune-with-your-changes)
  
## Reporting Issues

If you believe you have found a defect in Neptune or its documentation, use
the GitHub issue tracker to report
the problem to the Neptune maintainers.
When reporting the issue, please provide the version of Hugo in use (`neptune
version`) and your operating system.

## Code Contribution

Neptune has become a fully featured dynamic content site generator, so any new functionality must:

* be useful to many.
* fit naturally into _what Neptune does best._
* strive not to break existing sites.
* close or update an open [Neptune issue](https://github.com/goneptune/neptune/issues)

If it is of some complexity, the contributor is expected to maintain and support the new future (answer questions on the forum, fix any bugs etc.).

 If you are submitting a complex feature, create a small design proposal on the [Neptune issue tracker](https://github.com/goneptune/neptune/issues) before you start.


**Bug fixes are, of course, always welcome.**

## Submitting Patches

The Neptune project welcomes all contributors and contributions regardless of skill or experience level. If you are interested in helping with the project, we will help you with your contribution.

### Code Contribution Guidelines
Because we want to create the best possible product for our users and the best contribution experience for our developers, we have a set of guidelines which ensure that all contributions are acceptable. The guidelines are not intended as a filter or barrier to participation. If you are unfamiliar with the contribution process, the Neptune team will help you and teach you how to bring your contribution in accordance with the guidelines.

To make the contribution process as seamless as possible, we ask for the following:

* Go ahead and fork the project and make your changes.  We encourage pull requests to allow for review and discussion of code changes.
* When you’re ready to create a pull request, be sure to:
    * Sign the [CLA](https://cla-assistant.io/goneptune/neptune).
    * Have test cases for the new code. If you have questions about how to do this, please ask in your pull request.
    * Run `go fmt`.
    * Add documentation if you are adding new features or changing functionality.  The docs site lives in `/docs`.
    * Squash your commits into a single commit. `git rebase -i`. It’s okay to force update your pull request with `git push -f`.
    * Follow the **Git Commit Message Guidelines** below.
