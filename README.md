# mungectl

![GitHub License](https://img.shields.io/github/license/charmed-hpc/mungectl)
[![Matrix](https://img.shields.io/matrix/ubuntu-hpc%3Amatrix.org?logo=matrix&label=ubuntu-hpc)](https://matrix.to/#/#hpc:ubuntu.com)

mungectl is a command line tool for controlling and managing the operations of a
node's [MUNGE](https://dun.github.io/munge/) authentication daemon 🔐

For more information on how to use or contribute to mungectl,
check out the [Getting Started](#-getting-started) and [Development](#-development)
sections below 👇

## ✨ Getting started

### Installation

#### Option 1: Install from source

We use [just](https://just.systems) to manage this project. We recommend that it is
installed on your system before installing mungectl from source:

```shell
git clone https://github.com/charmed-hpc/mungectl.git
cd mungectl
just && just install
```

### Usage

#### Key management

mungectl can be used to manage the MUNGE key file shared between all the nodes in your
HPC cluster using the `key` subcommand:

```shell
mungectl key generate               # Generate new munge key file.
mungectl key get                    # Get munge key file as base64-encoded string.
cat new.key.b64 | mungectl key set  # Set new munge key using base64-encoded key.
```

## 🤔 What's next?

If you want to learn more about all the things you can do with mungectl,
here are some further resources for you to explore:

* [Open an issue](https://github.com/charmed-hpc/mungectl/issues)
* [Ask a question on GitHub](https://github.com/orgs/charmed-hpc/discussions/categories/q-a)

## 🛠️ Development

mungectl uses [just](https://just.systems) as its command runner. mungectl's
[justfile](./justfile) provides some useful recipes that will definitely help
you while you're hacking and wacking on mungectl:

```shell
just fmt   # Apply formatting standards to project.
just lint  # Check project against coding style standards.
just unit  # Run unit tests.
```

If you're interested in contributing your work to mungectl, take a look at our
[contributing guidelines](./CONTRIBUTING.md) for further details.

## 🤝 Project and Community

mungectl is a project of the [Ubuntu High-Performance Computing community](https://ubuntu.com/community/governance/teams/hpc).
Interested in contributing bug fixes, patches, documentation, or feedback?
Want to join the Ubuntu HPC community? You’ve come to the right place 🤩

Here’s some links to help you get started with joining the community:

* [Ubuntu Code of Conduct](https://ubuntu.com/community/ethos/code-of-conduct)
* [Contributing guidelines](./CONTRIBUTING.md)
* [Join the conversation on Matrix](https://matrix.to/#/#hpc:ubuntu.com)
* [Get the latest news on Discourse](https://discourse.ubuntu.com/c/hpc/151)
* [Ask and answer questions on GitHub](https://github.com/orgs/charmed-hpc/discussions/categories/q-a)

## 📋 License

mungectl is free software, distributed under the GNU General Public License,
version 3.0. See the [GPLv3 LICENSE](./LICENSE) file for further details.
